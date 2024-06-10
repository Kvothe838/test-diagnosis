package main

import (
	"TopDoctorsBackendChallenge/internal/app/controller"
	"TopDoctorsBackendChallenge/internal/app/server"
	"TopDoctorsBackendChallenge/internal/pkg/graceful"
	"TopDoctorsBackendChallenge/internal/pkg/logger"
	"context"
)

func main() {
	ctx := context.Background()

	setupRestAPI(ctx, "8080")

	if err := graceful.Wait(); err != nil {
		logger.CtxWarn(ctx, err)
	}
}

func setupRestAPI(ctx context.Context, port string) {
	ctrl := controller.New()
	srv := server.New(port)
	srv.RegisterHandler(ctrl)
	srv.StartAsync()

	logger.CtxInfof(ctx, "listening on :%s", port)
}
