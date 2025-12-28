package http

import "net/http"

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
