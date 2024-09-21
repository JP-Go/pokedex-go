package cache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	Val       []byte
	createdAt time.Time
}
type Cache struct {
	store map[string]CacheEntry
	mu    sync.RWMutex
}

func NewCache(evictTime time.Duration) *Cache {
	cache := Cache{
		store: make(map[string]CacheEntry),
		mu:    sync.RWMutex{},
	}
	cache.readLoop(evictTime)

	return &cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = CacheEntry{
		Val:       value,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) (entry CacheEntry, exists bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.store[key]
	if !ok {
		return CacheEntry{}, false
	}
	return val, true
}

func (c *Cache) readLoop(evictTime time.Duration) {
	ticker := time.NewTicker(evictTime)
	go func() {
		for range ticker.C {
			c.mu.Lock()
			for key, val := range c.store {
				if time.Since(val.createdAt) > evictTime {
					delete(c.store, key)
				}
			}
			c.mu.Unlock()

		}
	}()
}
