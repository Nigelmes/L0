package main

import (
	"context"
	"github.com/Nigelmes/L0/internal/config"
	"github.com/Nigelmes/L0/internal/handler"
	"github.com/Nigelmes/L0/internal/repository"
	"github.com/Nigelmes/L0/internal/server"
	"github.com/Nigelmes/L0/internal/streamingservice"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.GetConfig()
	logrus.Println("сonfiguration parsed successfully from .env file")

	db := repository.NewPostgresDB(cfg)
	logrus.Println("database connection successful")

	repo := repository.NewRepository(db)
	handlers := handler.NewHandler(repo)

	sc := streamingservice.NewNatsStream(cfg)
	sc.RunNatsSteaming(repo)

	server := new(server.Server)
	go func() {
		if err := server.Run(cfg, handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error running http server: %s", err.Error())
		}
	}()
	logrus.Print("server started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := sc.ShutDown(); err != nil {
		logrus.Errorf("error occured on nats streaming shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close : %s", err.Error())
	}
}
