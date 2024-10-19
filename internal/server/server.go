package server

import (
	"context"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/Ayomided/prog/internal/config"
	"github.com/Ayomided/prog/internal/handlers"
	"github.com/Ayomided/prog/internal/middleware"
)

func Run(cfg *config.Config, posts, templates fs.FS) error {
	var stopChan chan os.Signal
	templatesFS, err := fs.Sub(templates, "templates")
	if err != nil {
		return err
	}
	postsFS, err := fs.Sub(posts, "posts")
	if err != nil {
		return err
	}
	fileServer := http.FileServer(http.Dir(cfg.StaticPath))
	fileServerOg := http.FileServer(http.Dir(cfg.StaticPathOG))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /static/og-images/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=86400")

		ext := filepath.Ext(r.URL.Path)
		switch ext {
		case ".png":
			w.Header().Set("Content-Type", "image/png")
		case ".jpg", ".jpeg":
			w.Header().Set("Content-Type", "image/jpeg")
		case ".gif":
			w.Header().Set("Content-Type", "image/gif")
		default:
			http.Error(w, "Unsupported file type", http.StatusUnsupportedMediaType)
			return
		}

		http.StripPrefix("/static/og-images/", fileServerOg).ServeHTTP(w, r)
	})
	mux.HandleFunc("GET /sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		http.ServeFile(w, r, cfg.Sitemap)
	})
	mux.HandleFunc("GET /robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/txt")
		http.ServeFile(w, r, cfg.Robots)
	})
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))
	mux.Handle("GET /", handlers.HomeHandler(postsFS, templatesFS))
	mux.Handle("GET /about", handlers.AboutHandler(templatesFS))
	mux.Handle("GET /og-image/{path}", handlers.OGImageHandler(handlers.FileReader{}, postsFS))
	mux.Handle("GET /posts/{slug}", handlers.PostHandler(handlers.FileReader{}, postsFS, templatesFS))
	mux.Handle("GET /rss", handlers.RssHandler(postsFS))

	loggedMux := middleware.Logging(mux)
	corsLoggedMux := middleware.SetupCORS(loggedMux)

	log.Printf("Starting server on :%s\n", cfg.Port)

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: corsLoggedMux,
	}

	// create channel to listen for signals
	stopChan = make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error when running server: %s", err)
		}
	}()

	<-stopChan

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error when shutting down server: %v", err)
		return err
	}
	return nil
}
