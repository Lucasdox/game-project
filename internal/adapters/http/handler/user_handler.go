package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

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

func (h *UserHandler) LoadUserState(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := uuid.FromString(vars["userId"])
	if err != nil {
		log.Warn("Request with invalid UUID")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	state, err := h.Service.LoadUserState(id)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(state)

	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func (h *UserHandler) UpdateUserFriends(writer http.ResponseWriter, request *http.Request) {
	var command command.UpdateUserFriends
	vars := mux.Vars(request)
	id, err := uuid.FromString(vars["userId"])
	if err != nil {
		log.Warn("Request with invalid UUID")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&command)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	newFriends, err := h.Service.UpdateUserFriends(id, command)
	if err != nil {
		log.Warn("error inserting friends")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Infof("%d new friends created", newFriends)

	writer.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) ListUserFriends(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, err := uuid.FromString(vars["userId"])
	if err != nil {
		log.Warn("Request with invalid UUID")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userFriends := h.Service.ListUserFriends(id)
	res, _ := json.Marshal(userFriends)

	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}
