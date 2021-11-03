package service

import (
	"github.com/p12s/library-rest-api/pkg/models"
	"github.com/p12s/library-rest-api/pkg/repository"
)

// BookService - just book service
type BookService struct {
	repo       repository.Book
	authorRepo repository.Author // TODO it's realy need me??
}

// NewBookService - constructor
//func NewBookService(repo repository.Book, authorRepo repository.Author) *BookService {
func NewBookService(repo repository.Book) *BookService {
	//return &BookService{repo: repo, authorRepo: authorRepo}
	return &BookService{repo: repo}
}

// Create - create book
func (s *BookService) Create(book models.Book) (int, error) {
	return s.repo.Create(book)
}

// GetAll - getting all books
func (s *BookService) GetAll() ([]models.Book, error) {
	return s.repo.GetAll()
}

// GetById - get book
func (s *BookService) GetById(bookId int) (models.Book, error) {
	return s.repo.GetById(bookId)
}

// Delete - book
func (s *BookService) Delete(bookId int) error {
	return s.repo.Delete(bookId)
}

// Update - book
func (s *BookService) Update(bookId int, book models.Book) error {
	return s.repo.Update(bookId, book)
}
