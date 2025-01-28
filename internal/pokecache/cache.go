package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]CacheEntry
	mu      sync.Mutex
	done    chan bool
	closed  bool
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: map[string]CacheEntry{},
		mu:      sync.Mutex{},
		done:    make(chan bool),
		closed:  false,
	}

	go cache.reapLoop(interval)

	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if !cache.closed {
		newEntry := CacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
		cache.entries[key] = newEntry
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if !cache.closed {
		entry, ok := cache.entries[key]
		if !ok {
			return []byte{}, false
		}
		return entry.val, true
	}
	return []byte{}, false
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-cache.done:
			return
		case <-ticker.C:
			cache.mu.Lock()

			for key, entry := range cache.entries {
				cutoffTime := entry.createdAt.Add(interval)
				if cutoffTime.Before(time.Now()) {
					// Remove entry
					delete(cache.entries, key)
				}
			}

			cache.mu.Unlock()
		}
	}
}

func (cache *Cache) Close() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if !cache.closed {
		cache.done <- true
		close(cache.done)
		cache.closed = true
	}
}
