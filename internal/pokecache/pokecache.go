package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheMap map[string]cacheEntry
	mutex    *sync.Mutex
}

// NewCache -
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheMap: make(map[string]cacheEntry),
		mutex:    &sync.Mutex{},
	}
	go cache.reapLoop(interval)

	return cache
}

// Add -
func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

// Get -
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	value, exists := c.cacheMap[key]
	if !exists {
		return nil, false
	}
	return value.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.cacheMap {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cacheMap, k)
		}
	}
}
