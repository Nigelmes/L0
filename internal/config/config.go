package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"regexp"
)

type Config struct {
	Database struct {
		Host     string `env:"DBHOST" envDefault:"localhost"`
		Port     string `env:"DBPORT" envDefault:"5432"`
		User     string `env:"DBUSERNAME,required"`
		Dbname   string `env:"DBNAME,required"`
		Password string `env:"DBPASSWORD,required"`
	}
	Server struct {
		Host string `env:"SERVERHOST" envDefault:"localhost"`
		Port string `env:"SERVERPORT" envDefault:"8080"`
	}
	NatsStreamingCfg struct {
		StanClusterId string `env:"STANCLUSTERID"`
		ClientId      string `env:"CLIENTID"`
	}
}

var projectDirName = "wbl0"

func GetConfig() *Config {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		logrus.Fatalf("failed to load .env file: %e", err)
	}
	cfg := new(Config)
	if err = env.Parse(cfg); err != nil {
		logrus.Fatal(err)
	}
	return cfg
}
