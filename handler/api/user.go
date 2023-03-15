package api

import (
	"encoding/json"
	"net/http"
	"todo-web/entity"
	"todo-web/service"
	"todo-web/utils"
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
	var userForm entity.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewResponseError("terjadi kesalah ketika menerjemahkan request body"))
		return
	}

	if userForm.Email == "" || userForm.Password == "" || userForm.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewResponseError("jangan ada form yang kosong"))
		return
	}

	passHash, err := utils.GeneratePassword(userForm.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewResponseError("terjadi kesalah pada server"))
		return
	}

	user := entity.User{
		Email:    userForm.Email,
		Username: userForm.Username,
		Password: passHash,
	}
	newUser, err := u.userService.Register(r.Context(), &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewResponseError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "success", "data": newUser})
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
