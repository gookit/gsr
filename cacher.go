package gsr

import (
	"context"
	"io"
	"time"
)

// SimpleCacher interface definition
type SimpleCacher interface {
	// Closer close cache handle
	io.Closer
	// Clear all cache data
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

// ContextCacher interface.
type ContextCacher interface {
	SimpleCacher
	// WithContext and clone new cacher
	WithContext(ctx context.Context) ContextCacher
}

// ContextOpCacher interface.
type ContextOpCacher interface {
	SimpleCacher

	// HasWithCtx basic operation
	HasWithCtx(ctx context.Context, key string) bool
	DelWithCtx(ctx context.Context, key string) error
	GetWithCtx(ctx context.Context, key string) interface{}
	SetWithCtx(ctx context.Context, key string, val interface{}, ttl time.Duration) error

	// MGetWithCtx multi keys operation
	MGetWithCtx(ctx context.Context, keys []string) map[string]interface{}
	MSetWithCtx(ctx context.Context, values map[string]interface{}, ttl time.Duration) error
	MDelWithCtx(ctx context.Context, keys []string) error
}

// CodedCacher interface.
type CodedCacher interface {
	SimpleCacher

	// GetAs get and decode cache value to object ptr
	GetAs(key string, ptr interface{}) error
}
