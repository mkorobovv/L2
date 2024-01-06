package repo

import (
	"sync"

	"github.com/mkorobovv/L2/develop/server/models"
)

type Cache struct {
	mu   sync.Mutex // guards
	data map[string][]*models.Event
}

func NewCache() *Cache {
	return &Cache{
		mu:   sync.Mutex{},
		data: make(map[string][]*models.Event),
	}
}

func (c *Cache) GetDay(date string) []*models.Event {

	c.mu.Lock()
	defer c.mu.Unlock()

	events, ok := c.data[date]
	if !ok {
		return []*models.Event{}
	}
	return events
}

func (c *Cache) Create(event *models.Event) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[event.Date] = append(c.data[event.Date], event)

}

func (c *Cache) Delete(date, time string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := 0; i < len(c.data[date]); i++ {
		if c.data[date][i].Time == time {
			c.data[date] = append(c.data[date][:i], c.data[date][i+1:]...)
		}
	}

}
