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
func NewBookService(repo repository.Book, authorRepo repository.Author) *BookService {
	return &BookService{repo: repo, authorRepo: authorRepo}
}

// Create - create book
func (s *BookService) CreateBook(book models.Book) (int, error) {
	return s.repo.CreateBook(book)
}

// GetAll - getting all books
func (s *BookService) GetAllBook() ([]models.Book, error) {
	return s.repo.GetAllBook()
}

// GetById - get book
func (s *BookService) GetBookById(bookId int) (models.Book, error) {
	return s.repo.GetBookById(bookId)
}

// Delete - book
func (s *BookService) DeleteBook(bookId int) error {
	return s.repo.DeleteBook(bookId)
}

// Update - book
func (s *BookService) UpdateBook(bookId int, book models.Book) error {
	return s.repo.UpdateBook(bookId, book)
}
