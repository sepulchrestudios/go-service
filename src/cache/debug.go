package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/sepulchrestudios/go-service/src/log"
	"go.uber.org/zap"
)

// Debug represents a struct that wraps a caching mechanism with debug logging capabilities.
type Debug struct {
	implementation Contract
	logger         log.DebugContract
}

// CacheDebugAction represents the type of action being logged for cache operations.
type CacheDebugAction string

const (
	// CacheDebugActionRequest represents a cache request.
	CacheDebugActionRequest CacheDebugAction = "request"

	// CacheDebugActionResponse represents a cache response.
	CacheDebugActionResponse CacheDebugAction = "response"
)

// CacheDebugOperation represents the type of cache operation being performed.
type CacheDebugOperation string

const (
	// CacheDebugOperationClose represents a close operation.
	CacheDebugOperationClose CacheDebugOperation = "close"

	// CacheDebugOperationDelete represents a delete operation.
	CacheDebugOperationDelete CacheDebugOperation = "delete"

	// CacheDebugOperationExists represents an exists operation.
	CacheDebugOperationExists CacheDebugOperation = "exists"

	// CacheDebugOperationGet represents a get operation.
	CacheDebugOperationGet CacheDebugOperation = "get"

	// CacheDebugOperationSet represents a set operation.
	CacheDebugOperationSet CacheDebugOperation = "set"

	// CacheDebugOperationSetWithTTL represents a set-with-TTL operation.
	CacheDebugOperationSetWithTTL CacheDebugOperation = "setwithttl"
)

// logAction logs a cache action with the specified operation and any additional fields.
func (d *Debug) logAction(action CacheDebugAction, operation CacheDebugOperation, fields ...zap.Field) {
	if d == nil || d.implementation == nil || d.logger == nil {
		return
	}
	allFields := []zap.Field{
		zap.String("operation", string(operation)),
	}
	allFields = append(allFields, fields...)
	d.logger.Debug(fmt.Sprintf("cache action [%s]", action), allFields...)
}

// Close closes the connection to the cache.
func (d *Debug) Close() error {
	if d == nil || d.implementation == nil {
		return nil
	}
	d.logAction(CacheDebugActionRequest, CacheDebugOperationClose)
	err := d.implementation.Close()
	d.logAction(CacheDebugActionResponse, CacheDebugOperationClose,
		zap.Error(err))
	return err
}

// Delete destroys the item associated with the given key from the cache. The integer return value indicates the number
// of items that were deleted.
func (d *Debug) Delete(ctx context.Context, key string) (int64, error) {
	if d == nil || d.implementation == nil {
		return 0, nil
	}
	d.logAction(CacheDebugActionRequest, CacheDebugOperationDelete,
		zap.String("key", key))
	count, err := d.implementation.Delete(ctx, key)
	d.logAction(CacheDebugActionResponse, CacheDebugOperationDelete,
		zap.String("key", key), zap.Any("value", count), zap.Error(err))
	return count, err
}

// Exists checks if an item with the given key exists in the cache.
func (d *Debug) Exists(ctx context.Context, key string) (bool, error) {
	if d == nil || d.implementation == nil {
		return false, nil
	}
	d.logAction(CacheDebugActionRequest, CacheDebugOperationExists,
		zap.String("key", key))
	exists, err := d.implementation.Exists(ctx, key)
	d.logAction(CacheDebugActionResponse, CacheDebugOperationExists,
		zap.String("key", key), zap.Any("value", exists), zap.Error(err))
	return exists, err
}

// Get retrieves the item associated with the given key from the cache. If the key could not be found, this method
// returns nil.
func (d *Debug) Get(ctx context.Context, key string) ([]byte, error) {
	if d == nil || d.implementation == nil {
		return nil, nil
	}
	d.logAction(CacheDebugActionRequest, CacheDebugOperationGet,
		zap.String("key", key))
	value, err := d.implementation.Get(ctx, key)
	d.logAction(CacheDebugActionResponse, CacheDebugOperationGet,
		zap.String("key", key), zap.Any("value", value), zap.Error(err))
	return value, err
}

// Set stores the given value associated with the given key in the cache.
func (d *Debug) Set(ctx context.Context, key string, value []byte) error {
	if d == nil || d.implementation == nil {
		return nil
	}
	d.logAction(CacheDebugActionRequest, CacheDebugOperationSet,
		zap.String("key", key), zap.Any("value", value))
	err := d.implementation.Set(ctx, key, value)
	d.logAction(CacheDebugActionResponse, CacheDebugOperationSet,
		zap.String("key", key), zap.Error(err))
	return err
}

// SetWithTTL stores the given value associated with the given key in the cache along with a time-to-live (TTL)
// duration.
func (d *Debug) SetWithTTL(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	if d == nil || d.implementation == nil {
		return nil
	}
	d.logAction(CacheDebugActionRequest, CacheDebugOperationSetWithTTL,
		zap.String("key", key), zap.Any("value", value), zap.Duration("ttl", ttl))
	err := d.implementation.SetWithTTL(ctx, key, value, ttl)
	d.logAction(CacheDebugActionResponse, CacheDebugOperationSetWithTTL,
		zap.String("key", key), zap.Error(err))
	return err
}

// NewDebug takes an existing cache implementation and a debug-level logger then returns a wrapper cache that
// provides the existing implementation with debug logging capabilities.
func NewDebug(implementation Contract, logger log.DebugContract) *Debug {
	return &Debug{
		implementation: implementation,
		logger:         logger,
	}
}
