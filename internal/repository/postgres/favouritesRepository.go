package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/model"
	"crud/internal/lib/db"
	"crud/internal/repository/dbModel"
	"fmt"
)

type _favouritesRepository struct {
	db *db.Db
}

func NewFavouritesRepo(db *db.Db) repository.FavouritesRepository {
	return _favouritesRepository{db}
}

func (favouritesRepository _favouritesRepository) GetUserIdByLogin(ctx context.Context, login string) (int, error) {
	var id int

	err := favouritesRepository.db.PgConn.QueryRow(ctx,
		`SELECT u.id FROM users u WHERE u.login = $1`,
		login).Scan(&id)

	return id, err
}

func (favouritesRepository _favouritesRepository) CreateFavourite(ctx context.Context, login string, bookId int) (int, error) {
	var id int
	userId, err := favouritesRepository.GetUserIdByLogin(ctx, login)

	err = favouritesRepository.db.PgConn.QueryRow(ctx,
		`INSERT INTO favourites(user_id, book_id) 
					values ($1,$2) RETURNING book_id`,
		userId,
		bookId).Scan(&id)

	return id, err
}

func (favouritesRepository _favouritesRepository) GetFavourites(ctx context.Context, login string) ([]model.Book, []int, error) {
	var favourites []model.Book
	var ids []int
	var userId, err = favouritesRepository.GetUserIdByLogin(ctx, login)

	rows, err := favouritesRepository.db.PgConn.Query(ctx,
		`SELECT *
					FROM books b
					WHERE b.id in (SELECT f.book_id FROM favourites f WHERE f.user_id = $1)`,
		userId)
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка получения списка избранных книг: %s", err.Error())
	}

	for rows.Next() {
		var favourite dbModel.Book
		var id int

		err := rows.Scan(&id, &favourite.Title, &favourite.Description, &favourite.ImageURL,
			&favourite.Genre, &favourite.Year, &favourite.Pages, &favourite.Author)
		if err != nil {
			return nil, nil, fmt.Errorf("ошибка чтения избранного: %s", err.Error())
		}

		favourites = append(favourites, model.Book(favourite))
		ids = append(ids, id)
	}

	return favourites, ids, nil
}

func (favouritesRepository _favouritesRepository) DeleteFavourite(ctx context.Context, favouriteId int) error {

	err := favouritesRepository.db.PgConn.QueryRow(ctx,
		`DELETE FROM favourites f WHERE f.id=$1`,
		favouriteId)

	if err != nil {
		return fmt.Errorf("ошибка получения избранной книги: %s", err)
	}

	return nil
}
