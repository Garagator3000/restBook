package repository

import (
	"fmt"
	"github.com/Garagator3000/restBook"
	"github.com/jmoiron/sqlx"
)

type BookPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (r *BookPostgres) Create(book restBook.Book) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, author, price, year_of_publish, count) values ($1, $2, $3, $4, $5) RETURNING id", booksTable)
	row := r.db.QueryRow(query, book.Title, book.Author, book.Price, book.YearOfPublish, book.Count)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *BookPostgres) GetAll() ([]restBook.Book, error) {
	var books []restBook.Book
	query := fmt.Sprintf("SELECT * FROM %s", booksTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return []restBook.Book{}, err
	}
	for rows.Next() {
		var tmp restBook.Book
		err := rows.Scan(&tmp.Id, &tmp.Title, &tmp.Author, &tmp.Price, &tmp.YearOfPublish, &tmp.Count)
		if err != nil {
			return []restBook.Book{}, err
		}
		books = append(books, tmp)
	}
	return books, nil
}

func (r *BookPostgres) GetById(id int) (restBook.Book, error) {
	var book restBook.Book
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", booksTable)
	err := r.db.Get(&book, query, id)
	if err != nil {
		return restBook.Book{}, err
	}
	return book, nil
}

func (r *BookPostgres) DeleteById(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", booksTable)
	res, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	count, _ := res.RowsAffected()
	return int(count), nil
}

func (r *BookPostgres) UpdateById(id int, book restBook.Book) (restBook.Book, error) {
	query := fmt.Sprintf(`UPDATE %s SET
								title = $1, author = $2, price = $3, year_of_publish = $4, count = $5
								WHERE id = $6`, booksTable)

	_, err := r.db.Exec(query, book.Title, book.Author, book.Price, book.YearOfPublish, book.Count, id)
	if err != nil {
		return restBook.Book{}, err
	}
	newBook, err := r.GetById(id)
	if err != nil {
		return restBook.Book{}, err
	}
	return newBook, nil
}
