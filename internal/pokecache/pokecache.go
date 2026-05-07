package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]CacheEntry
	interval time.Duration
	mux      *sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.entries[key] = CacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	entry, exists := c.entries[key]
	c.mux.RUnlock()
	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mux.Lock()
		for key, entry := range c.entries {
			if time.Now().After(entry.createdAt.Add(c.interval)) {
				delete(c.entries, key)
			}
		}
		c.mux.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]CacheEntry),
		interval: interval,
		mux:      &sync.RWMutex{},
	}

	go cache.reapLoop()

	return cache
}
