package command

import "github.com/gofrs/uuid"

type UpdateUserFriends struct {
	Friends []uuid.UUID `json:"friends"`
}
