package cache

import (
	"sync"
	"wb-l0/domain"
)

type Cache struct {
	mu     sync.Mutex
	orders map[string]domain.Order
}

func NewCache() *Cache {
	return &Cache{
		orders: make(map[string]domain.Order),
	}
}

func (c *Cache) SetOrder(orderUID string, order domain.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.orders[orderUID] = order
}

func (c *Cache) GetOrder(orderUID string) (domain.Order, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	order, ok := c.orders[orderUID]
	return order, ok
}
