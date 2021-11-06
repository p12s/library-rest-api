package service

import (
	"github.com/p12s/library-rest-api/internal/repository"
)

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
