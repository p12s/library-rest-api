package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/p12s/library-rest-api/pkg/models"
)

// Authorization - signup/signin
type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

// Book - book commands
type Book interface {
	CreateBook(book models.Book) (int, error)
	GetAllBook() ([]models.Book, error)
	GetBookById(bookId int) (models.Book, error)
	DeleteBook(bookId int) error
	UpdateBook(bookId int, book models.Book) error
}

// Author - author commands (not implemented, because is not required)
type Author interface {
}

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
