package handler

import (
	"shorturl-service/internal/model"
	"shorturl-service/internal/service"

	"github.com/gin-gonic/gin"
)

type ShortUrlHandler struct {
	service service.IShortUrlService
}

func NewShortUrlHandler(s service.IShortUrlService) *ShortUrlHandler {
	return &ShortUrlHandler{
		service: s,
	}
}

func (h *ShortUrlHandler) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/shorten", h.Minimize)
	r.GET("/:id", h.Redirect)
}
func (s *ShortUrlHandler) Minimize(c *gin.Context) {

	var req model.MinimizeUrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.service.Minimize(c.Request.Context(), req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

func (s *ShortUrlHandler) Redirect(c *gin.Context) {

}
