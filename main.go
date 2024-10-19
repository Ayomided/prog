package main

import (
	"embed"
	"html/template"
	"log"

	"github.com/Ayomided/prog/internal/config"
	"github.com/Ayomided/prog/internal/server"
	"github.com/Ayomided/prog/internal/utils"
)

//go:embed posts
var posts embed.FS

//go:embed templates/*
var templates embed.FS

type Home struct {
	OGMeta   template.HTML
	Articles []utils.Post
}

func main() {
	if err := utils.GenerateSitemap(); err != nil {
		panic(err)
	}
	cfg := config.NewConfig()
	if err := server.Run(cfg, posts, templates); err != nil {
		log.Fatalf("could not run the server: %v", err)
	}
}
