package repository

import (
	"github.com/jmoiron/sqlx"
)

// Author - author commands (not implemented, because is not required)
type Author interface {
}

// AuthorPostgres - repo
type AuthorPostgres struct {
	db *sqlx.DB
}

// NewAuthorPostgres - constructor
func NewAuthorPostgres(db *sqlx.DB) *AuthorPostgres {
	return &AuthorPostgres{db: db}
}
