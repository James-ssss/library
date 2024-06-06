package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/model"
	"crud/internal/lib/db"
	"crud/internal/repository/dbModel"
	"fmt"
)

type _bookRepository struct {
	db *db.Db
}

func NewBookRepo(db *db.Db) repository.BookRepository {
	return _bookRepository{db}
}

func (bookRepository _bookRepository) CreateBook(ctx context.Context, book model.Book) (int, error) {
	bookDb := dbModel.Book(book)
	var id int

	err := bookRepository.db.PgConn.QueryRow(ctx,
		`INSERT INTO public.book(title, description, image, genre, year, pages, author) 
					values ($1,$2,$3,$4, $5, $6, $7) RETURNING id`,
		bookDb.Title,
		bookDb.Description,
		bookDb.ImageURL,
		bookDb.Genre,
		bookDb.Year,
		bookDb.Pages,
		bookDb.Author).Scan(&id)

	return id, err
}

func (bookRepository _bookRepository) GetBook(ctx context.Context, bookId int) (model.Book, error) {
	var book dbModel.Book

	err := bookRepository.db.PgConn.QueryRow(ctx,
		`SELECT p.title, p.description, p.image, p.genre, p.year, p.pages, p.author FROM public.book p WHERE p.id=$1`,
		bookId).Scan(&book.Title, &book.Description, &book.ImageURL, &book.Genre, &book.Year, &book.Pages, &book.Author)

	if err != nil {
		return model.Book{}, fmt.Errorf("ошибка получения книги: %s", err.Error())
	}

	return model.Book(book), nil
}

func (bookRepository _bookRepository) GetBooks(ctx context.Context) ([]model.Book, []int, error) {
	var books []model.Book
	var ids []int

	rows, err := bookRepository.db.PgConn.Query(ctx,
		`SELECT p.id, p.title, p.description, p.image, p.genre, p.year, p.pages, p.author FROM public.book p`)
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка получения списка книг: %s", err.Error())
	}

	for rows.Next() {
		var book dbModel.Book
		var id int

		err := rows.Scan(&id, &book.Title, &book.Description, &book.ImageURL, &book.Genre, &book.Year, &book.Pages, &book.Author)
		if err != nil {
			return nil, nil, fmt.Errorf("ошибка чтения книги: %s", err.Error())
		}

		books = append(books, model.Book(book))
		ids = append(ids, id)
	}

	return books, ids, nil
}

func (bookRepository _bookRepository) UpdateBook(ctx context.Context, book model.Book, bookId int) (model.Book, error) {
	bookDb := dbModel.Book(book)
	var id int

	err := bookRepository.db.PgConn.QueryRow(ctx,
		`UPDATE public.book
				SET title = $1,
    				description = $2,
   					image = $3,
    				genre = $4,
   					year = $5,
   					pages = $6,
   					author = $7
				WHERE id = $8`,
		bookDb.Title,
		bookDb.Description,
		bookDb.ImageURL,
		bookDb.Genre,
		bookDb.Year,
		bookDb.Pages,
		bookDb.Author,
		bookId).Scan(&id)

	return book, err
}

func (bookRepository _bookRepository) DeleteBook(ctx context.Context, bookId int) error {

	err := bookRepository.db.PgConn.QueryRow(ctx,
		`DELETE FROM public.book p WHERE p.id=$1`,
		bookId)

	if err != nil {
		return fmt.Errorf("ошибка получения книги: %s", err)
	}

	return nil
}

func (bookRepository _bookRepository) GetBooksByTitle(ctx context.Context, bookTitle string) ([]model.Book, []int, error) {
	var books []model.Book
	var ids []int

	rows, err := bookRepository.db.PgConn.Query(ctx,
		`SELECT p.id, p.title, p.description, p.image, p.genre, p.year, p.pages, p.author 
					FROM public.book p WHERE LOWER(p.title) LIKE '%' || LOWER($1) || '%'`, bookTitle)

	fmt.Println(rows)

	if err != nil {
		return nil, nil, fmt.Errorf("ошибка получения списка книг: %s", err.Error())
	}

	for rows.Next() {
		var book dbModel.Book
		var id int

		err := rows.Scan(&id, &book.Title, &book.Description, &book.ImageURL, &book.Genre, &book.Year, &book.Pages, &book.Author)
		if err != nil {
			return nil, nil, fmt.Errorf("ошибка чтения книги: %s", err.Error())
		}

		books = append(books, model.Book(book))
		ids = append(ids, id)
	}

	return books, ids, nil
}
