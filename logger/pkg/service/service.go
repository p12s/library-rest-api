package service

import (
	"context"

	"github.com/p12s/library-rest-api/logger/pb"
	"github.com/p12s/library-rest-api/logger/pkg/models"
	"github.com/p12s/library-rest-api/logger/pkg/repository"
)

// Logger
type Logger interface {
	Insert(ctx context.Context, item models.LogItem) error
}

// LogService
type LogService struct {
	repo *repository.Repository
}

// NewLogService
func NewLogService(repo *repository.Repository) *LogService {
	return &LogService{
		repo: repo,
	}
}

// Insert
func (s *LogService) Insert(ctx context.Context, req *pb.LoggerRequest) error {
	item := models.LogItem{
		Action:    req.GetAction().String(),
		Entity:    req.GetEntity().String(),
		EntityId:  req.GetEntityId(),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return s.repo.Insert(ctx, item)
}
