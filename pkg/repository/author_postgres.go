package repository

import (
	"github.com/jmoiron/sqlx"
)

// AuthorPostgres - repo
type AuthorPostgres struct {
	db *sqlx.DB
}

// NewAuthorPostgres - constructor
func NewAuthorPostgres(db *sqlx.DB) *AuthorPostgres {
	return &AuthorPostgres{db: db}
}
