package postgres

import (
	"TopDoctorsBackendChallenge/internal/pkg/logger"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type repository struct {
	conn *sqlx.DB
}

func New(ctx context.Context, user, dbName, password, host string) *repository {
	params := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s", user, dbName, password, host)
	db := sqlx.MustConnect("postgres", params)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		logger.CtxInfof(ctx, "Successfully connected to Postgres database with name %s", dbName)
	}

	return &repository{
		conn: db,
	}
}
