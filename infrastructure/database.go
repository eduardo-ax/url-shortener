package infrastructure

import (
	"context"
	"log"

	"github.com/eduardo-ax/url-shortener/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool() *pgxpool.Pool {
	dbPool, err := pgxpool.New(context.Background(), "postgres://db_user:db_password@localhost:5432/urlshortener?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return dbPool
}

func NewDatabase(pool *pgxpool.Pool) *Database {
	return &Database{
		pool: pool,
	}
}

type Database struct {
	pool *pgxpool.Pool
}

func (db *Database) Persist(ctx context.Context, url domain.URL) (int64, error) {
	var id int64
	row := db.pool.QueryRow(ctx, "insert into url(url) values ($1) returning id", url.Url)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
