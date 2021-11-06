package service

import (
	"github.com/p12s/library-rest-api/internal/repository"
)

// Author - author commands (not implemented, because is not required)
type Author interface {
}

// AuthorService - just service
type AuthorService struct {
	repo repository.Author
}

// NewAuthorService - constructor
func NewAuthorService(repo repository.Author) *AuthorService {
	return &AuthorService{repo: repo}
}
