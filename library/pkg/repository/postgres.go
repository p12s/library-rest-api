package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/p12s/library-rest-api/library/pkg/config"
)

const (
	usersTable      = "users"
	bookTable       = "book"
	authorTable     = "author"
	bookAuthorTable = "book_author"
)

// NewPostgresDB - open connect and ping trying
func NewPostgresDB(cfg config.Config) (*sqlx.DB, error) {
	driver := cfg.Db.Driver
	db, err := sqlx.Open(driver, fmt.Sprintf("host='%s' port=%d user=%s dbname='%s' password='%s' sslmode=%s",
		cfg.Db.Host, cfg.Db.Port, cfg.Db.User, cfg.Db.Name, cfg.Db.Password, cfg.Db.SslMode))
	if err != nil {
		fmt.Println("open err: ", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("ping err: ", err.Error())
		return nil, err
	}

	return db, nil
}
