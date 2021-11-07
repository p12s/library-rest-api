package models

// User - just a user
type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Book - just a book
type Book struct {
	Id      int      `json:"id" db:"id"`
	Title   string   `json:"title" db:"title"`
	Authors []Author `json:"authors,omitempty" db:"authors"`
}

// Author - just a author
type Author struct {
	Id         int    `json:"id" db:"id"`
	FirstName  string `json:"first_name" db:"first_name"`
	SecondName string `json:"second_name" db:"second_name"`
}
