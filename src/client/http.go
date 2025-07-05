package client

import (
	"context"
	"net/http"
)

// HTTPMethod is a type used to represent the HTTP method that should be used when sending a request.
type HTTPMethod int

// String returns the string representation of the HTTP method.
func (hm HTTPMethod) String() string {
	switch hm {
	case HTTPMethodConnect:
		return http.MethodConnect
	case HTTPMethodDelete:
		return http.MethodDelete
	case HTTPMethodGet:
		return http.MethodGet
	case HTTPMethodHead:
		return http.MethodHead
	case HTTPMethodOptions:
		return http.MethodOptions
	case HTTPMethodPatch:
		return http.MethodPatch
	case HTTPMethodPost:
		return http.MethodPost
	case HTTPMethodPut:
		return http.MethodPut
	case HTTPMethodTrace:
		return http.MethodTrace
	}
	return ""
}

const (
	// HTTPMethodConnect represents a CONNECT request over HTTP.
	HTTPMethodConnect HTTPMethod = iota

	// HTTPMethodDelete represents a DELETE request over HTTP.
	HTTPMethodDelete

	// HTTPMethodGet represents a GET request over HTTP.
	HTTPMethodGet

	// HTTPMethodHead represents a HEAD request over HTTP.
	HTTPMethodHead

	// HTTPMethodOptions represents an OPTIONS request over HTTP.
	HTTPMethodOptions

	// HTTPMethodPatch represents a PATCH request over HTTP.
	HTTPMethodPatch

	// HTTPMethodPost represents a POST request over HTTP.
	HTTPMethodPost

	// HTTPMethodPut represents a PUT request over HTTP.
	HTTPMethodPut

	// HTTPMethodTrace represents a TRACE request over HTTP.
	HTTPMethodTrace
)

// HTTPAwareClient is an interface that represents a client used for performing HTTP operations.
type HTTPAwareClient interface {
	// GetDefaultHeaders returns a map containing the headers that are included with every request by default.
	GetDefaultHeaders()

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
