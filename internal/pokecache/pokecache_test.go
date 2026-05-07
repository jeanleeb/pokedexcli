package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	c := NewCache(1 * time.Hour)

	c.Add("key", []byte("hello"))
	val, exists := c.Get("key")

	if !exists {
		t.Errorf("Expected key to exist in cache, but it does not")
	}
	if string(val) != "hello" {
		t.Errorf("Expected value to be 'hello', but got '%s'", string(val))
	}
}

func TestGetMissing(t *testing.T) {
	c := NewCache(1 * time.Hour)

	val, exists := c.Get("missing")

	if exists {
		t.Errorf("Expected key to not exist in cache, but it does")
	}
	if val != nil {
		t.Errorf("Expected value to be nil for missing key, but got '%s'", string(val))
	}
}

func TestReapDeletesExpired(t *testing.T) {
	interval := 10 * time.Millisecond
	c := NewCache(interval)
	c.Add("key", []byte("hello"))

	time.Sleep(50 * time.Millisecond)

	val, ok := c.Get("key")
	if ok {
		t.Errorf("Expected key to be expired and removed from cache, but it still exists")
	}
	if val != nil {
		t.Errorf("Expected value to be nil for expired key, but got '%s'", string(val))
	}
}

func TestReapPreservesUnexpired(t *testing.T) {
	const value = "hello"
	const interval = 10 * time.Minute
	c := NewCache(interval)
	c.Add("key", []byte(value))

	time.Sleep(50 * time.Millisecond)

	val, ok := c.Get("key")
	if !ok {
		t.Errorf("Expected key to still exist in cache, but it does not")
	}
	if string(val) != value {
		t.Errorf("Expected value to be 'hello', but got '%s'", string(val))
	}
}
