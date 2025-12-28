package http

import (
	"context"
)

// HTTPAwareClient is an interface that represents a client used for performing HTTP operations.
type HTTPAwareClient interface {
	// GetDefaultHeaders returns a map containing the headers that are included with every request by default.
	GetDefaultHeaders() map[string]string

	// Send sends an HTTP request using the specified method and URL. Returns a byte slice representing the response,
	// if any, as well as any error that was generated during the request process.
	Send(ctx context.Context, method HTTPMethod, url string) ([]byte, error)

	// SendWithBody performs the same operation as Send but allows for the inclusion of request body data.
	SendWithBody(ctx context.Context, method HTTPMethod, url string, body []byte) ([]byte, error)

	// SendWithHeaders sends an HTTP request using the specified method, URL, and override headers. The "override headers"
	// are intended to function as a map where headers for *this specific request* will be replaced if they already exist
	// or added if they are not already present. Returns a byte slice representing the response, if any, as well as any
	// error that was generated during the request process.
	SendWithHeaders(ctx context.Context, method HTTPMethod, url string, overrideHeaders map[string]string) ([]byte, error)

	// SendWithHeadersAndBody performs the same operation as SendWithHeaders but allows for the inclusion of request
	// body data.
	SendWithHeadersAndBody(ctx context.Context, method HTTPMethod, url string, overrideHeaders map[string]string, body []byte) ([]byte, error)

	// SetDefaultHeaders allows you to set the headers that will be included with every request by default.
	SetDefaultHeaders(headers map[string]string)
}
