package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisConnectionArguments is a struct representing the general properties expected when making a connection
// to a Redis cache.
type RedisConnectionArguments struct {
	// CacheIdentifier in the embedded struct refers to the Redis database ID.
	CacheConnectionArguments

	Addr     string
	Password string
	Username string
}

// Redis represents a Redis caching mechanism.
type Redis struct {
	client *redis.Client
}

// NewRedis creates and returns a new Redis cache instance along with any error that may have occurred.
func NewRedis(ctx context.Context, connectionArguments *RedisConnectionArguments) (*Redis, error) {
	err := ValidateRedisConnectionArguments(connectionArguments)
	if err != nil {
		return nil, err
	}
	redisOptions := &redis.Options{
		Addr: connectionArguments.Addr,
	}
	databaseID, err := strconv.ParseInt(connectionArguments.CacheIdentifier, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrRedisCannotParseDatabaseIDAsInteger, err)
	}
	redisOptions.DB = int(databaseID)
	// Set credentials provider if username or password is provided because auth is optional
	if connectionArguments.Username != "" || connectionArguments.Password != "" {
		redisOptions.CredentialsProviderContext = func(ctx context.Context) (string, string, error) {
			return connectionArguments.Username, connectionArguments.Password, nil
		}
	}
	return NewRedisWithOptions(ctx, redisOptions)
}

// NewRedisWithOptions creates and returns a new Redis cache instance along with any error that may have occurred.
//
// The provided redis.Options pointer is used to configure the underlying Redis client fully and to give additional
// flexibility past what the NewRedis() function provides.
func NewRedisWithOptions(ctx context.Context, options *redis.Options) (*Redis, error) {
	client := redis.NewClient(options)
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotConnect, err)
	}
	return &Redis{
		client: client,
	}, nil
}

// ValidateRedisConnectionArguments takes a RedisConnectionArguments struct pointer and returns an error
// if any of the expected fields are missing. Returns nil if the validation checks pass.
func ValidateRedisConnectionArguments(connectionArguments *RedisConnectionArguments) error {
	if connectionArguments == nil {
		return ErrRedisNoConnectionArguments
	}
	if connectionArguments.CacheIdentifier == "" {
		return ErrNoCacheIdentifier
	}
	if connectionArguments.Addr == "" {
		return ErrRedisNoConnectionAddr
	}
	return nil
}

// Delete destroys the item associated with the given key from the cache. The integer return value indicates the number
// of items that were deleted.
func (r *Redis) Delete(ctx context.Context, key string) (int64, error) {
	if r == nil || r.client == nil {
		return 0, nil
	}
	return r.client.Del(ctx, key).Result()
}

// Exists checks if an item with the given key exists in the cache.
func (r *Redis) Exists(ctx context.Context, key string) (bool, error) {
	if r == nil || r.client == nil {
		return false, nil
	}
	result, err := r.client.Exists(ctx, key).Result()
	return result > 0, err
}

// Get retrieves the item associated with the given key from the cache. If the key could not be found, this method
// returns nil.
func (r *Redis) Get(ctx context.Context, key string) ([]byte, error) {
	if r == nil || r.client == nil {
		return nil, nil
	}
	result, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	return result, err
}

// Set stores the given value associated with the given key in the cache.
func (r *Redis) Set(ctx context.Context, key string, value []byte) error {
	if r == nil || r.client == nil {
		return nil
	}
	return r.SetWithTTL(ctx, key, value, 0)
}

// SetWithTTL stores the given value associated with the given key in the cache along with a time-to-live (TTL)
// duration.
func (r *Redis) SetWithTTL(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	if r == nil || r.client == nil {
		return nil
	}
	return r.client.Set(ctx, key, value, ttl).Err()
}
