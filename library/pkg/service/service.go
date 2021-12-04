package service

import (
	_ "github.com/golang/mock/mockgen/model"
	"github.com/p12s/library-rest-api/library/pkg/repository"
)

//go:generate mockgen -destination mocks/mock.go -package service github.com/p12s/library-rest-api/library/pkg/service Authorization,Author,Book

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
