package utils

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"surena/node/env"
)

func CreateLogger(name string) *logrus.Entry {
	logger := logrus.New()

	// set log level
	logLevel := getLogrusLevel(env.GetLogLevel())
	logger.SetLevel(logLevel)

	// set log format
	logger.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"label", "module"},
	})

	// set log output to stdout
	logger.SetOutput(os.Stdout)

	// set log output to file
	logDirectory := env.GetLogDirectory()
	logger.Hooks.Add(lfshook.NewHook(
		lfshook.PathMap{
			logrus.TraceLevel: fmt.Sprintf("%s/%s.log", logDirectory, name),
			logrus.DebugLevel: fmt.Sprintf("%s/%s.log", logDirectory, name),
			logrus.InfoLevel:  fmt.Sprintf("%s/%s.log", logDirectory, name),
			logrus.WarnLevel:  fmt.Sprintf("%s/%s.log", logDirectory, name),
			logrus.ErrorLevel: fmt.Sprintf("%s/%s.log", logDirectory, name),
			logrus.FatalLevel: fmt.Sprintf("%s/%s.log", logDirectory, name),
			logrus.PanicLevel: fmt.Sprintf("%s/%s.log", logDirectory, name),
		},
		&logrus.JSONFormatter{},
	))

	return logger.WithField("label", name)
}

func getLogrusLevel(level string) logrus.Level {
	switch strings.ToLower(level) {
	case "trace":
		return logrus.TraceLevel
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}
