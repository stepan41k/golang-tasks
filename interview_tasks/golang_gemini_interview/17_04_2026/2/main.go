package main

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	value any
	ttl   time.Duration
}

type Cache struct {
	storage         map[string]any
	mu              *sync.RWMutex
	cleanupInterval time.Duration
	cleanupTimeout  time.Duration
}

func newCache(cleanupInterval time.Duration, cleanupTimeout time.Duration) *Cache {
	newCache := &Cache{
		storage:         make(map[string]any),
		mu: &sync.RWMutex{},
		cleanupInterval: cleanupInterval,
		cleanupTimeout: cleanupTimeout,
	}

	newCache.cleanupCache()

	return newCache
}

func (c *Cache) cleanupCache() {
	ticker := time.NewTicker(c.cleanupInterval)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				c.mu.Lock()
				startTime := time.Now()
				
				for k, _ := range c.storage {
					delete(c.storage, k)
					if time.Since(startTime) > c.cleanupTimeout {
						return
					}
				}
			}
		}
	}()
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = Item{value: value, ttl: ttl}
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.storage[key]
	return val, ok
}

func main() {
	newC := newCache(5 * time.Second, 1 * time.Second)
	newC.Set("a", 1, 2 * time.Second)
	newC.Set("b", 2, 3 * time.Second)
	newC.Set("c", 3, 1 * time.Second)
	newC.Set("d", 4, 4 * time.Second)
	newC.Set("e", 5, 3 * time.Second)
	newC.Set("f", 6, 3 * time.Second)
	newC.Set("g", 7, 5 * time.Second)
	
	time.Sleep(2 * time.Second)
	
	fmt.Println(newC.storage)
}
