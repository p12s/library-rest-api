package grpcClient

import (
	"fmt"

	"github.com/p12s/library-rest-api/library/pb"
	"github.com/p12s/library-rest-api/library/pkg/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Client
type Client struct {
	Service pb.LoggerServiceClient
}

// New
func New(cfg config.Config) (*Client, error) {
	conn, err := grpc.Dial(cfg.Logger.Host, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("error grpc client: %s\n", err.Error())
		return nil, fmt.Errorf("error grpc client: %w\n", err)
	}

	return &Client{
		Service: pb.NewLoggerServiceClient(conn),
	}, nil
}
