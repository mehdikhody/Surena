package xray

import (
	"surena/node/xray/core"
)

var xray *Xray

type Xray struct {
	XrayInterface
	Core core.CoreInterface
}

type XrayInterface interface {
	GetCore() core.CoreInterface
}

func init() {
	xray = &Xray{
		Core: core.Get(),
	}
}

func Get() XrayInterface {
	return xray
}

func (x *Xray) GetCore() core.CoreInterface {
	return x.Core
}
