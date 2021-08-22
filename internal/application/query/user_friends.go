package query

import "github.com/gofrs/uuid"

type Friend struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Highscore int32      `json:"highscore,omitempty"`
}

type UserFriends struct {
	Friends []*Friend `json:"friends,omitempty"`
}
