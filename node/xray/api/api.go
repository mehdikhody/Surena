package api

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"surena/node/utils"
	"surena/node/xray/api/services"
	"surena/node/xray/config"
)

var api *API

type API struct {
	Logger       *logrus.Entry
	Client       *grpc.ClientConn
	StatsService *services.StatsService
}

type APIInterface interface {
}

func init() {
	logger := utils.CreateLogger("xray").WithField("module", "api")
	logger.Debug("initializing Xray api")

	cf, err := config.Get()
	if err != nil {
		logger.Error("could not get Xray config")
		return
	}

	address, err := cf.GetAPIAddress()
	if err != nil {
		logger.Error("could not find Xray api address from config")
		return
	}

	credentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	client, err := grpc.Dial(address, credentials)
	if err != nil {
		logger.Error("could not dial Xray api")
		return
	}

	api = &API{
		Logger:       logger,
		Client:       client,
		StatsService: services.NewStatsService(client),
	}
}

func Get() (APIInterface, error) {
	if api == nil {
		return nil, fmt.Errorf("xray api not initialized")
	}

	return api, nil
}
