package gsr

import (
	"io"
	"time"
)

// SimpleCacher interface definition
type SimpleCacher interface {
	// Closer close cache handle
	io.Closer
	// Clear clear all caches
	Clear() error

	// Has basic operation
	Has(key string) bool
	Del(key string) error
	Get(key string) interface{}
	Set(key string, val interface{}, ttl time.Duration) error

	// GetMulti multi operation
	GetMulti(keys []string) map[string]interface{}
	SetMulti(values map[string]interface{}, ttl time.Duration) error
	DelMulti(keys []string) error
}

// CodedCacher interface.
type CodedCacher interface {
	SimpleCacher

	// GetAs get and decode cache value to object ptr
	GetAs(key string, ptr interface{}) error
}
