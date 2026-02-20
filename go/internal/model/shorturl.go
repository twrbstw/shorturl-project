package model

type MinimizeUrlRequest struct {
	OriginalUrl string `json:"original_url"`
}

type MinimizeUrlResponse struct {
	ShortURL string `json:"short_url"`
}
