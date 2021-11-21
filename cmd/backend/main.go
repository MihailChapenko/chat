package main

import (
	"github.com/MihailChapenko/chat/config"
	"github.com/MihailChapenko/chat/db"
	"github.com/MihailChapenko/chat/pkg/handlers"
	"github.com/MihailChapenko/chat/pkg/openapi3"
	"log"
	"net/http"
	"os"
)

func main() {
	config.Init("../../config/dev.yaml")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cfg := config.Get()

	db.Init(cfg.DB)
	h := openapi3.Handler(handlers.NewHandler())

	log.Println("Starting server on port :8080")
	http.ListenAndServe(":"+cfg.Server.Port, h)
}
