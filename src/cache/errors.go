package cache

import "errors"

// ErrCannotConnect is a sentinel error describing a failure to connect to the cache.
var ErrCannotConnect = errors.New("cannot connect to cache")

// ErrNoCacheIdentifier is a sentinel error representing a blank cache identifier when attempting to connect
// to the cache.
var ErrNoCacheIdentifier = errors.New("cache identifier cannot be blank")
