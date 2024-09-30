package server

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/Ayomided/prog/internal/config"
	"github.com/Ayomided/prog/internal/handlers"
	"github.com/Ayomided/prog/internal/middleware"
)

func Run(cfg *config.Config, posts, templates fs.FS) error {
	mux := http.NewServeMux()
	mux.Handle("GET /", handlers.HomeHandler(templates))
	mux.Handle("GET /about", handlers.AboutHandler(templates))
	mux.Handle("GET /posts/{slug}", handlers.PostHandler(handlers.FileReader{}, posts, templates))
	mux.Handle("GET /rss", handlers.RssHandler())

	loggedMux := middleware.Logging(mux)
	corsLoggedMux := middleware.SetupCORS(loggedMux)

	log.Printf("Starting server on :%s\n", cfg.Port)
	return http.ListenAndServe(":"+cfg.Port, corsLoggedMux)
}
