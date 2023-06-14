package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"log"
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
}

func GetConfig() *Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
	cfg := new(Config)
	if err = env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}