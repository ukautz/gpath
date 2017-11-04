package gpath

import (
	"strings"
	"sync"
)

// traversals implements concurrency safe map (I want to support Go < 1.9, so no sync.Map ..)
type cache struct {
	data map[string]interface{}
	mux  *sync.RWMutex
}

func newCache(data map[string]interface{}) *cache {
	return &cache{
		data: data,
		mux:  new(sync.RWMutex),
	}
}

func (c *cache) get(key string) (interface{}, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	if v, ok := c.data[key]; ok {
		return v, true
	}
	return nil, false
}

func (c *cache) set(key string, value interface{}) interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data[key] = value
	return value
}

func (c *cache) clear(prefix string) int {
	count := 0
	keys := []string{}
	c.mux.Lock()
	defer c.mux.Unlock()
	for key, _ := range c.data {
		keys = append(keys, key)
	}
	for _, key := range keys {
		if strings.HasPrefix(key, prefix) {
			count++
			delete(c.data, key)
		}
	}
	return count
}
