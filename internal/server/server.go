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
	templatesFS, err := fs.Sub(templates, "templates")
	if err != nil {
		return err
	}
	postsFS, err := fs.Sub(posts, "posts")
	if err != nil {
		return err
	}
	mux := http.NewServeMux()
	mux.Handle("GET /", handlers.HomeHandler(postsFS, templatesFS))
	mux.Handle("GET /about", handlers.AboutHandler(templatesFS))
	mux.Handle("GET /posts/{slug}", handlers.PostHandler(handlers.FileReader{}, postsFS, templatesFS))
	mux.Handle("GET /rss", handlers.RssHandler(postsFS))

	loggedMux := middleware.Logging(mux)
	corsLoggedMux := middleware.SetupCORS(loggedMux)

	log.Printf("Starting server on :%s\n", cfg.Port)
	return http.ListenAndServe(":"+cfg.Port, corsLoggedMux)
}
