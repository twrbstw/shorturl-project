package router

import (
	"shorturl-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h handler.Handler) {
	h.HealthHandler.RegisterRoutes(r)
	h.ShortUrlHandler.RegisterRoutes(r)
}
