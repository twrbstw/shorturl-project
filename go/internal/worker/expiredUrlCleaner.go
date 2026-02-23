package worker

import (
	"context"
	"shorturl-service/internal/repository"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type ExpiredUrlCleaner struct {
	repo repository.IShortUrlRepository
}

func NewExpiredUrlCleaner(repo repository.IShortUrlRepository) Worker {
	return &ExpiredUrlCleaner{repo: repo}
}

// process implements IExpiredUrlCleaner.
func (w *ExpiredUrlCleaner) Process(ctx context.Context) {
	for {
		log.Info("Deleting expired url...")
		w.repo.DeleteExpired(ctx)

		time.Sleep(5 * time.Minute)
	}

}
