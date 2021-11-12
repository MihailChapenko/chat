package handlers

import (
	"encoding/json"
	"github.com/MihailChapenko/chat/pkg/openapi3"
	"io/ioutil"
	"net/http"
)

var users = make(map[string]openapi3.CreateUserRequest)

type ChatServer struct{}

func (c ChatServer) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input openapi3.CreateUserRequest
	var output openapi3.CreateUserResponse

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &input)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	key := input.UserName
	users[key] = input

	output.Id = &input.UserName
	output.UserName = &input.UserName
	out, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(out)
}

func (c ChatServer) LoginUser(w http.ResponseWriter, r *http.Request) {
	var input openapi3.LoginUserRequest
	var output openapi3.LoginUserResonse

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &input)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if _, ok := users[input.UserName]; !ok {
		http.Error(w, "wrong username", 403)
		return
	}
	output.Url = "google.com"

	out, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(out)
}
