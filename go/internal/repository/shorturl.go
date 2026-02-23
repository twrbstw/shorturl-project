package repository

import (
	"context"
	"database/sql"
	"shorturl-service/internal/model"
)

type IShortUrlRepository interface {
	Save(ctx context.Context, code, longURL string) error
	FindByCode(ctx context.Context, code string) (*model.ShortUrlData, error)
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

func (r *shortUrlRepository) FindByCode(ctx context.Context, code string) (*model.ShortUrlData, error) {
	query := `
		SELECT id, code, long_url, created_at, expired_at
		FROM short_urls
		WHERE code = $1
	`

	var data model.ShortUrlData
	err := r.db.QueryRowContext(ctx, query, code).Scan(
		&data.ID,
		&data.Code,
		&data.LongURL,
		&data.CreatedAt,
		&data.ExpiredAt)
	return &data, err
}
