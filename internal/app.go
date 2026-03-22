package app

import (
	"fmt"
	config "gem/internal/configs"
	registry "gem/internal/registry"
	"gem/pkg/db"
	"log"
	"net/http"
	"time"
)

func Start() {

	cfg, err := config.Load()

	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	db, err := db.New(
		cfg.DB.Addr,
		cfg.DB.MaxOpenConns,
		cfg.DB.MaxIdleConns,
		cfg.DB.MaxIdleTime,
	)

	if err != nil {
		log.Fatal("Failed to initiate db: ", err)
	}

	defer db.Close()
	fmt.Printf("database connection pool established")

	r := registry.New(db, cfg)

	appRouter := RegisterRoutes(r)

	server := &http.Server{
		Addr:         cfg.Addr,
		Handler:      appRouter,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	fmt.Printf("🚀 Server starting on %s\n", cfg.Addr)

	err = server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
