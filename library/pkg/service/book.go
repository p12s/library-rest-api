package service

import (
	"github.com/p12s/library-rest-api/library/pkg/models"
	"github.com/p12s/library-rest-api/library/pkg/repository"
)

// Book - book commands
type Book interface {
	CreateBook(book models.Book) (int, error)
	GetAllBook() ([]models.Book, error)
	GetBookById(bookId int) (models.Book, error)
	DeleteBook(bookId int) error
	UpdateBook(bookId int, book models.Book) error
}

// BookService - just book service
type BookService struct {
	repo       repository.Book
	authorRepo repository.Author
}

// NewBookService - constructor
func NewBookService(repo repository.Book, authorRepo repository.Author) *BookService {
	return &BookService{repo: repo, authorRepo: authorRepo}
}

// CreateBook - create book
func (s *BookService) CreateBook(book models.Book) (int, error) {
	// TODO in the right way I would do smt like this,
	// so that in each file was only access to its own table:
	/*
		authorIds, err := s.authorRepo.CreateOrGetAuthors(book.Authors)
		...
		s.repo.CreateBook(book.Title, authorIds)
	*/

	// but now, for simplicity, saving to adjacent tables is in one place:
	return s.repo.CreateBook(book)
}

// GetAllBook - getting all books
func (s *BookService) GetAllBook() ([]models.Book, error) {
	// TODO without adjacent authors yet
	return s.repo.GetAllBook()
}

// GetBookById - get book
func (s *BookService) GetBookById(bookId int) (models.Book, error) {
	// TODO without adjacent authors yet
	return s.repo.GetBookById(bookId)
}

// DeleteBook - book
func (s *BookService) DeleteBook(bookId int) error {
	return s.repo.DeleteBook(bookId)
}

// UpdateBook - book
func (s *BookService) UpdateBook(bookId int, book models.Book) error {
	// TODO without adjacent authors yet
	return s.repo.UpdateBook(bookId, book)
}
