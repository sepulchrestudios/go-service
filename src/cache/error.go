package cache

import "errors"

// ErrCannotConnect is a sentinel error describing a failure to connect to the cache.
var ErrCannotConnect = errors.New("cannot connect to cache")

// ErrNoCacheIdentifier is a sentinel error representing a blank cache identifier when attempting to connect
// to the cache.
var ErrNoCacheIdentifier = errors.New("cache identifier cannot be blank")

// ErrRedisCannotParseDatabaseIDAsInteger is a sentinel error describing a failure to parse the Redis database ID as
// an integer.
var ErrRedisCannotParseDatabaseIDAsInteger = errors.New("cannot parse Redis database ID as integer")

// ErrRedisNoConnectionDatabaseAddr is a sentinel error representing a blank address string when attempting to make
// a Redis cache connection.
var ErrRedisNoConnectionAddr = errors.New("address in Redis connection arguments cannot be blank")

// ErrRedisNoConnectionArguments is a sentinel error representing a nil connection arguments pointer when attempting
// to make a Redis cache connection.
var ErrRedisNoConnectionArguments = errors.New("connection arguments for Redis cannot be nil")
