package cache

// Cache represents a generic interface for a caching mechanism.
type Cache interface {
	// Delete removes the item associated with the given key from the cache.
	Delete(key string) error

	// Exists checks if an item with the given key exists in the cache.
	Exists(key string) (bool, error)

	// Get retrieves the item associated with the given key from the cache.
	Get(key string) ([]byte, error)

	// Set stores the given value associated with the given key in the cache.
	Set(key string, value []byte) error

	// SetWithTTL stores the given value associated with the given key in the cache along with a time-to-live (TTL)
	// in seconds.
	SetWithTTL(key string, value []byte, ttlSeconds int) error
}
