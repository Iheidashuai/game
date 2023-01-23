package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func NewDB() (*DB, error) {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/game")
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
