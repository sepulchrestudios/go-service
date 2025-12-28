package cache

// CacheConnectionArguments is a struct representing the general properties expected when making a connection
// to a cache.
type CacheConnectionArguments struct {
	// CacheIdentifier is the identifier of the cache instance depending on the implementation.
	CacheIdentifier string
}
