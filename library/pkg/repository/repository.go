package repository

import (
	"github.com/jmoiron/sqlx"
)

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
