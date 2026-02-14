package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Minute)

	key := "test-url"
	val := []byte("test data")

	cache.Add(key, val)

	cacheData, ok := cache.Get(key)
	if !ok {
		t.Errorf("Expected to find key in Cache")
	}

	if string(cacheData) != string(val) {
		t.Errorf("Expected %s, got %s", val, cacheData)
	}
}
