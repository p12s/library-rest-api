package main

import (
	"runtime"

	"github.com/joho/godotenv"
	"github.com/p12s/library-rest-api/logger-2/config"
	"github.com/p12s/library-rest-api/logger-2/consumer"
	"github.com/sirupsen/logrus"
)

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

	queueConsumer, err := consumer.New(*cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize queue consumer: %s\n", err.Error())
	}
	defer queueConsumer.Close()

	queueConsumer.Consume()
}
