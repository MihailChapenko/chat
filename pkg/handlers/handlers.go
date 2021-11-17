package handlers

import (
	"github.com/MihailChapenko/chat/internal/services/user"
	"github.com/MihailChapenko/chat/pkg/errors"
	"net/http"
)

type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}

type ChatServer struct {
	userService user.Service
}

func NewHandler() Handler {
	us := user.NewService()

	return &ChatServer{
		userService: us,
	}
}

func (c ChatServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	res, err := user.CreateUser(w, r)
	if err != nil {
		http.Error(w, err.Error(), err.(errors.ErrorResponse).Status)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func (c ChatServer) LoginUser(w http.ResponseWriter, r *http.Request) {
	res, err := user.LoginUser(w, r)
	if err != nil {
		http.Error(w, err.Error(), err.(errors.ErrorResponse).Status)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(res)
}
