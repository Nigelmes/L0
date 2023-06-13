package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Database struct {
		Host     string `env:"DBHOST"`
		Port     string `env:"DBPORT"`
		User     string `env:"DBUSERNAME"`
		Dbname   string `env:"DBNAME"`
		Password string `env:"DBPASSWORD"`
	}
	Server struct {
		Host string `env:"SERVERHOST"`
		Port string `env:"SERVERPORT"`
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
