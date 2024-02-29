package xray

import (
	"github.com/xtls/xray-core/common/errors"
	"surena/node/utils"
	"surena/node/xray/api"
	"surena/node/xray/config"
	"surena/node/xray/core"
)

var xray *Xray
var logger = utils.CreateLogger("xray")

type Xray struct {
	XrayInterface
	Config config.ConfigInterface
	Core   core.CoreInterface
	API    api.APIInterface
}

type XrayInterface interface {
	GetConfig() config.ConfigInterface
	GetCore() core.CoreInterface
}

func init() {
	logger.Debug("initializing Xray")

	cf, err := config.Get()
	if err != nil {
		logger.Error("could not load Xray config")
		return
	}

	xray = &Xray{
		Config: cf,
	}

	xray.Core, err = core.Initialize()
	if err != nil {
		logger.Error("could not initialize Xray core")
		return
	}

	xray.API, _ = api.Get()
}

func Initialize() (XrayInterface, error) {
	if xray == nil {
		return nil, errors.New("Xray not initialized")
	}

	return xray, nil
}

func Get() XrayInterface {
	if xray == nil {
		panic("Xray not initialized")
	}

	return xray
}

func (x *Xray) GetConfig() config.ConfigInterface {
	return x.Config
}

func (x *Xray) GetCore() core.CoreInterface {
	return x.Core
}
