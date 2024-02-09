package api

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"surena/node/xray/api/services"
)

type API struct {
	client       *grpc.ClientConn
	StatsService *services.StatsService
}

func New(host string, port int) (*API, error) {
	target := fmt.Sprintf("%s:%d", host, port)
	credentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	client, err := grpc.Dial(target, credentials)

	if err != nil {
		return nil, err
	}

	return &API{
		client:       client,
		StatsService: services.NewStatsService(client),
	}, nil
}
