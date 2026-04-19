package main

import (
	"fmt"
	"sync"
	"time"
)

type Item struct {
	value any
	ttl   time.Time
}

type Cache struct {
	storage         map[string]Item
	mu              sync.RWMutex
	cleanupInterval time.Duration
	cleanupTimeout  time.Duration
	stopCh          chan struct{}
}

func newCache(cleanupInterval time.Duration, cleanupTimeout time.Duration) *Cache {
	newCache := &Cache{
		storage:         make(map[string]Item),
		mu:              sync.RWMutex{},
		cleanupInterval: cleanupInterval,
		cleanupTimeout:  cleanupTimeout,
		stopCh:          make(chan struct{}),
	}

	go newCache.runCleanup()

	return newCache
}

func (c *Cache) runCleanup() {
	ticker := time.NewTicker(c.cleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.cleanup()
		case <-c.stopCh:
			return
		}
	}
}

func (c *Cache) cleanup() {
	startTime := time.Now()
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.storage {
		if startTime.After(v.ttl) {
			delete(c.storage, k)
		}
	}
}

func (c *Cache) stopCleanup() {
	close(c.stopCh)
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	c.storage[key] = Item{value: value, ttl: time.Now().Add(ttl)}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.storage[key]
	if time.Since(val.ttl) > 0 {
		delete(c.storage, key)
		return nil, false
	}
	return val, ok
}

func main() {
	newC := newCache(3*time.Second, 1*time.Second)
	newC.Set("a", 1, 2*time.Second)
	newC.Set("b", 2, 3*time.Second)
	newC.Set("c", 3, 1*time.Second)
	newC.Set("d", 4, 4*time.Second)
	newC.Set("e", 5, 3*time.Second)
	newC.Set("f", 6, 3*time.Second)
	newC.Set("g", 7, 5*time.Second)

	time.Sleep(3 * time.Second)

	fmt.Println(newC.storage)
}
