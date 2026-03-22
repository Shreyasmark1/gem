package app

import (
	appMiddleware "gem/internal/middlewares"
	"gem/internal/registry"
	"gem/internal/services/auth"
	"gem/internal/services/mine"
	"gem/pkg/response"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterApiRoutes(r *registry.ServiceRegistry) http.Handler {
	apiRouter := chi.NewRouter()

	apiRouter.Use(appMiddleware.JWTMiddleware(r.Config.JWTSecret))

	apiRouter.Mount("/v1/auth", auth.RegisterAuthRoutes(r.AuthService))
	apiRouter.With(appMiddleware.RequireAuth()).Mount("/v1/mines", mine.RegisterMineRoutes(r.MineService))

	return apiRouter

}

func RegisterRoutes(r *registry.ServiceRegistry) http.Handler {

	appRouter := chi.NewRouter()
	appRouter.Use(middleware.Recoverer)

	appRouter.HandleFunc("/", __test__)
	appRouter.Mount("/api", RegisterApiRoutes(r))

	return appRouter
}

func __test__(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	values := r.URL.Query()

	response.Send(w, values)

}
