package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

func NewLogger(name string) (*zap.SugaredLogger, error) {
	err := makeLogsDirectory()
	if err != nil {
		return nil, err
	}

	logFile := getLogFile(name)
	logLevel := getZapLogLevel()

	config := zap.NewProductionConfig()
	config.Encoding = "console"
	config.Level = zap.NewAtomicLevelAt(logLevel)
	config.OutputPaths = []string{"stdout", logFile}
	config.ErrorOutputPaths = []string{"stderr", logFile}

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}

	sugar := logger.Sugar()
	defer sugar.Sync()

	return sugar, nil
}

func GetLogsDirectory() string {
	logsDirectory := os.Getenv("LOGS_DIRECTORY")
	if logsDirectory == "" {
		logsDirectory = "logs"
	}

	absolutePath, _ := filepath.Abs(logsDirectory)
	return absolutePath
}

func GetLogLevel() string {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	return logLevel
}

func makeLogsDirectory() error {
	logsDirectory := GetLogsDirectory()
	return os.MkdirAll(logsDirectory, os.ModePerm)
}

func getLogFile(name string) string {
	logsDirectory := GetLogsDirectory()
	return filepath.Join(logsDirectory, name+".log")
}

func getZapLogLevel() zapcore.Level {
	switch GetLogLevel() {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
