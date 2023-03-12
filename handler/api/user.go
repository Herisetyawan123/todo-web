package api

import (
	"encoding/json"
	"net/http"
	"todo-web/entity"
	"todo-web/service"
)

type UserApi interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Token(w http.ResponseWriter, r *http.Request)
	Verification(w http.ResponseWriter, r *http.Request)
}

type userApi struct {
	userService service.UserService
}

func NewUserApi(service service.UserService) UserApi {
	return &userApi{service}
}

func (u *userApi) GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(entity.User{ID: 1, Email: "herisetyawan233@gmail.com", Password: "123456", Username: "herisetyawan"})
	return
}

func (u *userApi) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: untuk menghandle register
	json.NewEncoder(w).Encode(entity.User{ID: 1, Email: "herisetyawan233@gmail.com", Password: "123456", Username: "herisetyawan"})
	return
}

func (u *userApi) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: untuk menghandle login
	json.NewEncoder(w).Encode(entity.User{ID: 1, Email: "herisetyawan233@gmail.com", Password: "123456", Username: "herisetyawan"})
	return
}

func (u *userApi) Token(w http.ResponseWriter, r *http.Request) {
	// TODO: untuk menghandle token
	json.NewEncoder(w).Encode(entity.User{ID: 1, Email: "herisetyawan233@gmail.com", Password: "123456", Username: "herisetyawan"})
	return
}

func (u *userApi) Verification(w http.ResponseWriter, r *http.Request) {
	// TODO: untuk menghandle vericitaion
	json.NewEncoder(w).Encode(entity.User{ID: 1, Email: "herisetyawan233@gmail.com", Password: "123456", Username: "herisetyawan"})
	return
}
