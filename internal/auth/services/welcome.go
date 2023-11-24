package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) Welcome() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	}
}