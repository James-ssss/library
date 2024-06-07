package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/lib/db"
	"database/sql"
	"errors"
	"fmt"
)

type userDB struct {
	Login    string `db:"login"`
	Password string `db:"pass"`
}

type _authRepo struct {
	*db.Db
}

func NewRepo(db *db.Db) repository.AuthRepository {
	return _authRepo{db}
}

func (repo _authRepo) GetUser(ctx context.Context, login, hashPassword string) (string, error) {
	var user userDB

	row := repo.PgConn.QueryRow(ctx, `SELECT * FROM users WHERE login=$1 AND password=$2`, login, hashPassword)

	if err := row.Scan(&user); err != nil {
		return "", fmt.Errorf("не смогли получить юзера: %x", err)
	}

	return login, nil
}

func (repo _authRepo) Register(ctx context.Context, login, hashPassword string) (string, error) {
	var uLogin string
	row := repo.PgConn.QueryRow(ctx, `SELECT login FROM users WHERE login=$1`, login)

	if err := row.Scan(&uLogin); !errors.Is(err, sql.ErrNoRows) && err == nil {
		return "", fmt.Errorf("пользователь с таким логином уже зарегистрирован")
	}

	_, err := repo.PgConn.Exec(
		ctx,
		`INSERT INTO users(login, password) values ($1, $2)`,
		login, hashPassword,
	)

	if err != nil {
		return "", fmt.Errorf("не смогли создать: %x", err)
	}
	fmt.Printf("Register login %s password %s \n", login, hashPassword)

	return login, nil
}
