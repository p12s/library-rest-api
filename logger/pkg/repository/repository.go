package repository

import (
	"context"

	"github.com/p12s/library-rest-api/logger/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

const LOG_COLLECTION_NAME = "logs"
const DB_NAME = "audit"

// Repository
type Repository struct {
	db *mongo.Client
}

// NewRepository
func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		db: db,
	}
}

// Insert - save data to DB
func (r *Repository) Insert(ctx context.Context, item models.LogItem) error {
	_, err := r.db.Database(DB_NAME).Collection(LOG_COLLECTION_NAME).InsertOne(ctx, item)

	return err
}
