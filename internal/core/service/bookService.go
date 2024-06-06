package service

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"errors"
	"log/slog"
)

type _bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) service.BookService {
	return _bookService{repo: repo}
}

func (bookService _bookService) CreateBook(ctx context.Context, book model.Book) (int, error) {
	id, err := bookService.repo.CreateBook(ctx, book)

	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("ошибка создания книги")
	}

	return id, nil
}

func (bookService _bookService) GetBook(ctx context.Context, bookId int) (model.Book, error) {
	book, err := bookService.repo.GetBook(ctx, bookId)

	if err != nil {
		slog.Error(err.Error())
		return book, errors.New("ошибка вывод книги")
	}

	return book, nil
}

func (bookService _bookService) GetBooks(ctx context.Context) ([]model.Book, []int, error) {
	books, ids, err := bookService.repo.GetBooks(ctx)

	if err != nil {
		slog.Error(err.Error())
		return nil, nil, errors.New("ошибка вывода книг")
	}

	return books, ids, nil
}

func (bookService _bookService) DeleteBook(ctx context.Context, bookId int) error {
	err := bookService.repo.DeleteBook(ctx, bookId)

	if err != nil {
		slog.Error(err.Error())
		return errors.New("ошибка удаления книги")
	}
	return nil
}

func (bookService _bookService) UpdateBook(ctx context.Context, book model.Book, bookId int) (model.Book, error) {
	uBook, err := bookService.repo.UpdateBook(ctx, book, bookId)

	if err != nil {
		slog.Error(err.Error())
		return uBook, errors.New("ошибка обновления книги")
	}

	return uBook, nil
}

func (bookService _bookService) GetBooksByCondition(ctx context.Context, num int, bookCondition string) ([]model.Book, []int, error) {
	books, ids, err := bookService.repo.GetBooksByCondition(ctx, num, bookCondition)

	if err != nil {
		slog.Error(err.Error())
		return nil, nil, errors.New("ошибка вывода книг")
	}

	return books, ids, nil
}
