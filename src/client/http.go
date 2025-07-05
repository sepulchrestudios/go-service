package client

import (
	"errors"
	"net/http"
)

// ErrNetHTTPClientCannotBeNil is a sentinel error representing an attempt to use a nil http.Client pointer.
var ErrNetHTTPClientCannotBeNil = errors.New("http.Client instance cannot be nil")

// HTTPClient represents a struct that provides basic HTTP client connectivity capabilities.
type HTTPClient struct {
	client *http.Client
}

// MakeDefaultNetHTTPClient builds and returns a new default http.Client pointer instance that can be used immediately
// when creating HTTPClient instances.
func MakeDefaultNetHTTPClient() *http.Client {
	return &http.Client{}
}

// NewHTTPClient builds and returns a new HTTPClient pointer instance as well as any error that may have occurred
// during creation.
func NewHTTPClient() (*HTTPClient, error) {
	return NewHTTPClientFromClient(MakeDefaultNetHTTPClient())
}

// NewHTTPClientFromClient builds and returns a new HTTPClient pointer instance based on the provided http.Client pointer
// as well as any error that may have occurred during creation.
func NewHTTPClientFromClient(client *http.Client) (*HTTPClient, error) {
	if client == nil {
		return nil, ErrNetHTTPClientCannotBeNil
	}
	return &HTTPClient{
		client: client,
	}, nil
}
