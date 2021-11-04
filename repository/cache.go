package repository

import (
"time"

"github.com/ReneKroon/ttlcache/v2"
)

type CacheStore struct {
	cache *ttlcache.Cache
}

//NewCacheStore constructor: returns instance of CacheStore
func NewCacheStore(c *ttlcache.Cache) *CacheStore {
	return &CacheStore{
		cache: c,
	}
}

// SetTTL sets global TTL expiration of keys inside cache store
func (c *CacheStore) SetTTL(ttl time.Duration) {
	c.cache.SetTTL(ttl)
}

// Get retrieves key from cache store
func (c *CacheStore) Get(key string) (interface{}, bool) {
	if value, exists := c.cache.Get(key); exists == nil {
		return value, true
	}
	return nil, false
}

// Set adds key and data into cache store
func (c *CacheStore) Set(key string, data interface{}) error {
	err := c.cache.Set(key, data)
	if err != nil {
		return err
	}
	return nil
}

// SetCacheSizeLimit sets the max number of items stored in cache
func (c *CacheStore) SetCacheSizeLimit(limit int) {
	c.cache.SetCacheSizeLimit(limit)
}

// Count returns the number of items in the cache
func (c *CacheStore) Count() int {
	return c.cache.Count()
}

