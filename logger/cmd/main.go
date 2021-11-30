package main

import (
	"context"
	"runtime"
	"time"

	"github.com/joho/godotenv"
	"github.com/p12s/library-rest-api/logger/pkg/config"
	"github.com/p12s/library-rest-api/logger/pkg/handler"
	"github.com/p12s/library-rest-api/logger/pkg/repository"
	"github.com/p12s/library-rest-api/logger/pkg/server"
	"github.com/p12s/library-rest-api/logger/pkg/service"
	"github.com/sirupsen/logrus"
)

// @title gRPC-logger
// @version 0.0.1
// @description Simple gRPC-logger for saving logs into MongoDB
// @host localhost:81
// @BasePath /
func main() {
	runtime.GOMAXPROCS(1)
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error reading env variables from file: %s\n", err.Error())
	}

	cfg, err := config.New()
	if err != nil {
		logrus.Fatalf("error loading env variables: %s\n", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := repository.NewMongoDB(ctx, repository.Config{
		Username: cfg.Db.User,
		Password: cfg.Db.Password,
		Url:      cfg.Db.Uri,
	})
	if err != nil {
		logrus.Fatalf("failed to initialize mongodb: %s\n", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewLogService(repos)
	handler := handler.NewHandler(service)
	srv := server.New(handler)

	logrus.Printf("Server started on port %d. %s", cfg.Server.Port, time.Now().Format(time.RFC3339))

	srv.ListenAndServe(cfg.Server.Port)
}
