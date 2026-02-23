package handler

import (
	"net/http"
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
	r.GET("/:code", h.Revert)
}
func (s *ShortUrlHandler) Minimize(c *gin.Context) {

	var req model.MinimizeUrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.service.Minimize(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (s *ShortUrlHandler) Revert(c *gin.Context) {
	code := c.Param("code")

	resp, err := s.service.Revert(c.Request.Context(), model.RedirectUrlRequest{Code: code})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
