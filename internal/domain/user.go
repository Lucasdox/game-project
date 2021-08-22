package domain

import(
	"database/sql"

	"github.com/gofrs/uuid"
)

type User struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
	GamesPlayed sql.NullInt32 `json:"gamesPlayed,omitempty"`
	Score sql.NullInt64 `json:"score,omitempty"`
}

type UserRepository interface {
	Create(uName string) (*User, error)
	UpdateUserState(userId uuid.UUID, gamesPlayed uint8, score uint) error
	FindUser(userId uuid.UUID) *User
	UpdateFriends(userId uuid.UUID, friendLst []uuid.UUID) (touchedRows int64, err error)
}
