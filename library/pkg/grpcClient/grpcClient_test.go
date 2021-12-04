package grpcClient

import (
	"context"
	"net"
	"testing"

	"github.com/p12s/library-rest-api/library/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	DEFAULT_ENTITY_ID = 1
	BUF_SIZE          = 1024 * 1024
)

var fakeListener *bufconn.Listener

type server struct{}

func (*server) Log(context.Context, *pb.LoggerRequest) (*pb.EmptyResponse, error) {
	return &pb.EmptyResponse{}, nil
}

func init() {
	fakeListener = bufconn.Listen(BUF_SIZE)
	s := grpc.NewServer()

	pb.RegisterLoggerServiceServer(s, &server{})

	go func() {
		if err := s.Serve(fakeListener); err != nil {
			assert.ObjectsAreEqual(nil, err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return fakeListener.Dial()
}

func TestLog(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		Entity pb.LoggerRequest_Entities
		Action pb.LoggerRequest_Actions
	}{
		{"Client can send user register log", pb.LoggerRequest_USER, pb.LoggerRequest_REGISTER},
		{"Client can send user login log", pb.LoggerRequest_USER, pb.LoggerRequest_LOGIN},
		{"Client can send user update log", pb.LoggerRequest_USER, pb.LoggerRequest_UPDATE},
		{"Client can send user get log", pb.LoggerRequest_USER, pb.LoggerRequest_GET},
		{"Client can send user delete log", pb.LoggerRequest_USER, pb.LoggerRequest_DELETE},

		{"Client can send user delete log", pb.LoggerRequest_BOOK, pb.LoggerRequest_CREATE},
		{"Client can send user delete log", pb.LoggerRequest_BOOK, pb.LoggerRequest_UPDATE},
		{"Client can send user delete log", pb.LoggerRequest_BOOK, pb.LoggerRequest_GET},
		{"Client can send user delete log", pb.LoggerRequest_BOOK, pb.LoggerRequest_DELETE},
	}

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		assert.Equal(t, nil, err)
	}
	defer conn.Close()
	client := pb.NewLoggerServiceClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			response, err := client.Log(ctx, &pb.LoggerRequest{
				Entity:    tt.Entity,
				Action:    tt.Action,
				EntityId:  DEFAULT_ENTITY_ID,
				Timestamp: timestamppb.Now(),
			})

			if err != nil {
				assert.Equal(t, nil, err)
			}
			assert.IsType(t, &pb.EmptyResponse{}, response)
		})
	}
}
