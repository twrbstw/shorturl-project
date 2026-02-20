package repository

import (
	"context"
	"database/sql"
)

type IShortUrlRepository interface {
	Save(ctx context.Context, code, longURL string) error
	FindByCode(ctx context.Context, code string) (string, error)
}

type shortUrlRepository struct {
	db *sql.DB
}

func NewShortUrlRepository(db *sql.DB) IShortUrlRepository {
	return &shortUrlRepository{db: db}
}

func (r *shortUrlRepository) Save(ctx context.Context, code, longURL string) error {
	query := `
		INSERT INTO short_urls (code, long_url)
		VALUES ($1, $2)
	`
	_, err := r.db.ExecContext(ctx, query, code, longURL)
	return err
}

func (r *shortUrlRepository) FindByCode(ctx context.Context, code string) (string, error) {
	query := `
		SELECT long_url
		FROM short_urls
		WHERE code = $1
	`

	var longURL string
	err := r.db.QueryRowContext(ctx, query, code).Scan(&longURL)
	return longURL, err
}
