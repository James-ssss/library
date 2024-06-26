package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

const connection = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s application_name=%s"

type Db struct {
	PgConn *pgxpool.Pool
}

func New(ctx context.Context) *Db {

	connectionString := fmt.Sprintf(connection, "localhost", "5432", "postgres", "laugh3425", "postgres", "disable")

	conn, err := pgxpool.New(ctx, connectionString)

	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping(ctx)

	if err != nil {
		fmt.Println("Ошибка пинга:", err)
		log.Fatal(err)
	}

	fmt.Println("Подключено")

	return &Db{PgConn: conn}
}
