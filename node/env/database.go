package env

import (
	"os"
	"path/filepath"
)

func GetDatabasePath() string {
	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "node.db"
	}

	absPath, err := filepath.Abs(dbPath)
	if err != nil {
		return dbPath
	}

	return absPath
}
