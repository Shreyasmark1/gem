package registry

import (
	"database/sql"
	config "gem/internal/configs"
	"gem/internal/services/auth"
	"gem/internal/services/mine"
	"gem/internal/services/user"
	"gem/internal/store"
)

type ServiceRegistry struct {
	Store       *store.Store
	UserService *user.UserService
	AuthService *auth.AuthService
	MineService *mine.MineService
	Config      *config.Config
}

func New(db *sql.DB, cfg *config.Config) *ServiceRegistry {
	store := store.NewStore(db)

	userSvc := user.NewUserService(store)

	authSvc := auth.NewAuthService(store, userSvc, cfg.JWTSecret)

	mineSvc := mine.NewMineService(store)

	return &ServiceRegistry{
		Store:       store,
		UserService: userSvc,
		AuthService: authSvc,
		MineService: mineSvc,
		Config:      cfg,
	}
}
