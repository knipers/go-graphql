package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Author struct {
	db   *sql.DB
	ID   string
	Name string
	//Books []*Book
}

func NewAuthor(db *sql.DB) *Author {
	return &Author{db: db}
}

func (a *Author) Create(Name string) (Author, error) {
	id := uuid.New().String()
	_, err := a.db.Exec("INSERT INTO authors (id, name) VALUES ($1, $2)", id, Name)

	if err != nil {
		return Author{}, err
	}

	return Author{ID: id, Name: Name}, nil
}
