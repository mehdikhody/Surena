package utils

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strconv"
)

func LoadEnv() {
	const envFile = ".env"
	envPath, err := filepath.Abs(envFile)
	if err != nil {
		return
	}

	_, err = os.Stat(envPath)
	if os.IsNotExist(err) {
		return
	}

	_ = godotenv.Load(envPath)
}

func GetTimezone() string {
	timezone := os.Getenv("TZ")
	if timezone == "" {
		timezone = "UTC"
	}

	return timezone
}

func GetDatabasePath() string {
	dbpath := os.Getenv("DATABASE_PATH")
	if dbpath == "" {
		dbpath = "db/node.db"
	}

	absolutePath, err := filepath.Abs(dbpath)
	if err != nil {
		return dbpath
	}

	return absolutePath
}

func GetServerHost() string {
	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	return host
}

func GetServerPort() int {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "3000"
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	return portInt
}
