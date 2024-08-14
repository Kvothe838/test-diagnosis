package main

import (
	"context"
	"flag"

	"github.com/Kvothe838/test-diagnosis/config"
	"github.com/Kvothe838/test-diagnosis/internal/app/controller"
	"github.com/Kvothe838/test-diagnosis/internal/app/server"
	"github.com/Kvothe838/test-diagnosis/internal/database/postgres"
	"github.com/Kvothe838/test-diagnosis/internal/pkg/clock"
	"github.com/Kvothe838/test-diagnosis/internal/pkg/graceful"
	"github.com/Kvothe838/test-diagnosis/internal/pkg/logger"
	"github.com/Kvothe838/test-diagnosis/internal/pkg/uuid"
	"github.com/Kvothe838/test-diagnosis/internal/services"
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
