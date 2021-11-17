package user

import (
	"encoding/json"
	"github.com/MihailChapenko/chat/pkg/errors"
	"github.com/MihailChapenko/chat/pkg/hasher"
	"github.com/MihailChapenko/chat/pkg/openapi3"
	"io/ioutil"
	"net/http"
)

var users = make(map[string]openapi3.CreateUserRequest)

type Service interface{}

type service struct{}

func NewService() Service {
	return &service{}
}

func CreateUser(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	var input openapi3.CreateUserRequest
	var output openapi3.CreateUserResponse

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &input)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}

	if _, ok := users[input.UserName]; ok {
		return nil, errors.BadRequest("such user is exists")
	}

	pass, err := hasher.HashPassword(input.Password)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	input.Password = pass
	key := input.UserName
	users[key] = input

	output.Id = &input.UserName
	output.UserName = &input.UserName
	out, err := json.Marshal(output)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}

	return out, nil
}

func LoginUser(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	var input openapi3.LoginUserRequest
	var output openapi3.LoginUserResonse

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &input)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}

	v, ok := users[input.UserName]
	if !ok {
		return nil, errors.BadRequest("wrong username")
	}

	if ok := hasher.CheckPassword(input.Password, v.Password); !ok {
		return nil, errors.Unauthorized("wrong password")
	}

	output.Url = "google.com"

	out, err := json.Marshal(output)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}

	return out, nil
}
