package config

import (
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/infra/conf"
	"github.com/xtls/xray-core/infra/conf/serial"
	"os"
	"surena/node/env"
	"surena/node/utils"
)

var config *Config

type Config struct {
	Logger *logrus.Entry
	Config *conf.Config
}

type ConfigInterface interface {
	GetAPIAddress() (string, error)
}

func init() {
	logger := utils.CreateLogger("xray").WithField("module", "config")
	logger.Debug("initializing Xray config")

	configFile, err := os.Open(env.GetXrayConfigPath())
	if err != nil {
		logger.Error("could not load Xray config")
		return
	}

	defer configFile.Close()

	cf, err := serial.DecodeJSONConfig(configFile)
	if err != nil {
		logger.Error("could not decode Xray config")
		return
	}

	config = &Config{
		Logger: logger,
		Config: cf,
	}
}

func Get() (*Config, error) {
	if config == nil {
		return nil, errors.New("Xray config not initialized")
	}

	return config, nil
}
