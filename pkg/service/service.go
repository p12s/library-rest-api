package service

import (
	"github.com/p12s/library-rest-api/pkg/models"
	"github.com/p12s/library-rest-api/pkg/repository"
)

// Authorization - signup/signin
type Authorization interface {
	CreateUser(user models.User) (int, error)
}

// Book - book commands
type Book interface {
	Create(book models.Book) (int, error)
	GetAll() ([]models.Book, error)
	GetById(bookId int) (models.Book, error)
	Delete(bookId int) error
	Update(bookId int, book models.Book) error
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
		//Book:          NewBookService(repos.Book, repos.Author),
		Book:   NewBookService(repos.Book),
		Author: NewAuthorService(repos.Author),
	}
}
