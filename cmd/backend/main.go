package main

import (
	"github.com/MihailChapenko/chat/pkg/handlers"
	"github.com/MihailChapenko/chat/pkg/openapi3"
	"log"
	"net/http"
)

func main() {
	port := "8080"
	h := openapi3.Handler(handlers.ChatServer{})

	log.Println("Starting server on port :8080")
	http.ListenAndServe(process.env.PORT || port, h)
}
