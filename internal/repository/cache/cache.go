package cache

import (
	"github.com/Nigelmes/L0/internal/models"
	"github.com/jinzhu/gorm"
	"sync"
)

type Cache struct {
	mx        sync.RWMutex
	cacheItem map[string]models.Order
}

func NewCache(db *gorm.DB) *Cache {
	return &Cache{cacheItem: make(map[string]models.Order)}
}

func (c *Cache) Set(order models.Order) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.cacheItem[order.OrderUid] = order
}

func (c *Cache) GetByUUID(uuid string) (models.Order, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.cacheItem[uuid]
	return val, ok
}
