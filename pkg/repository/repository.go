package repository

import (
	"github.com/Garagator3000/restBook"
	"github.com/jmoiron/sqlx"
)

type Book interface {
	Create(book restBook.Book) (int, error)
	GetAll() ([]restBook.Book, error)
	GetById(id int) (restBook.Book, error)
	DeleteById(id int) (int, error)
	UpdateById(id int, book restBook.Book) (restBook.Book, error)
}

type Repository struct {
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Book: NewBookPostgres(db),
	}
}
