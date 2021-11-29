package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p12s/library-rest-api/library/pkg/models"
)

// Authorization - signup/signin
type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

// AuthPostgres - repo
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres - constructor
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser - creating user
func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (name, username, password_hash) 
		values ($1, $2, $3) RETURNING id`, usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("create user: %w", err)
	}

	return id, nil
}

// GetUser - getting user
func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf(`SELECT id FROM %s WHERE username=$1 AND password_hash=$2`, usersTable)
	err := r.db.Get(&user, query, username, password)
	if err != nil {
		return user, fmt.Errorf("get user: %w", err)
	}

	return user, err
}
