package service

import (
	"github.com/p12s/library-rest-api/pkg/models"
	"github.com/p12s/library-rest-api/pkg/repository"
)

// Authorization - signup/signin
type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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

// Service - just service
type Service struct {
	Authorization
	Book
	Author
}

// NewService - constructor
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Book:          NewBookService(repos.Book, repos.Author),
		Author:        NewAuthorService(repos.Author),
	}
}
