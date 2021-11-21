package handlers

import (
	"github.com/MihailChapenko/chat/internal/services/user"
	"github.com/MihailChapenko/chat/pkg/errors"
	"net/http"
)

//Handler user handler interface
type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}

//ChatServer describe chet server struct
type ChatServer struct {
	userService user.Service
}

//NewHandler create new handler instance
func NewHandler() Handler {
	us := user.NewService()

	return &ChatServer{
		userService: us,
	}
}

//CreateUser create user handler
func (c ChatServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	res, err := c.userService.CreateUser(w, r)
	if err != nil {
		http.Error(w, err.Error(), err.(errors.ErrorResponse).Status)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

//LoginUser login user handler
func (c ChatServer) LoginUser(w http.ResponseWriter, r *http.Request) {
	res, err := c.userService.LoginUser(w, r)
	if err != nil {
		http.Error(w, err.Error(), err.(errors.ErrorResponse).Status)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(res)
}
