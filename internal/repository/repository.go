package repository

import (
	"github.com/Nigelmes/L0/internal/models"
	"github.com/Nigelmes/L0/internal/repository/cache"
	"github.com/Nigelmes/L0/internal/repository/database"
	"github.com/jinzhu/gorm"
)

type OrderRepo interface {
	GetAll() ([]models.Order, error)
	Create(order models.Order) error
}

type CacheRepo interface {
	Set(order models.Order)
	GetByUUID(uuid string) (models.Order, bool)
}

type Repository struct {
	OrderRepo
	CacheRepo
}

func NewRepository(db *gorm.DB) *Repository {
	rdb := database.NewOrderPostgres(db)
	rch := cache.NewCache(db)
	items, err := rdb.GetAll()
	if err != nil {
		return &Repository{
			OrderRepo: rdb,
			CacheRepo: rch,
		}
	}
	for _, item := range items {
		rch.Set(item)
	}
	return &Repository{
		OrderRepo: rdb,
		CacheRepo: rch,
	}
}
