package cache

import (
	"context"
	"time"
)

// Cache represents a generic interface for a caching mechanism.
type Cache interface {
	// Close closes the connection to the cache.
	Close() error

	// Delete removes the item associated with the given key from the cache.
	Delete(ctx context.Context, key string) (int64, error)

	// Exists checks if an item with the given key exists in the cache.
	Exists(ctx context.Context, key string) (bool, error)

	// Get retrieves the item associated with the given key from the cache.
	Get(ctx context.Context, key string) ([]byte, error)

	// Set stores the given value associated with the given key in the cache.
	Set(ctx context.Context, key string, value []byte) error

	// SetWithTTL stores the given value associated with the given key in the cache along with a time-to-live (TTL)
	// duration.
	SetWithTTL(ctx context.Context, key string, value []byte, ttl time.Duration) error
}
