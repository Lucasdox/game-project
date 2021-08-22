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
	UPDATE_USER = `UPDATE game.public.user SET games_played = $1, score = GREATEST(score, $2) WHERE id = $3;`
	SELECT_USER = `SELECT id, name, games_played, score FROM game.public.user WHERE id = $1;`
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

		return nil, err
	}

	return user, err
}

func (r *UserRepositoryImpl) UpdateUserState(userId uuid.UUID, gamesPlayed uint8, score uint) error {
	_, err := r.pool.Exec(context.Background(), UPDATE_USER, gamesPlayed, score, userId)

	return err
}

func (r *UserRepositoryImpl) FindUser(userId uuid.UUID) *domain.User {
	var user domain.User
	row := r.pool.QueryRow(context.Background(), SELECT_USER, userId)

	err := row.Scan(&user.Id, &user.Name, &user.GamesPlayed, &user.Score)
	if err != nil {
		log.Warnf("User with id %s not found", userId)
		return nil
	}

	return &user
}
