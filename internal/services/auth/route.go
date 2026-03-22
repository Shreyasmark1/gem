package auth

import (
	middleware "gem/internal/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterAuthRoutes(service *AuthService) *chi.Mux {

	handler := NewAuthHandler(service)
	router := chi.NewRouter()

	router.Post("/sign-up", middleware.ErrorHandler(handler.SignUp))
	router.Post("/verify", middleware.ErrorHandler(handler.VerifyAuth))
	router.Post("/login", middleware.ErrorHandler(handler.Login))

	return router
}
