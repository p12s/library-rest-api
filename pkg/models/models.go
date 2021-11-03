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
	Id          int      `json:"id" db:"id"`
	Title       string   `json:"title" db:"title"`
	Description string   `json:"description" db:"description"`
	Authors     []Author `json:"authors" db:"authors"`
}

// Author - book author
type Author struct {
	Id        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}

// UpdateListInput - данные для обновление списка
// type UpdateListInput struct {
// 	Title       *string `json:"title"`
// 	Description *string `json:"description"`
// }

// // Validate - валидация при обновлении списка
// func (i UpdateListInput) Validate() error {
// 	if i.Title == nil && i.Description == nil {
// 		return errors.New("update structure has no values")
// 	}

// 	return nil
// }

// // UpdateItemInput - данные для обновления элемента
// type UpdateItemInput struct {
// 	Title       *string `json:"title"`
// 	Description *string `json:"description"`
// 	Done        *bool   `json:"done"`
// }

// // Validate - валидация при обновлении элемента
// func (i UpdateItemInput) Validate() error {
// 	if i.Title == nil && i.Description == nil && i.Done == nil {
// 		return errors.New("update structure has no values")
// 	}

// 	return nil
// }
