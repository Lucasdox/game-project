package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"

	"game-project/internal/application/command"
	"game-project/internal/application/query"
)

func TestUserHandler_List(t *testing.T) {
	cases := []struct {
		fakeServiceImpl *fakeServiceImpl
		expectedResult []*query.User
		expectedStatus int

	} {
		{
			fakeServiceImpl: &fakeServiceImpl{
				listResult: []*query.User{
					{
						Id:   uuid.UUID{},
						Name: "Jake",
					},
				},
			},
			expectedResult: []*query.User{
				{
					Id:   uuid.UUID{},
					Name: "Jake",
				},
			},
			expectedStatus: http.StatusOK,
		},
	}
	for _, tc := range cases {
		var response []*query.User
		handler := &UserHandler{Service: tc.fakeServiceImpl}
		r, _ := http.NewRequest("GET", "/user", nil)
		w := httptest.NewRecorder()
		router(handler).ServeHTTP(w, r)

		json.NewDecoder(w.Body).Decode(&response)
		if w.Code != tc.expectedStatus {
			t.Fatalf("wrong status retrieved, should be %d and received %d instead", w.Code, tc.expectedStatus)
		}
		if !reflect.DeepEqual(tc.expectedResult, response) {
			t.Fatalf("body response and expected result does not match. expected %+v and received %+v", tc.expectedResult, response)
		}
	}
}

func TestUserHandler_Create(t *testing.T) {
	cases := []struct {
		fakeServiceImpl *fakeServiceImpl
		expectedResult *query.User
		body string
		expectedErr error
		expectedStatus int
	} {
		{
			fakeServiceImpl: &fakeServiceImpl{
				createUserResult: &query.User{
					Id:   uuid.UUID{},
					Name: "Paul",
				},
			},
			body: "{\n\t\"name\": \"Jefferson\"\n}",
			expectedErr: nil,
			expectedResult: &query.User{
				Id:   uuid.UUID{},
				Name: "Paul",
			},
			expectedStatus: http.StatusCreated,
		},
	}
	for _, tc := range cases {
		var response *query.User
		handler := &UserHandler{Service: tc.fakeServiceImpl}
		r, _ := http.NewRequest("POST", "/user", strings.NewReader(tc.body))
		w := httptest.NewRecorder()
		router(handler).ServeHTTP(w, r)

		json.NewDecoder(w.Body).Decode(&response)
		if w.Code != tc.expectedStatus {
			t.Fatalf("wrong status retrieved, should be %d and received %d instead", w.Code, tc.expectedStatus)
		}
		if !reflect.DeepEqual(tc.expectedResult, response) {
			t.Fatalf("body response and expected result does not match. expected %+v and received %+v", tc.expectedResult, response)
		}
	}
}

func TestUserHandler_LoadUserState(t *testing.T) {
	cases := []struct {
		fakeServiceImpl *fakeServiceImpl
		expectedResult *query.UserGameStateQuery
		expectedErr error
		expectedStatus int
	} {
		{
			fakeServiceImpl: &fakeServiceImpl{
				loadUserStateResult: &query.UserGameStateQuery{
					GamesPlayed: 2,
					Score:       300,
				},
			},
			expectedErr: nil,
			expectedResult: &query.UserGameStateQuery{
				GamesPlayed: 2,
				Score:       300,
			},
			expectedStatus: http.StatusOK,
		},
	}
	for _, tc := range cases {
		var response *query.UserGameStateQuery
		handler := &UserHandler{Service: tc.fakeServiceImpl}
		id, _ := uuid.NewV4()
		r, _ := http.NewRequest("GET", fmt.Sprintf("/user/%s/state", id), nil)
		w := httptest.NewRecorder()
		router(handler).ServeHTTP(w, r)

		json.NewDecoder(w.Body).Decode(&response)
		if w.Code != tc.expectedStatus {
			t.Fatalf("wrong status retrieved, should be %d and received %d instead", w.Code, tc.expectedStatus)
		}
		if !reflect.DeepEqual(tc.expectedResult, response) {
			t.Fatalf("body response and expected result does not match. expected %+v and received %+v", tc.expectedResult, response)
		}
	}
}

func TestUserHandler_ListUserFriends(t *testing.T) {
	cases := []struct {
		fakeServiceImpl *fakeServiceImpl
		expectedResult *query.UserFriends
		expectedErr error
		expectedStatus int
	} {
		{
			fakeServiceImpl: &fakeServiceImpl{
				listUserFriendsResult: &query.UserFriends{
					Friends: []*query.Friend{
						{
							Id:        uuid.UUID{},
							Name:      "Jeferson",
							Highscore: 20,
						},
						{
							Id:        uuid.UUID{},
							Name:      "Claudio",
							Highscore: 35,
						},
					},
				},
			},
			expectedErr: nil,
			expectedResult: &query.UserFriends{
				Friends: []*query.Friend{
					{
						Id:        uuid.UUID{},
						Name:      "Jeferson",
						Highscore: 20,
					},
					{
						Id:        uuid.UUID{},
						Name:      "Claudio",
						Highscore: 35,
					},
				},
			},
			expectedStatus: http.StatusOK,
		},
	}
	for _, tc := range cases {
		var response *query.UserFriends
		handler := &UserHandler{Service: tc.fakeServiceImpl}
		id, _ := uuid.NewV4()
		r, _ := http.NewRequest("GET", fmt.Sprintf("/user/%s/friends", id), nil)
		w := httptest.NewRecorder()
		router(handler).ServeHTTP(w, r)

		json.NewDecoder(w.Body).Decode(&response)
		if w.Code != tc.expectedStatus {
			t.Fatalf("wrong status retrieved, should be %d and received %d instead", w.Code, tc.expectedStatus)
		}
		if !reflect.DeepEqual(tc.expectedResult, response) {
			t.Fatalf("body response and expected result does not match. expected %+v and received %+v", tc.expectedResult, response)
		}
	}
}

type fakeServiceImpl struct {
	listResult []*query.User
	createUserResult *query.User
	loadUserStateResult *query.UserGameStateQuery
	nUserFriendsUpdated int64
	listUserFriendsResult *query.UserFriends
	err error
}

func (f fakeServiceImpl) ListUser() []*query.User {
	return f.listResult
}

func (f fakeServiceImpl) CreateUser(user command.CreateUser) (*query.User, error) {
	return f.createUserResult, f.err
}

func (f fakeServiceImpl) UpdateUserState(userId uuid.UUID, command command.UpdateUserState) error {
	return f.err
}

func (f fakeServiceImpl) LoadUserState(userId uuid.UUID) (*query.UserGameStateQuery, error) {
	return f.loadUserStateResult, f.err
}

func (f fakeServiceImpl) UpdateUserFriends(userId uuid.UUID, command command.UpdateUserFriends) (int64, error) {
	return f.nUserFriendsUpdated, f.err
}

func (f fakeServiceImpl) ListUserFriends(userId uuid.UUID) (*query.UserFriends, error) {
	return f.listUserFriendsResult, f.err
}

func router(handler *UserHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user", handler.List).Methods("GET")
	r.HandleFunc("/user", handler.Create).Methods("POST")
	r.HandleFunc("/user/{userId}/state", handler.UpdateUserState).Methods("PUT")
	r.HandleFunc("/user/{userId}/state", handler.LoadUserState).Methods("GET")
	r.HandleFunc("/user/{userId}/friends", handler.UpdateUserFriends).Methods("PUT")
	r.HandleFunc("/user/{userId}/friends", handler.ListUserFriends).Methods("GET")

	return r
}
