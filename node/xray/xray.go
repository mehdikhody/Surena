package xray

import (
	"surena/node/xray/api"
	"surena/node/xray/core"
)

var xray *Xray

type Xray struct {
	Core *core.Core
	API  *api.API
}

func New() *Xray {
	xray = &Xray{}

	var err error

	xray.Core, err = core.New()
	if err != nil {
		panic(err)
	}

	xray.API, err = api.New("127.0.0.1", 30002)
	if err != nil {
		panic(err)
	}

	return xray
}

func Get() *Xray {
	return xray
}
