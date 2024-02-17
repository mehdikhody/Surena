package env

import (
	"os"
	"path/filepath"
)

func GetLogLevel() string {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	return logLevel
}

func GetLogDirectory() string {
	logDirectory := os.Getenv("LOG_DIRECTORY")
	if logDirectory == "" {
		logDirectory = "logs"
	}

	absPath, err := filepath.Abs(logDirectory)
	if err != nil {
		return logDirectory
	}

	return absPath
}
