package main

import (
	"github.com/MihailChapenko/chat/config"
	"github.com/MihailChapenko/chat/db"
	"github.com/MihailChapenko/chat/pkg/handlers"
	"github.com/MihailChapenko/chat/pkg/logger"
	"github.com/MihailChapenko/chat/pkg/openapi3"
	"log"
	"net/http"
	"os"
)

func main() {
	config.Init("../../config/dev.yaml")
	cfg := config.Get()
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Server.Port
	}

	logger.Init()
	db.Init(cfg.DB)
	h := openapi3.Handler(handlers.NewHandler())

	log.Println("Starting server on port :8080")
	http.ListenAndServe(":"+cfg.Server.Port, h)
}
