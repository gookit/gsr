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
	Get(key string) interface{}
	Set(key string, val interface{}, ttl time.Duration) (err error)
	Del(key string) error

	// GetMulti multi operation
	GetMulti(keys []string) map[string]interface{}
	SetMulti(values map[string]interface{}, ttl time.Duration) (err error)
	DelMulti(keys []string) error
}
