package http

import "errors"

// ErrCannotCreateHTTPRequest is a sentinel error representing a failure while creating an HTTP request instance.
var ErrCannotCreateHTTPRequest = errors.New("cannot create http request")

// ErrCannotReadHTTPResponseBody is a sentinel error representing a failure while reading the body of an HTTP response.
var ErrCannotReadHTTPResponseBody = errors.New("cannot read http response body")

// ErrHTTPClientCannotBeNil is a sentinel error representing an attempt to use a nil HTTPClient pointer.
var ErrHTTPClientCannotBeNil = errors.New("HTTPClient instance cannot be nil")

// ErrHTTPRequestFailed is a sentinel error representing a failure when sending an HTTP request itself.
var ErrHTTPRequestFailed = errors.New("http request failed")

// ErrNetHTTPClientCannotBeNil is a sentinel error representing an attempt to use a nil http.Client pointer.
var ErrNetHTTPClientCannotBeNil = errors.New("http.Client instance cannot be nil")
