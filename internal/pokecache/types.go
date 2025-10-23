package pokecache

import (
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
