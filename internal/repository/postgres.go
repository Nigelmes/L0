package repository

import (
	"fmt"
	"github.com/Nigelmes/L0/internal/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

func NewPostgresDB(cfg *config.Config) *gorm.DB {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Dbname,
		cfg.Database.Password,
	)

	db, err := gorm.Open("postgres", conn)
	if err != nil {
		logrus.Fatalf("error connection database, %s", err.Error())
		return nil
	}
	err = db.DB().Ping()
	if err != nil {
		logrus.Fatalf("error pinging the database:: %s", err.Error())
		return nil
	}
	return db
}
