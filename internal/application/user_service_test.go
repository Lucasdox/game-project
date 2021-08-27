package application

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/gofrs/uuid"

	"game-project/internal/application/command"
	"game-project/internal/application/query"
	"game-project/internal/domain"
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

func TestUserServiceImpl_ListUserFriends(t *testing.T) {
	cases := []struct {
		fakeRepository *fakeUserRepository
		expectedResult *query.UserFriends
		expectedErr error
	}{
		{
			fakeRepository: &fakeUserRepository{
				listFriendsMock: []*domain.User{
					{
						Id:   uuid.UUID{},
						Name: "Jessica",
						GamesPlayed: sql.NullInt32{
							Int32: 1,
							Valid: true,
						},
						Score: sql.NullInt64{
							Int64: 200,
							Valid: true,
						},
					},
					{
						Id:          uuid.UUID{},
						Name:        "Paul",
					},
				},
				findUserMock: &domain.User{
					Id:          uuid.UUID{},
					Name:        "Don",
				},
			},
			expectedResult: &query.UserFriends{
				Friends: []*query.Friend{
					{
						Id:        uuid.UUID{},
						Name:      "Jessica",
						Highscore: 200,
					},
					{
						Id:        uuid.UUID{},
						Name:      "Paul",
						Highscore: 0,
					},
				},
			},
			expectedErr:    nil,
		},
		{
			fakeRepository: &fakeUserRepository{
				listFriendsMock: []*domain.User{
					{
						Id:   uuid.UUID{},
						Name: "Jessica",
						GamesPlayed: sql.NullInt32{
							Int32: 1,
							Valid: true,
						},
						Score: sql.NullInt64{
							Int64: 200,
							Valid: true,
						},
					},
					{
						Id:          uuid.UUID{},
						Name:        "Paul",
					},
				},
				findUserMock: nil,
			},
			expectedResult: nil,
			expectedErr:    errors.New(fmt.Sprintf("no user with id %s found", uuid.UUID{})),
		},
	}

	for _, tc := range cases {
		service := UserServiceImpl{repository: tc.fakeRepository}
		response, err := service.ListUserFriends(uuid.UUID{})

		if tc.expectedErr != nil {
			if err.Error() != tc.expectedErr.Error() {
				t.Errorf("expected err %s, actual: %s", tc.expectedErr, err)
			}
		} else {
			if err != tc.expectedErr {
				t.Errorf("expected err %s, actual: %s", tc.expectedErr, err)
			}
		}

		if tc.expectedResult != nil {
			if !reflect.DeepEqual(response, tc.expectedResult){
				t.Errorf("expected respone to be %v, actual: %v", tc.expectedResult, response)
			}
		} else {
			if response != nil {
				t.Error("result should be nil")
			}
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
