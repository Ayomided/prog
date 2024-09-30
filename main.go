package main

import (
	"log"

	"github.com/Ayomided/prog/internal/config"
	"github.com/Ayomided/prog/internal/server"
)

func main() {
	cfg := config.NewConfig()
	if err := server.Run(cfg); err != nil {
		log.Fatalf("could not run the server: %v", err)
	}
}
