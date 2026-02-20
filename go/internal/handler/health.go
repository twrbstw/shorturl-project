package handler

import (
	"shorturl-service/internal/service"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	service service.IHealthService
}

func NewHealthHandler(s service.IHealthService) *HealthHandler {
	return &HealthHandler{
		service: s,
	}
}

func (h *HealthHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/health/liveness", h.Liveness)
	r.GET("/health/readiness", h.Readiness)
}

func (h *HealthHandler) Liveness(c *gin.Context) {

}

func (h *HealthHandler) Readiness(c *gin.Context) {

}
