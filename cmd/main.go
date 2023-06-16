package main

import (
	"github.com/Nigelmes/L0/internal/config"
	"github.com/Nigelmes/L0/internal/repository"
	"github.com/Nigelmes/L0/internal/streamingservice"
	"github.com/sirupsen/logrus"
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

	sc := streamingservice.NewNatsStream(cfg)
	sc.RunNatsSteaming(repo)
	//streamingservice.RunNatsSteaming(repo, cfg)

	//a := repo.CacheRepo.GetAlls()
	//for _, b := range a{
	//	fmt.Println(b)
	//}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	sc.ShutDown()
}
