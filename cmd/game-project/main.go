package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"game-project/internal/adapters/http/handler"
	"game-project/internal/adapters/postgresql"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
)

type ApplicationHandler struct {
	UserHandler  handler.UserHandler
}

func NewApplicationHandler(u handler.UserHandler) ApplicationHandler {
	return ApplicationHandler{
		UserHandler: u,
	}
}

func Router(appHandler ApplicationHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user", appHandler.UserHandler.List).Methods("GET")
	r.HandleFunc("/user", appHandler.UserHandler.Create).Methods("POST")
	r.HandleFunc("/user/{userId}/state", appHandler.UserHandler.UpdateUserState).Methods("PUT")
	r.HandleFunc("/user/{userId}/state", appHandler.UserHandler.LoadUserState).Methods("GET")
	r.HandleFunc("/user/{userId}/friends", appHandler.UserHandler.UpdateUserFriends).Methods("PUT")
	r.HandleFunc("/user/{userId}/friends", appHandler.UserHandler.ListUserFriends).Methods("GET")

	log.Info("Application routers succesfully configured")

	return r
}

func main() {
	log.Info("Starting application")
	host := os.Getenv("pgHost")
	if host == "" {
		host = "localhost"
	}

	m, err := migrate.New(
		"file://data/migrations",
		fmt.Sprintf("postgresql://root:root@%s:5432/game?sslmode=disable", host),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Warn("Migrations did not run, error: ", err)
		} else {
			log.Fatal(err)
		}
	}

	pool := postgresql.CreatePool(host)

	userRepository := postgresql.NewUserRepository(pool)

	appHandler := NewApplicationHandler(handler.NewUserHandler(userRepository))
	router := Router(appHandler)

	http.ListenAndServe(":8080", router)
}