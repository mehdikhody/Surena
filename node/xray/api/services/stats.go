package services

import (
	"context"
	"github.com/xtls/xray-core/app/stats/command"
	"google.golang.org/grpc"
	"time"
)

type StatsService struct {
	client  *grpc.ClientConn
	service command.StatsServiceClient
}

func NewStatsService(client *grpc.ClientConn) *StatsService {
	service := command.NewStatsServiceClient(client)

	return &StatsService{
		client:  client,
		service: service,
	}
}

func (s *StatsService) QueryStats(reset bool) (*command.QueryStatsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	stats, err := s.service.QueryStats(ctx, &command.QueryStatsRequest{
		Reset_: reset,
	})

	if err != nil {
		return nil, err
	}

	return stats, nil
}
