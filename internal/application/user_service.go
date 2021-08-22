package application

import (
	"errors"

	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"

	"game-project/internal/application/command"
	"game-project/internal/application/query"
	"game-project/internal/domain"
)

type UserService interface {
	CreateUser(user command.CreateUser) (*domain.User, error)
	UpdateUserState(userId uuid.UUID, command command.UpdateUserState) error
	LoadUserState(userId uuid.UUID) (*query.UserGameStateQuery, error)
}

type UserServiceImpl struct {
	repository domain.UserRepository
}

func (s *UserServiceImpl) LoadUserState(userId uuid.UUID) (*query.UserGameStateQuery, error) {
	usr := s.repository.FindUser(userId)
	var err error
	if usr == nil {
		return nil, errors.New("no user found")
	}
	state := query.UserGameStateQuery{
		GamesPlayed: usr.GamesPlayed.Int32,
		Score:       usr.Score.Int64,
	}

	return &state, err
}

func (s *UserServiceImpl) CreateUser(user command.CreateUser) (*domain.User, error) {
	usr, err := s.repository.Create(user.Name)
	if err != nil {
		log.Warn("error inserting user", err)
	}

	return usr, err
}

func (s *UserServiceImpl) UpdateUserState(userId uuid.UUID, command command.UpdateUserState) error {
	usrInDb := s.repository.FindUser(userId)
	if usrInDb == nil {
		log.Warnf("no user with id %s found", usrInDb.Id)
		return errors.New("no user found")
	}
	return s.repository.UpdateUserState(userId, command.GamesPlayed, command.Score)
}

func NewUserService(r domain.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repository: r}
}
