package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mut     sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time // a time.time that represents when the entry was created
	val       []byte    // a byte string that represents the raw data being cached
}

// returns a new cache
func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
	}

	// kick off a thread to reap the entries
	go cache.reapLoop(interval)

	// returns immediately while loop runs
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	// access the cache instance to add to it
	c.mut.Lock() // lock the mutex while making edits
	defer c.mut.Unlock()

	// create or edit the entry in the cache provided by the key
	c.entries[key] = cacheEntry{
		createdAt: time.Now(), // created right now
		val:       val,        // value of the entry is the []byte passed to the function
	}
}

func (c *Cache) Get(key string) (b []byte, found bool) {
	c.mut.Lock()
	defer c.mut.Unlock()

	data, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return data.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	// Infinite loop that:
	// 1. Waits for interval Duration
	// 2. Checks all entries
	// 3. Deletes entries older than interval
	// 4. Repeats forever
	for {
		// sleep until interval passes
		time.Sleep(interval)

		c.mut.Lock()

		// check entries
		for key, val := range c.entries {
			if time.Since(val.createdAt) > interval {
				delete(c.entries, key)
			}
		}
		c.mut.Unlock()
	}
}
