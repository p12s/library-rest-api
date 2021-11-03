package repository

import (
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
func (r *BookPostgres) Create(book models.Book) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	// идем по-списку авторов, сохраняем авторов если таких нет
	// сохраняем книгу
	// сохраняем ид книга-автор
	/*
		var id int
		createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
		row := tx.QueryRow(createListQuery, list.Title, list.Description)
		if err := row.Scan(&id); err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				return 0, rollbackErr
			}
			return 0, err
		}

		createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
		_, err = tx.Exec(createUsersListQuery, userId, id)
		if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				return 0, rollbackErr
			}
			return 0, err
		}
	*/
	return 99, tx.Commit()
}

// GetAll - getting all books
func (r *BookPostgres) GetAll() ([]models.Book, error) {
	var lists []models.Book

	// query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
	// 	todoListsTable, usersListsTable)
	// err := r.db.Select(&lists, query, userId)

	return lists, nil
}

// GetById - getting book
func (r *BookPostgres) GetById(bookId int) (models.Book, error) {
	var book models.Book

	// query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
	// 							INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
	// 	todoListsTable, usersListsTable)
	// err := r.db.Get(&list, query, userId, listId)

	return book, nil
}

// Delete - remove book
func (r *BookPostgres) Delete(bookId int) error {
	// query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
	// 	todoListsTable, usersListsTable)
	// _, err := r.db.Exec(query, userId, listId)

	return nil
}

// Update - удаление СПИСКА
func (r *BookPostgres) Update(bookId int, book models.Book) error {
	// setValues := make([]string, 0)
	// args := make([]interface{}, 0)
	// argId := 1

	// if input.Title != nil {
	// 	setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
	// 	args = append(args, *input.Title)
	// 	argId++
	// }

	// if input.Description != nil {
	// 	setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
	// 	args = append(args, *input.Description)
	// 	argId++
	// }

	// // title=$1
	// // description=$1
	// // title=$1, description=$2
	// setQuery := strings.Join(setValues, ", ")

	// query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
	// 	todoListsTable, setQuery, usersListsTable, argId, argId+1)
	// args = append(args, listId, userId)

	// logrus.Debugf("updateQuery: %s", query)
	// logrus.Debugf("args: %s", args)

	// _, err := r.db.Exec(query, args...)
	return nil
}
