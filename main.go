package main

import (
	"embed"
	"log"

	"github.com/Ayomided/prog/internal/config"
	"github.com/Ayomided/prog/internal/server"
)

//go:embed posts
var posts embed.FS

//go:embed templates/*
var templates embed.FS

func main() {
	cfg := config.NewConfig()
	if err := server.Run(cfg, posts, templates); err != nil {
		log.Fatalf("could not run the server: %v", err)
	}
}
