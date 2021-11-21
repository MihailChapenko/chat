package user

import (
	"encoding/json"
	"github.com/MihailChapenko/chat/internal/models"
	postgres "github.com/MihailChapenko/chat/internal/repository/Postgres/user"
	"github.com/MihailChapenko/chat/pkg/errors"
	"github.com/MihailChapenko/chat/pkg/hasher"
	"github.com/MihailChapenko/chat/pkg/openapi3"
	"github.com/MihailChapenko/chat/pkg/token_generator"
	"io/ioutil"
	"net/http"
	"strconv"
)

//Service user service interface
type Service interface {
	CreateUser(w http.ResponseWriter, r *http.Request) ([]byte, error)
	LoginUser(w http.ResponseWriter, r *http.Request) ([]byte, error)
}

//service describe user service struct
type service struct {
	userRepo postgres.UserRepository
}

//NewService create new user service instance
func NewService() Service {
	ur := postgres.NewUserRepository()

	return &service{
		userRepo: ur,
	}
}

//CreateUser create user service
func (s service) CreateUser(w http.ResponseWriter, r *http.Request) ([]byte, error) {
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

	user, err := s.userRepo.FindByUsername(input.UserName)
	if user.ID != 0 {
		return nil, errors.BadRequest("user with such username is exists")
	}

	pass, err := hasher.HashPassword(input.Password)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	input.Password = pass
	res, err := s.userRepo.Create(models.User{Username: input.UserName, Password: input.Password})

	id := strconv.Itoa(int(res.ID))
	output.Id = &id
	output.UserName = &res.Username
	out, err := json.Marshal(output)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}

	return out, nil
}

//LoginUser login user service
func (s service) LoginUser(w http.ResponseWriter, r *http.Request) ([]byte, error) {
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

	user, err := s.userRepo.FindByUsername(input.UserName)
	if user.ID == 0 {
		return nil, errors.BadRequest("user with such username is not exists")
	}

	if ok := hasher.CheckPassword(input.Password, user.Password); !ok {
		return nil, errors.Unauthorized("wrong password")
	}

	output.Url = token_generator.GenerateSecureToken(30)

	out, err := json.Marshal(output)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}

	return out, nil
}
