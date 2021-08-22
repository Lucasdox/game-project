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
	CreateUser(user command.CreateUser) (*query.User, error)
	UpdateUserState(userId uuid.UUID, command command.UpdateUserState) error
	LoadUserState(userId uuid.UUID) (*query.UserGameStateQuery, error)
	UpdateUserFriends(userId uuid.UUID, command command.UpdateUserFriends) (int64, error)
	ListUserFriends(userId uuid.UUID) *query.UserFriends
}

type UserServiceImpl struct {
	repository domain.UserRepository
}

func (s *UserServiceImpl) ListUserFriends(userId uuid.UUID) *query.UserFriends {
	return s.repository.ListFriends(userId)
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

func (s *UserServiceImpl) CreateUser(user command.CreateUser) (*query.User, error) {
	usr, err := s.repository.Create(user.Name)
	if err != nil {
		log.Warn("error inserting user", err)
	}

	query := query.User{
		Id:   usr.Id,
		Name: usr.Name,
	}

	return &query, err
}

func (s *UserServiceImpl) UpdateUserState(userId uuid.UUID, command command.UpdateUserState) error {
	usrInDb := s.repository.FindUser(userId)
	if usrInDb == nil {
		log.Warnf("no user with id %s found", usrInDb.Id)
		return errors.New("no user found")
	}
	return s.repository.UpdateUserState(userId, command.GamesPlayed, command.Score)
}

func (s *UserServiceImpl) UpdateUserFriends(userId uuid.UUID, command command.UpdateUserFriends) (int64, error) {
	n, err := s.repository.UpdateFriends(userId, command.Friends)
	if err != nil {
		log.Warnf("could not insert friends for userid: %s", n)
		return 0, err
	}
	return n, err
}

func NewUserService(r domain.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repository: r}
}
