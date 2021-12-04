package repository

import (
	_ "github.com/golang/mock/mockgen/model"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -destination mocks/mock.go -package repository github.com/p12s/library-rest-api/library/pkg/repository Authorization,Author,Book

// Repository - repo
type Repository struct {
	Authorization
	Book
	Author
}

// NewRepository - constructor
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Book:          NewBookPostgres(db),
		Author:        NewAuthorPostgres(db),
	}
}
