package application

import (
	"game-project/internal/application/command"
	"game-project/internal/domain"
)

type UserService interface {
	CreateUser(user command.CreateUser) (*domain.User, error)
}

type UserServiceImpl struct {
	repository domain.UserRepository
}

func (s *UserServiceImpl) CreateUser(user command.CreateUser) (*domain.User, error) {
	return s.repository.Create(user.Name)
}

func NewUserService(r domain.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repository: r}
}
