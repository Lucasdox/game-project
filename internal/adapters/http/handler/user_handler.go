package handler

import (
	"encoding/json"
	"net/http"

	"game-project/internal/application"
	"game-project/internal/application/command"
)

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