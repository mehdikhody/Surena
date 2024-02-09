package xray

import "surena/node/xray/api"

var xray *Xray

type Xray struct {
	API *api.API
}

func Init() *Xray {
	xray = &Xray{}

	var err error
	xray.API, err = api.New("127.0.0.1", 30002)
	if err != nil {
		panic(err)
	}

	return xray
}

func GetXray() *Xray {
	return xray
}
