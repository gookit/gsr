package gsr

import (
	"io"
	"time"
)

// SimpleCacher interface definition
type SimpleCacher interface {
	// Closer close
	io.Closer
	// Clear clear
	Clear() error

	// Has basic operation
	Has(key string) bool
	Del(key string) error
	Get(key string) interface{}
	// GetAs get cache value and unmarshal as ptr.
	GetAs(key string, ptr interface{}) error
	Set(key string, val interface{}, ttl time.Duration) error

	// GetMulti multi operation
	GetMulti(keys []string) map[string]interface{}
	SetMulti(values map[string]interface{}, ttl time.Duration) error
	DelMulti(keys []string) error
}
