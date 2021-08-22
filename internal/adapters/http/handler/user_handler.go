package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"

	"game-project/internal/application"
	"game-project/internal/application/command"
	"game-project/internal/domain"
)

func NewUserHandler(repository domain.UserRepository) *UserHandler {
	return &UserHandler{Service: application.NewUserService(repository)}
}

type UserHandler struct {
	Service application.UserService
}

func (h *UserHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var command command.CreateUser

	err := json.NewDecoder(request.Body).Decode(&command)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.Service.CreateUser(command)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(u)

	writer.WriteHeader(http.StatusCreated)
	writer.Write(res)
}

func (h *UserHandler) UpdateUserState(writer http.ResponseWriter, request *http.Request) {
	var command command.UpdateUserState
	vars := mux.Vars(request)
	id, err := uuid.FromString(vars["userId"])
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&command)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.UpdateUserState(id, command)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
