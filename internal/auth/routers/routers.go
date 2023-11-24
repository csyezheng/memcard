package routers

import (
	"github.com/csyezheng/memcard/internal/auth/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(service *services.Service) http.Handler {
	router := gin.Default()
	router.GET("/", service.Welcome())
	return router
}