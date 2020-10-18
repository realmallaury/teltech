package cache

import (
	"time"
)

// Store describes cache operations.
type Store interface {
	StoreRecord(key string, value interface{})
	GetRecord(key string) (interface{}, bool)
}

// InMemoryStore is in memory implementation of cache store.
type InMemoryStore struct {
	cache *Cache
}

// StoreRecord stores record to cache.
func (i *InMemoryStore) StoreRecord(key string, value interface{}) {
	i.cache.Add(key, value)
}

// GetRecord gets record from cache.
func (i *InMemoryStore) GetRecord(key string) (interface{}, bool) {
	return i.cache.Get(key)
}

// NewStore returns new in memory cache store instance.
func NewStore(cacheSize int, recordTTL time.Duration) *InMemoryStore {
	return &InMemoryStore{
		cache: New(cacheSize, recordTTL),
	}
}
