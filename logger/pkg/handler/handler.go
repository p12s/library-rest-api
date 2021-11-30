package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/p12s/library-rest-api/logger/pb"
	"github.com/p12s/library-rest-api/logger/pkg/service"
	"github.com/sirupsen/logrus"
)

// Handler
type Handler struct {
	services *service.LogService
}

// NewHandler
func NewHandler(services *service.LogService) *Handler {
	return &Handler{services: services}
}

// Log - logging
func (h *Handler) Log(ctx context.Context, req *pb.LoggerRequest) (*pb.EmptyResponse, error) {
	err := h.services.Insert(ctx, req)
	if err != nil {
		logrus.Errorf("log handler: %s %s %d %s %s\n", time.Now().Format(time.RFC3339), req.Action, req.EntityId, req.Entity, req.Timestamp)
		return nil, fmt.Errorf("log handler: %w\n", err)
	}

	logrus.Printf("save log: %s %s %d %s %s\n", time.Now().Format(time.RFC3339), req.Action, req.EntityId, req.Entity, req.Timestamp)

	return &pb.EmptyResponse{}, err
}
