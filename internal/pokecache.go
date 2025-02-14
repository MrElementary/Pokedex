package internal

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache_map map[string]cacheEntry
	mux       *sync.Mutex
}

// function that will create a cache to use when we run the program
// it also does nothing with the interval but passes it to reapLoop
// which will use it to execute an action every interval period
// in a separate routine
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache_map: make(map[string]cacheEntry),
		mux:       &sync.Mutex{},
	}

	// need explanation here
	go c.reapLoop(interval)

	return c
}

// function that will add an entry to the cache
func (c *Cache) Add(key string, value []byte) {
	// using the sync import to lock our Cache struct to lock access
	// since the maps are not thread-safe
	c.mux.Lock()
	defer c.mux.Unlock()

	c.cache_map[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
	// fmt.Println("added to cache")
}

// retrieves an entry from the cache_map if the key exists
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	return_entry, ok := c.cache_map[key]
	// fmt.Println("retrieved from cache")

	return return_entry.val, ok
}

// creates a recurring action with the reap function
// that is called every interval period for as long
// as the program is running.
func (c *Cache) reapLoop(interval time.Duration) {

	// creates a ticker for length of the interval period
	// for us it will be every 5 min as per the main() func
	ticker := time.NewTicker(interval)

	// every 5 min it calls reap to clean the cache
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	// removes entries older than 5 minutes.
	// e.g. if entry has value.createdAt of 13:03
	// and now.Add() is executed at 13:10.add(-5minutes) which is 13:05
	// it will see entry is older than 13:05 and remove it.
	for key, value := range c.cache_map {
		if value.createdAt.Before(now.Add(-last)) {
			delete(c.cache_map, key)
		}
	}
}
