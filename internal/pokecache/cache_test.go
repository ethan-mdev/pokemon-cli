package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache()
	key := "testKey"
	value := []byte("testValue")
	cache.Add(key, value)
	cache.interval = 1 * time.Second // Set a short interval for testing

	// Test retrieval
	retrievedValue, exists := cache.Get(key)
	if !exists {
		t.Fatalf("Expected key %s to exist in cache", key)
	}
	if string(retrievedValue) != string(value) {
		t.Fatalf("Expected value %s, got %s", value, retrievedValue)
	}

	// Test expiration
	time.Sleep(6 * time.Second)
	_, exists = cache.Get(key)
	if exists {
		t.Fatalf("Expected key %s to be expired", key)
	}
}
