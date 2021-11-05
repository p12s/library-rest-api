package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p12s/library-rest-api/pkg/models"
)

// AuthPostgres - репозиторий
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres - конструктор объекта репозитория
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser - создание пользователя
func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	tx, err := r.db.Begin() // TODO убрать отсюда транзакцию в создание книги
	if err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return 0, rollbackErr
		}
		return 0, err
	}

	return id, tx.Commit()
}

// GetUser - получение пользователя из БД
func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
