package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p12s/library-rest-api/pkg/models"
)

// BookPostgres - repo
type BookPostgres struct {
	db *sqlx.DB
}

// NewBookPostgres - constructor
func NewBookPostgres(db *sqlx.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

// Create - book creating
func (r *BookPostgres) CreateBook(book models.Book) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("create book transaction begin: %w", err)
	}

	var authorIds []int
	for _, author := range book.Authors {
		var authorId int
		query := fmt.Sprintf(`SELECT id FROM %s WHERE first_name = $1 AND second_name = $2`, authorTable)
		err := r.db.Get(&authorId, query, author.FirstName, author.SecondName)

		if err == nil {
			authorIds = append(authorIds, authorId)
			continue
		}

		if err == sql.ErrNoRows {
			createAuthorQuery := fmt.Sprintf(`INSERT INTO %s (first_name, second_name) VALUES ($1, $2) RETURNING id`, authorTable)
			row := tx.QueryRow(createAuthorQuery, author.FirstName, author.SecondName)
			if err := row.Scan(&authorId); err != nil {
				rollbackErr := tx.Rollback()
				if rollbackErr != nil {
					return 0, fmt.Errorf("create author rollback: %w", rollbackErr)
				}
				return 0, fmt.Errorf("create author: %w", err)
			}
			authorIds = append(authorIds, authorId)
			continue
		}

		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				return 0, fmt.Errorf("create author rollback: %w", rollbackErr)
			}
			return 0, fmt.Errorf("create author: %w", err)
		}
	}

	var bookId int
	createBookQuery := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", bookTable)
	row := tx.QueryRow(createBookQuery, book.Title)
	if err := row.Scan(&bookId); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return 0, fmt.Errorf("create book rollback: %w", rollbackErr)
		}
		return 0, fmt.Errorf("create book: %w", err)
	}

	createBookAuthorQuery := fmt.Sprintf("INSERT INTO %s (book_id, author_id) VALUES ", bookAuthorTable)
	for i, authorId := range authorIds {
		createBookAuthorQuery += fmt.Sprintf("(%d, %d)", bookId, authorId)
		if i < len(authorIds)-1 {
			createBookAuthorQuery += ","
		}
	}

	_, err = tx.Exec(createBookAuthorQuery)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return 0, fmt.Errorf("create book-author rollback: %w", rollbackErr)
		}
		return 0, fmt.Errorf("create book-author record: %w", err)
	}

	return bookId, tx.Commit()
}

// GetAll - getting all books
func (r *BookPostgres) GetAllBook() ([]models.Book, error) {
	var books []models.Book

	query := fmt.Sprintf(`SELECT b.id, b.title
		FROM %s b INNER JOIN %s ba on ba.book_id = b.id 
		INNER JOIN %s a on a.id = ba.author_id
		GROUP BY b.id, b.title
		ORDER BY b.id ASC`,
		bookTable, bookAuthorTable, authorTable)
	err := r.db.Select(&books, query)
	if err != nil {
		return []models.Book{}, fmt.Errorf("get all books: %w", err)
	}

	return books, nil
}

// GetById - getting book
func (r *BookPostgres) GetBookById(bookId int) (models.Book, error) {
	var book models.Book

	query := fmt.Sprintf(`SELECT id, title FROM %s WHERE id = $1`, bookTable)
	err := r.db.Get(&book, query, bookId)
	if err != nil {
		return book, fmt.Errorf("get book by id: %w", err)
	}

	return book, nil
}

// Delete - remove book
func (r *BookPostgres) DeleteBook(bookId int) error {
	query := fmt.Sprintf(`DELETE FROM %s b USING %s ba WHERE b.id = ba.book_id AND b.id=$1`,
		bookTable, bookAuthorTable)
	_, err := r.db.Exec(query, bookId)
	if err != nil {
		return fmt.Errorf("delete book by id: %w", err)
	}

	return nil
}

// Update - change book
func (r *BookPostgres) UpdateBook(bookId int, book models.Book) error {
	query := fmt.Sprintf(`UPDATE %s SET title='%s' WHERE id = $1`, bookTable, book.Title)

	_, err := r.db.Exec(query, bookId)
	if err != nil {
		return fmt.Errorf("update book by id: %w", err)
	}

	return nil
}
