package main

import (
	"github.com/MihailChapenko/chat/pkg/handlers"
	"github.com/MihailChapenko/chat/pkg/openapi3"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	h := openapi3.Handler(handlers.ChatServer{})

	log.Println("Starting server on port :8080")
	http.ListenAndServe(":"+port, h)
}
