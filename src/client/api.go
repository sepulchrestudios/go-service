package client

import "errors"

// ErrHTTPClientCannotBeNil is a sentinel error representing an attempt to use a nil HTTPClient pointer.
var ErrHTTPClientCannotBeNil = errors.New("HTTPClient instance cannot be nil")

type HTTPAPIClient struct {
	httpclient *HTTPClient
}

// NewHTTPAPIClient builds and returns a new HTTPAPIClient pointer instance as well as any error that may have occurred
// during the creation.
func NewHTTPAPIClient(httpclient *HTTPClient) (*HTTPAPIClient, error) {
	if httpclient == nil {
		return nil, ErrHTTPClientCannotBeNil
	}
	return &HTTPAPIClient{
		httpclient: httpclient,
	}, nil
}
