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

type Database struct {
	pool *pgxpool.Pool
}

func NewDatabase(pool *pgxpool.Pool) *Database {
	return &Database{
		pool: pool,
	}
}

func (db *Database) Close() {
	db.pool.Close()
}

func (db *Database) Persist(ctx context.Context, url domain.URL) (int64, error) {
	var id int64
	row := db.pool.QueryRow(ctx, "insert into url(longurl) values ($1) returning id", url.URL)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (db *Database) GetIdByUrl(ctx context.Context, url domain.URL) (int64, error) {
	var id int64
	row := db.pool.QueryRow(ctx, "select id from url where longurl = $1", url.URL)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func (db *Database) GetById(ctx context.Context, id int64) (string, error) {
	var url string
	r := db.pool.QueryRow(ctx, "select longurl from url where id = $1", id)
	if err := r.Scan(&url); err != nil {
		return "", err
	}
	return url, nil
}
