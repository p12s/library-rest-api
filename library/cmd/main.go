package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/p12s/library-rest-api/library/pkg/config"
	"github.com/p12s/library-rest-api/library/pkg/grpcClient"
	"github.com/p12s/library-rest-api/library/pkg/handler"
	"github.com/p12s/library-rest-api/library/pkg/queueClient"
	"github.com/p12s/library-rest-api/library/pkg/repository"
	"github.com/p12s/library-rest-api/library/pkg/service"
	"github.com/sirupsen/logrus"
)

// @title Library app REST-API
// @version 0.0.1
// @description Simple library application for adding/getting books
// @host localhost:80
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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

	db, err := repository.NewPostgresDB(*cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s\n", err.Error())
	}
	repos := repository.NewRepository(db)

	services := service.NewService(repos)
	grpcLogger, err := grpcClient.New(*cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize grpc client: %s\n", err.Error())
	}
	queueLogger, err := queueClient.New(*cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize queue client: %s\n", err.Error())
	}

	handlers := handler.NewHandler(services, grpcLogger, queueLogger)

	srv := new(Server)
	go func() {
		if err := srv.Run(cfg.Server.Port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error while running http server: %s\n", err.Error())
		}
	}()
	logrus.Print("app started on port ", cfg.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("app shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred on db connection close: %s", err.Error())
	}
	if err := queueLogger.Close(); err != nil {
		logrus.Errorf("error occurred on queue connection close: %s", err.Error())
	}
}

// Server - http server
type Server struct {
	httpServer *http.Server
}

// Run - start
func (s *Server) Run(port int, handler http.Handler) error {
	address := ":" + strconv.Itoa(port)
	s.httpServer = &http.Server{
		Addr:           address,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown - grace-full
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
