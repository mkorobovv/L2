package repo

import (
	"sync"

	"github.com/mkorobovv/L2/develop/server/models"
)

type Cache struct {
	mu   sync.Mutex // guards
	data map[string]*models.EventCalendar
}

func NewCache() *Cache {
	return &Cache{
		mu:   sync.Mutex{},
		data: make(map[string]*models.EventCalendar),
	}
}

func (c *Cache) Get(date string) *models.EventCalendar {
	c.mu.Lock()
	defer c.mu.Unlock()

	events, ok := c.data[date]
	if !ok {
		return nil
	}
	return events
}
