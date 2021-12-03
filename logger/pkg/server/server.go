package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/p12s/library-rest-api/logger/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server
type Server struct {
	grpcSrv      *grpc.Server
	loggerServer pb.LoggerServiceServer
}

// New
func New(loggerServer pb.LoggerServiceServer) *Server {
	return &Server{
		grpcSrv:      grpc.NewServer(),
		loggerServer: loggerServer,
	}
}

// ListenAndServe
func (s *Server) ListenAndServe(port int) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Fatalf("error grpc listening: %s\n", err.Error())
	}

	pb.RegisterLoggerServiceServer(s.grpcSrv, s.loggerServer)
	reflection.Register(s.grpcSrv)

	go func() {
		if err := s.grpcSrv.Serve(listen); err != nil {
			logrus.Fatalf("error while running grpc server: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Printf("Server shutting down. %s", time.Now().Format(time.RFC3339))
	s.grpcSrv.GracefulStop()
}
