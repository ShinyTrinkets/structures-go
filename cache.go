package structs

import (
	"encoding/json"
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	items map[string]item
}

type item struct {
	value interface{}
}

func NewCache() *SafeMap {
	return &SafeMap{items: make(map[string]item)}
}

func (c *SafeMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.items)
}

// Size returns how many items are stored in the cache.
func (c *SafeMap) Size() int {
	c.RLock()
	defer c.RUnlock()
	return len(c.items)
}

func (c *SafeMap) get(key string) (item, bool) {
	c.RLock()
	cacheItem, ok := c.items[key]
	c.RUnlock()
	return cacheItem, ok
}

func (c *SafeMap) Exists(key string) bool {
	_, ok := c.get(key)
	return ok
}

// Get returns stored record.
// First returned: the stored value.
// Second returned: existence flag like in the map.
func (c *SafeMap) Get(key string) (interface{}, bool) {
	cacheItem, ok := c.get(key)
	return cacheItem.value, ok
}

// Set adds record in the cache.
func (c *SafeMap) Set(key string, value interface{}) {
	cacheItem := item{value: value}
	c.Lock()
	defer c.Unlock()
	c.items[key] = cacheItem
}

// Delete deletes the key and its value from the cache.
func (c *SafeMap) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.items, key)
}
