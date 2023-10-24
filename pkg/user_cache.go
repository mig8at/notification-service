package pkg

import (
	"sync"

	"github.com/mig8at/notifications/internal/model"
)

type UserCache struct {
	mu    sync.RWMutex
	Users map[model.NotificationType]map[int]model.User
}

func NewUserCache() *UserCache {
	return &UserCache{
		Users: make(map[model.NotificationType]map[int]model.User),
	}
}

func (c *UserCache) Set(t model.NotificationType, id int, user model.User) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.Users[t]; !ok {
		c.Users[t] = make(map[int]model.User)
	}
	c.Users[t][id] = user
}

func (c *UserCache) Get(t model.NotificationType, id int) (model.User, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	user, exists := c.Users[t][id]
	return user, exists
}

func (c *UserCache) Delete(t model.NotificationType, id int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.Users[t], id)
}
