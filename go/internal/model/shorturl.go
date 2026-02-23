package model

import "time"

type ShortUrlData struct {
	ID        int64     `db:"id"`
	Code      string    `db:"code"`
	LongURL   string    `db:"long_url"`
	CreatedAt time.Time `db:"created_at"`
	ExpiredAt time.Time `db:"expired_at"`
}

type MinimizeUrlRequest struct {
	OriginalUrl string `json:"original_url"`
}

type MinimizeUrlResponse struct {
	ShortURL string `json:"short_url"`
}

type RedirectUrlRequest struct {
	Code string
}

type RedirectUrlResponse struct {
	OriginalUrl string `json:"original_url"`
}
