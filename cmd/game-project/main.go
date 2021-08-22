package main

import (
	"net/http"

	"game-project/internal/adapters/http/handler"
	"game-project/internal/adapters/postgresql"

	"github.com/gorilla/mux"
)

type ApplicationHandler struct {
	UserHandler  *handler.UserHandler
}

func NewApplicationHandler(u *handler.UserHandler) *ApplicationHandler {
	return &ApplicationHandler{
		UserHandler: u,
	}
}

func Router(appHandler *ApplicationHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user", appHandler.UserHandler.Create).Methods("POST")
	r.HandleFunc("/user/{userId}/state", appHandler.UserHandler.UpdateUserState).Methods("PUT")
	r.HandleFunc("/user/{userId}/state", appHandler.UserHandler.LoadUserState).Methods("GET")
	r.HandleFunc("/user/{userId}/friends", appHandler.UserHandler.UpdateUserFriends).Methods("PUT")
	r.HandleFunc("/user/{userId}/friends", appHandler.UserHandler.ListUserFriends).Methods("GET")

	return r
}

func main() {
	pool := postgresql.CreatePool()
	userRepository := postgresql.NewUserRepository(pool)

	appHandler := NewApplicationHandler(handler.NewUserHandler(userRepository))
	router := Router(appHandler)

	http.ListenAndServe(":8080", router)

}