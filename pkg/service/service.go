package service

import (
	"github.com/Garagator3000/restBook"
	"github.com/Garagator3000/restBook/pkg/repository"
)

type Book interface {
	Create(book restBook.Book) (int, error)
	GetAll() ([]restBook.Book, error)
	GetById(id int) (restBook.Book, error)
	DeleteById(id int) (int, error)
	UpdateById(id int, book restBook.Book) (restBook.Book, error)
}

type Service struct {
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repos),
	}
}
