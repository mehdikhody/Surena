package xray

import (
	"github.com/sirupsen/logrus"
	"github.com/xtls/xray-core/common/errors"
	"surena/node/utils"
	"surena/node/xray/api"
	"surena/node/xray/config"
	"surena/node/xray/core"
)

var xray *Xray

type Xray struct {
	XrayInterface
	Logger *logrus.Entry
	Config config.ConfigInterface
	Core   core.CoreInterface
	API    api.APIInterface
}

type XrayInterface interface {
	GetConfig() config.ConfigInterface
	GetCore() core.CoreInterface
}

func init() {
	logger := utils.CreateLogger("xray")
	logger.Debug("initializing Xray")

	cf, err := config.Get()
	if err != nil {
		logger.Error("could not load Xray config")
		return
	}

	xray = &Xray{Config: cf}
	xray.Core, _ = core.Get()
	xray.API, _ = api.Get()
}

func Get() (XrayInterface, error) {
	if xray == nil {
		return nil, errors.New("Xray not initialized")
	}

	return xray, nil
}

func (x *Xray) GetConfig() config.ConfigInterface {
	return x.Config
}

func (x *Xray) GetCore() core.CoreInterface {
	return x.Core
}
