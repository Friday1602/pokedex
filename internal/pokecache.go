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

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]cacheEntry),
	}
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

func (c *Cache) reapLoop() {
	time.Sleep(5 * time.Second)
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.store, "a")
}
