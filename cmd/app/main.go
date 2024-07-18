package main

import (
	"TopDoctorsBackendChallenge/config"
	"TopDoctorsBackendChallenge/internal/app/controller"
	"TopDoctorsBackendChallenge/internal/app/server"
	"TopDoctorsBackendChallenge/internal/database/postgres"
	"TopDoctorsBackendChallenge/internal/pkg/clock"
	"TopDoctorsBackendChallenge/internal/pkg/graceful"
	"TopDoctorsBackendChallenge/internal/pkg/logger"
	"TopDoctorsBackendChallenge/internal/pkg/uuid"
	"TopDoctorsBackendChallenge/internal/services"
	"context"
	"flag"
)

func main() {
	filePath := flag.String("config", "", "path of configuration file")
	flag.Parse()

	ctx := context.Background()

	var sources []config.Source
	conf := config.New(ctx, *filePath, sources...)

	db := postgres.New(ctx, conf.Postgres.User, conf.Postgres.DbName, conf.Postgres.Password, conf.Postgres.Host)
	UUID := uuid.NewReal()
	clock := clock.NewReal()

	interactor := services.NewInteractor(db, UUID, clock)

	setupRestAPI(ctx, interactor, conf.Port)

	if err := graceful.Wait(); err != nil {
		logger.CtxWarn(ctx, err)
	}
}

func setupRestAPI(ctx context.Context, interactor controller.Interactor, port string) {
	ctrl := controller.New(interactor)
	srv := server.New(port)
	srv.RegisterHandler(ctrl)
	srv.StartAsync()

	logger.CtxInfof(ctx, "listening on :%s", port)
}
