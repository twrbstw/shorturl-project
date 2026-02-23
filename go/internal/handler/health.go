package handler

import (
	"net/http"
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
	health := r.Group("/health")
	health.GET("/liveness", h.Liveness)
	health.GET("/readiness", h.Readiness)
}

func (h *HealthHandler) Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "alive",
	})
}

func (h *HealthHandler) Readiness(c *gin.Context) {
	if err := h.service.CheckReadiness(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "not ready",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ready",
	})
}
