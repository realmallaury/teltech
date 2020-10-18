package cache

import (
	"container/list"
	"sync"
	"time"
)

// Cache is an LRU cache.
type Cache struct {
	// cacheSize is the maximum number of cache entries before an item is evicted.
	cacheSize int

	// recordTTL is duration of record inactivity past which is evicted.
	recordTTL time.Duration

	ll    *list.List
	cache map[interface{}]*list.Element
	mux   sync.Mutex
}

// Key can be any comparable type.
type Key interface{}

type entry struct {
	key   Key
	ttl   time.Time
	value interface{}
}

// New creates a new Cache instance,
// zero values are 1000 entries and 1 min cache element TTL.
func New(cacheSize int, recordTTL time.Duration) *Cache {
	if cacheSize == 0 {
		cacheSize = 1000
	}

	if recordTTL == 0 {
		recordTTL = 1 * time.Minute
	}

	return &Cache{
		cacheSize: cacheSize,
		recordTTL: recordTTL,
		ll:        list.New(),
		cache:     make(map[interface{}]*list.Element),
	}
}

// Add adds a value to the cache.
func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil {
		return
	}

	c.mux.Lock()
	defer c.mux.Unlock()
	c.checkSizeAndExpiredRecords()

	if element, ok := c.cache[key]; ok {
		c.ll.MoveToFront(element)
		element.Value.(*entry).ttl = time.Now().Add(c.recordTTL)
		element.Value.(*entry).value = value
		return
	}

	element := c.ll.PushFront(&entry{key, time.Now().Add(c.recordTTL), value})
	c.cache[key] = element
}

// Get looks up a key's value from the cache.
func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	if c.cache == nil {
		return
	}

	c.mux.Lock()
	defer c.mux.Unlock()
	c.checkSizeAndExpiredRecords()

	if element, hit := c.cache[key]; hit {
		element.Value.(*entry).ttl = time.Now().Add(c.recordTTL)
		c.ll.MoveToFront(element)
		return element.Value.(*entry).value, true
	}

	return
}

// checkSizeAndExpiredRecords checks cache size and remove elements from the back,
// removes expired records from the cache.
func (c *Cache) checkSizeAndExpiredRecords() {
	element := c.ll.Back()
	for element != nil {
		entry := element.Value.(*entry)
		if time.Now().After(entry.ttl) {
			c.ll.Remove(element)
			delete(c.cache, entry.key)
		}

		element = element.Prev()
	}

	element = c.ll.Back()
	for element != nil && c.ll.Len() > c.cacheSize {
		entry := element.Value.(*entry)
		c.ll.Remove(element)
		delete(c.cache, entry.key)

		element = element.Prev()
	}
}
