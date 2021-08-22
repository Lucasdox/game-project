package domain

import(
	"github.com/gofrs/uuid"
)

type User struct {
	Id uuid.UUID `json:"id"`
	Name string `json:"name"`
}

type UserRepository interface {
	Create(uName string) (User, error)
}
