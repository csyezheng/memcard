package routers

import (
	"github.com/csyezheng/memcard/internal/auth/services"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterRoutes(service *services.Service) http.Handler {
	r := chi.NewRouter()
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", service.Signup)
		r.Post("/login", service.Login)
	})
	return r
}
