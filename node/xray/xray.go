package xray

import (
	"surena/node/xray/api"
	"surena/node/xray/core"
)

var xray *Xray
var xrayInitialized bool = false

type Xray struct {
	core *core.Core
	api  *api.API
}

func Initialize() *Xray {
	if xrayInitialized {
		panic("Xray already initialized")
	}

	xray = &Xray{}

	var err error

	xray.core, err = core.New()
	if err != nil {
		panic(err)
	}

	xray.api, err = api.New("127.0.0.1", 30002)
	if err != nil {
		panic(err)
	}

	xrayInitialized = true
	return xray
}

func Get() *Xray {
	if !xrayInitialized {
		panic("Xray not initialized")
	}

	return xray
}

func (x *Xray) GetCore() *core.Core {
	return x.core
}

func (x *Xray) GetAPI() *api.API {
	return x.api
}
