package repository

import (
	"crud/internal/core/interface/repository"
	"crud/internal/lib/db"
	"crud/internal/repository/postgres"
)

type RepositoryManager struct {
	repository.AuthRepository
	repository.BookRepository
	repository.FavouritesRepository
}

func NewRepositoryManager(db *db.Db) RepositoryManager {
	return RepositoryManager{
		postgres.NewRepo(db),
		postgres.NewBookRepo(db),
		postgres.NewFavouritesRepo(db),
	}
}
