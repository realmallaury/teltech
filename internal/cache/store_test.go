package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewStore(t *testing.T) {
	assert := assert.New(t)

	store := NewStore(1, 1*time.Second)
	assert.NotNil(store)
	assert.Equal(1, store.cache.cacheSize)
	assert.Equal(1*time.Second, store.cache.recordTTL)
}

func TestStoreRecord(t *testing.T) {
	assert := assert.New(t)

	store := NewStore(100, 100*time.Millisecond)

	store.StoreRecord("1", 2)
	assert.Equal(1, store.cache.ll.Len())
	assert.Equal(2, store.cache.cache["1"].Value.(*entry).value)

	store.StoreRecord("3", 4)
	assert.Equal(2, store.cache.ll.Len())
	assert.Equal(4, store.cache.cache["3"].Value.(*entry).value)

	store.StoreRecord("5", 6)
	assert.Equal(3, store.cache.ll.Len())
	assert.Equal(6, store.cache.cache["5"].Value.(*entry).value)
}

func TestGetRecord(t *testing.T) {
	assert := assert.New(t)

	store := NewStore(100, 100*time.Millisecond)

	store.StoreRecord("1", 2)
	value, ok := store.GetRecord("1")
	assert.Equal(2, value)
	assert.True(ok)
}

func TestRecordTTL(t *testing.T) {
	assert := assert.New(t)

	store := NewStore(100, 200*time.Millisecond)
	store.StoreRecord("1", 2)

	time.Sleep(100 * time.Millisecond)

	value, ok := store.GetRecord("1")
	assert.Equal(2, value)
	assert.True(ok)

	time.Sleep(200 * time.Millisecond)

	_, ok = store.GetRecord("1")
	assert.False(ok)
}

func TestCacheSize(t *testing.T) {
	assert := assert.New(t)

	store := NewStore(1, 100*time.Millisecond)
	store.cache.Add("1", 2)

	value, ok := store.GetRecord("1")
	assert.Equal(2, value)
	assert.True(ok)

	store.StoreRecord("2", 3)
	value, ok = store.GetRecord("2")
	assert.Equal(3, value)
	assert.True(ok)

	_, ok = store.GetRecord("1")
	assert.False(ok)
}
