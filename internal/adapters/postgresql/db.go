package postgresql

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

func CreatePool() *pgxpool.Pool{
	connString := "postgresql://root:root@localhost:5432/game?sslmode=disable"
	config, err := pgxpool.ParseConfig(connString)

	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("error starting the database: ", err)
	}

	log.Info("Database initialized")
	return pool
}
