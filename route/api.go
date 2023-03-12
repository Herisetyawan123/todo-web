package route

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"todo-web/handler"
	"todo-web/handler/api"
	"todo-web/middleware"
	"todo-web/repository"
	"todo-web/service"
)

func Api(mux *http.ServeMux, db *gorm.DB) *http.ServeMux {

	// user
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserApi(userService)

	// route

	// users
	MuxRoute(mux, Get, PathApi("/users"), userHandler.GetUser)
	MuxRoute(mux, Get, PathApi("/users/token"), userHandler.Token)
	MuxRoute(mux, Get, PathApi("/users/verification"), userHandler.Verification)
	MuxRoute(mux, Get, PathApi("/auth/register"), userHandler.Register)
	MuxRoute(mux, Get, PathApi("/auth/login"), userHandler.Login)

	MuxRoute(mux, Get, PathApi("/"), handler.Welcome)
	return mux
}

func MuxRoute(mux *http.ServeMux, method Method, path string, handler func(w http.ResponseWriter, r *http.Request)) {
	var handle http.Handler
	if method.String() == "GET" {
		handle = middleware.Get(http.HandlerFunc(handler))
	} else {
		handle = middleware.Get(http.HandlerFunc(handler))
	}
	fmt.Printf("[%s]: http://localhost:8080%s \n", method, path)
	mux.Handle(path, handle)
}
