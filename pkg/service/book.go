package service

import (
	"github.com/Garagator3000/restBook"
	"github.com/Garagator3000/restBook/pkg/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (service *BookService) Create(book restBook.Book) (int, error) {
	return service.repo.Create(book)
}

func (service *BookService) GetAll() ([]restBook.Book, error) {
	return service.repo.GetAll()
}

func (service *BookService) GetById(id int) (restBook.Book, error) {
	return service.repo.GetById(id)
}

func (service *BookService) DeleteById(id int) (int, error) {
	return service.repo.DeleteById(id)
}

func (service *BookService) UpdateById(id int, book restBook.Book) (restBook.Book, error) {
	return service.repo.UpdateById(id, book)
}
