package service

import (
	"context"
	"errors"
	"shorturl-service/internal/model"
	"shorturl-service/internal/repository"
	"shorturl-service/utils"

	"github.com/jackc/pgx/v5/pgconn"
)

type IShortUrlService interface {
	Minimize(ctx context.Context, req model.MinimizeUrlRequest) (*model.MinimizeUrlResponse, error)
	Redirect() string
}

type shortUrlService struct {
	repo repository.IShortUrlRepository
}

func NewShortUrlService(repo repository.IShortUrlRepository) IShortUrlService {
	return &shortUrlService{repo: repo}
}

// Minimize implements IShortUrlService.
func (s *shortUrlService) Minimize(ctx context.Context, req model.MinimizeUrlRequest) (*model.MinimizeUrlResponse, error) {
	maxAttempts := 3
	for i := 0; i < maxAttempts; i++ {

		code, err := utils.GenerateCode(6)
		if err != nil {
			return nil, err
		}

		err = s.repo.Save(ctx, code, req.OriginalUrl)
		if err == nil {
			return &model.MinimizeUrlResponse{ShortURL: "https://shorten-url/" + code}, nil
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			continue // retry
		}

		return nil, err
	}
	return nil, errors.New("CODE_GENERATION_FAILED")
}

// Redirect implements IShortUrlService.
func (s *shortUrlService) Redirect() string {
	panic("unimplemented")
}
