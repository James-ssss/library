package service

import (
	"context"
	"crud/internal/core/model"
)

type AuthService interface {
	Register(ctx context.Context, login, password, role string) (string, error)
	GenerateToken(ctx context.Context, login, password string) (string, error)
}

type BookService interface {
	CreateBook(ctx context.Context, book model.Book) (int, error)
	DeleteBook(ctx context.Context, bookId int) error
	GetBook(ctx context.Context, bookId int) (model.Book, error)
	GetBooks(ctx context.Context) ([]model.Book, []int, error)
	UpdateBook(ctx context.Context, book model.Book, bookId int) (model.Book, error)

	GetBooksByCondition(ctx context.Context, num int, bookGenre string) ([]model.Book, []int, error)
}

type FavouritesService interface {
	CreateFavourite(ctx context.Context, login string, bookId int) (int, error)
	GetFavourites(ctx context.Context, login string) ([]model.Book, []int, error)
	DeleteFavourite(ctx context.Context, bookId int) error
}
