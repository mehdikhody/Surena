package env

import (
	"os"
	"strconv"
)

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
