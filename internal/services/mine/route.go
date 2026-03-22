package mine

import (
	middleware "gem/internal/middlewares"

	"github.com/go-chi/chi/v5"
)

func RegisterMineRoutes(service *MineService) *chi.Mux {

	handler := NewMinehHandler(service)
	router := chi.NewRouter()

	router.Post("", middleware.ErrorHandler(handler.SaveMine))

	return router
}
