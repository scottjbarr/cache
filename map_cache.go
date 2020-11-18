package cache

import (
	"sync"
)

// MapCache is a simple map backed Cache.
//
// Access to the map is protected by a mutex.
type MapCache struct {
	store map[string]interface{}
	mu    *sync.RWMutex
}

// NewMapCache returns a new MapCache.
func NewMapCache() MapCache {
	return MapCache{
		store: map[string]interface{}{},
		mu:    &sync.RWMutex{},
	}
}

// Get implements the CacheReader interface.
func (c MapCache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	d, ok := c.store[key]
	if !ok {
		// cache miss, nothing to return
		return nil, nil
	}

	return d, nil
}

// Set implements the CacheWriter interface.
func (c MapCache) Set(key string, data interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.store[key] = data

	return nil
}
