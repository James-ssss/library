package service

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"errors"
	"log/slog"
)

type _favouritesService struct {
	repo repository.FavouritesRepository
}

func NewFavouritesService(repo repository.FavouritesRepository) service.FavouritesService {
	return _favouritesService{repo: repo}
}

func (favouritesService _favouritesService) CreateFavourite(ctx context.Context, login string, bookId int) (int, error) {
	id, err := favouritesService.repo.CreateFavourite(ctx, login, bookId)

	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("ошибка создания книги")
	}

	return id, nil
}

func (favouritesService _favouritesService) GetFavourites(ctx context.Context, login string) ([]model.Book, []int, error) {
	books, ids, err := favouritesService.repo.GetFavourites(ctx, login)

	if err != nil {
		slog.Error(err.Error())
		return nil, nil, errors.New("ошибка вывода книг")
	}

	return books, ids, nil
}

func (favouritesService _favouritesService) DeleteFavourite(ctx context.Context, bookId int) error {
	err := favouritesService.repo.DeleteFavourite(ctx, bookId)

	if err != nil {
		slog.Error(err.Error())
		return errors.New("ошибка удаления книги")
	}
	return nil
}
