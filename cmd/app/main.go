package main

import (
	"TopDoctorsBackendChallenge/internal/app/controller"
	"TopDoctorsBackendChallenge/internal/app/server"
	"TopDoctorsBackendChallenge/internal/database/postgres"
	"TopDoctorsBackendChallenge/internal/pkg/graceful"
	"TopDoctorsBackendChallenge/internal/pkg/logger"
	"TopDoctorsBackendChallenge/internal/services"
	"context"
)

func main() {
	ctx := context.Background()

	db := postgres.New(ctx, "postgres", "topDoctors", "admin", "localhost")
	interactor := services.NewInteractor(db)

	setupRestAPI(ctx, interactor, "8080")

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
