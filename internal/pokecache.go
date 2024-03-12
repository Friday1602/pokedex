package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createAt time.Time
	val      []byte
}

type Cache struct {
	store map[string]cacheEntry
	mu    sync.Mutex
}

// receive interval to remove cache with live longer than the interval
func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		store: make(map[string]cacheEntry),
	}

	go newCache.reapLoop(interval)
	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = cacheEntry{
		createAt: time.Now(),
		val:      val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cache, ok := c.store[key]
	return cache.val, ok

}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.store {
			if time.Since(value.createAt) > time.Duration(interval) {
				delete(c.store, key)
			}
		}
		c.mu.Unlock()
	}

}
