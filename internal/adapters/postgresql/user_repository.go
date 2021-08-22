package postgresql

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"

	"game-project/internal/domain"
)

const (
	INSERT_USER = `INSERT into game.public.user (id, name) VALUES ($1, $2);`
)

type UserRepositoryImpl struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepositoryImpl{
	return &UserRepositoryImpl{pool: pool}
}

func (r *UserRepositoryImpl) Create(uName string) (*domain.User, error) {
	uuid, _ := uuid.NewV4()
	user := &domain.User{
		Id: uuid ,
		Name: uName,
	}

	_, err := r.pool.Exec(context.Background(), INSERT_USER, user.Id, user.Name)
	if err != nil {
		log.Warn("error inserting user", err)
		return nil, err
	}

	return user, err
}
