package application

import (
	"fmt"
	"game-project/internal/application/command"
	"game-project/internal/application/query"
	"game-project/internal/domain"
	"github.com/gofrs/uuid"
	"reflect"
	"testing"
)

func TestUserServiceImpl_CreateUser(t *testing.T) {
	cases := []struct{
		fakeRepository *fakeUserRepository
		command command.CreateUser
		expectedResult *query.User
		expectedErr error
	} {
		{
			fakeRepository: &fakeUserRepository{
				createMock:      &domain.User{
					Id:          uuid.UUID{},
					Name:        "Jeferson",
				},
			},
			command: command.CreateUser{
				Name: "Jeferson",
			},
			expectedResult: &query.User{
				Id:   uuid.UUID{},
				Name: "Jeferson",
			},
			expectedErr:    nil,
		},
	}

	for _, tc := range cases {
		service := UserServiceImpl{repository: tc.fakeRepository}
		result, err := service.CreateUser(tc.command)
		if (err != tc.expectedErr) {
			t.Errorf(fmt.Sprintf("expected err: %s, actual err: %s", err, tc.expectedErr))
		}
		if (!reflect.DeepEqual(result, tc.expectedResult)) {
			t.Errorf(fmt.Sprintf("expected result: %v, actual: %v", tc.expectedResult, result))
		}
	}
}

func TestUserServiceImpl_ListUser(t *testing.T) {
	cases := []struct {
		fakeRepository *fakeUserRepository
		expectedResult []*query.User
		expectedErr error
	}{
		{
			fakeRepository: &fakeUserRepository{
				listMock:        []*domain.User{
					{
						Id:          uuid.UUID{},
						Name:        "Jake",
					},
				},
				errMock:         nil,
			},
			expectedResult: []*query.User{
				{
					Id:   uuid.UUID{},
					Name: "Jake",
				},
			},
			expectedErr:    nil,
		},
	}

	for _, tc := range cases {
		service := UserServiceImpl{repository: tc.fakeRepository}
		result := service.ListUser()
		if !reflect.DeepEqual(result, tc.expectedResult) {
			t.Errorf("expected result: %v , actual: %v", tc.expectedResult, result)
		}
	}
}

type fakeUserRepository struct {
	listMock []*domain.User
	createMock *domain.User
	findUserMock *domain.User
	listFriendsMock []*domain.User
	errMock error
}

func (f fakeUserRepository) List() []*domain.User {
	return f.listMock
}

func (f fakeUserRepository) Create(uName string) (*domain.User, error) {
	return f.createMock, f.errMock
}

func (f fakeUserRepository) UpdateUserState(userId uuid.UUID, gamesPlayed uint8, score uint) error {
	panic("implement me")
}

func (f fakeUserRepository) FindUser(userId uuid.UUID) *domain.User {
	return f.findUserMock
}

func (f fakeUserRepository) UpdateFriends(userId uuid.UUID, friendLst []uuid.UUID) (touchedRows int64, err error) {
	panic("implement me")
}

func (f fakeUserRepository) ListFriends(userId uuid.UUID) []*domain.User {
	return f.listFriendsMock
}
