package service

import (
	"context"
	"errors"
	"shorturl-service/internal/model"
	"shorturl-service/internal/repository"
	"shorturl-service/utils"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
)

type IShortUrlService interface {
	Minimize(ctx context.Context, req model.MinimizeUrlRequest) (*model.MinimizeUrlResponse, error)
	Revert(ctx context.Context, req model.RedirectUrlRequest) (*model.RedirectUrlResponse, error)
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
	for range maxAttempts {

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
func (s *shortUrlService) Revert(ctx context.Context, req model.RedirectUrlRequest) (*model.RedirectUrlResponse, error) {
	data, err := s.repo.FindByCode(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	if time.Now().After(data.ExpiredAt) {
		return nil, errors.New("URL_EXPIRED")
	}
	return &model.RedirectUrlResponse{OriginalUrl: data.LongURL}, nil
}
