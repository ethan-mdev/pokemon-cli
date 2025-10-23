package pokecache

import (
	"time"
)

// Creates a new Cache instance and starts the cleanup goroutine
func NewCache() *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: 5 * time.Minute,
	}
	go func(c *Cache) {
		c.readLoop()
	}(cache)
	return cache
}

// Retrieves a value from the cache by key
func (c *Cache) Get(key string) ([]byte, bool) {
	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

// Adds a new key-value pair to the cache
func (c *Cache) Add(key string, val []byte) {
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Periodically cleans up expired entries from the cache
func (c *Cache) readLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		now := time.Now()
		for key, entry := range c.entries {
			if now.Sub(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
	}
}
