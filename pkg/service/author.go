package service

import (
	"github.com/p12s/library-rest-api/pkg/repository"
)

// AuthorService - just service
type AuthorService struct {
	repo repository.Author
}

// NewAuthorService - constructor
func NewAuthorService(repo repository.Author) *AuthorService {
	return &AuthorService{repo: repo}
}
