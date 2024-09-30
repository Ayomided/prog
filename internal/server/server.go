package server

import (
	"log"
	"net/http"

	"github.com/Ayomided/prog.git/internal/config"
	"github.com/Ayomided/prog.git/internal/middleware"
)

func Run(cfg *config.Config) error {
	mux := http.NewServeMux()
	mux.Handle("GET /", http.NotFoundHandler())

	loggedMux := middleware.Logging(mux)
	corsLoggedMux := middleware.SetupCORS(loggedMux)

	log.Printf("Starting server on :%s\n", cfg.Port)
	return http.ListenAndServe(":"+cfg.Port, corsLoggedMux)
}
