package http

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

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

// HTTPClient represents a struct that provides basic HTTP client connectivity capabilities.
type HTTPClient struct {
	client  *http.Client
	headers map[string]string
	mu      sync.Mutex
}

// MakeDefaultNetHTTPClient builds and returns a new default http.Client pointer instance that can be used immediately
// when creating HTTPClient instances.
func MakeDefaultNetHTTPClient() *http.Client {
	return &http.Client{}
}

// MakeEmptyRequestBody returns an empty byte slice that represents no body data being provided in a request. This
// primarily exists for readability as well as allowing consuming logic to use the same method call but conditionally
// choose whether to supply body data.
func MakeEmptyRequestBody() []byte {
	return []byte{}
}

// MakeEmptyRequestHeaders returns an empty map that represents no override headers being provided in a request. This
// primarily exists for readability as well as allowing consuming logic to use the same method call but conditionally
// choose whether to supply headers.
func MakeEmptyRequestHeaders() map[string]string {
	return map[string]string{}
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
		client:  client,
		headers: map[string]string{},
	}, nil
}

// Send sends an HTTP request using the specified method and URL. Returns a byte slice representing the response,
// if any, as well as any error that was generated during the request process.
func (hc *HTTPClient) Send(ctx context.Context, method HTTPMethod, url string) ([]byte, error) {
	if hc == nil {
		return []byte{}, ErrHTTPClientCannotBeNil
	}
	return hc.SendWithHeadersAndBody(ctx, method, url, MakeEmptyRequestHeaders(), MakeEmptyRequestBody())
}

// SendWithBody performs the same operation as Send but allows for the inclusion of request body data.
func (hc *HTTPClient) SendWithBody(ctx context.Context, method HTTPMethod, url string, body []byte) ([]byte, error) {
	if hc == nil {
		return []byte{}, ErrHTTPClientCannotBeNil
	}
	return hc.SendWithHeadersAndBody(ctx, method, url, MakeEmptyRequestHeaders(), body)
}

// SendWithHeaders sends an HTTP request using the specified method, URL, and override headers. The "override headers"
// are intended to function as a map where headers for *this specific request* will be replaced if they already exist
// or added if they are not already present. Returns a byte slice representing the response, if any, as well as any
// error that was generated during the request process.
func (hc *HTTPClient) SendWithHeaders(
	ctx context.Context, method HTTPMethod, url string, overrideHeaders map[string]string,
) ([]byte, error) {
	if hc == nil {
		return []byte{}, ErrHTTPClientCannotBeNil
	}
	return hc.SendWithHeadersAndBody(ctx, method, url, overrideHeaders, MakeEmptyRequestBody())
}

// SendWithHeadersAndBody performs the same operation as SendWithHeaders but allows for the inclusion of request
// body data.
func (hc *HTTPClient) SendWithHeadersAndBody(
	ctx context.Context, method HTTPMethod, url string, overrideHeaders map[string]string, body []byte,
) ([]byte, error) {
	if hc == nil {
		return []byte{}, ErrHTTPClientCannotBeNil
	}
	if hc.client == nil {
		return []byte{}, ErrNetHTTPClientCannotBeNil
	}
	if body == nil {
		body = []byte{}
	}

	// build the request based on the passed context
	req, err := http.NewRequestWithContext(ctx, method.String(), url, bytes.NewReader(body))
	if err != nil {
		return []byte{}, fmt.Errorf("%w: %w", ErrCannotCreateHTTPRequest, err)
	}

	// force the default headers from the map if we have any
	if len(hc.headers) > 0 {
		for defaultHeaderKey, defaultHeaderValue := range hc.headers {
			req.Header.Set(defaultHeaderKey, defaultHeaderValue)
		}
	}

	// set the request-specific headers for this go-around
	if len(overrideHeaders) > 0 {
		for headerKey, headerValue := range overrideHeaders {
			req.Header.Set(headerKey, headerValue)
		}
	}

	// send the request
	resp, err := hc.client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("%w: %w", ErrHTTPRequestFailed, err)
	}

	// read the response
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("%w: %w", ErrCannotReadHTTPResponseBody, err)
	}
	return respBody, nil
}

// GetDefaultHeaders returns a map containing the headers that are included with every request by default.
func (hc *HTTPClient) GetDefaultHeaders() map[string]string {
	if hc == nil || hc.headers == nil {
		return map[string]string{}
	}
	// ensure we don't get a collision if two or more goroutines try to read concurrently
	hc.mu.Lock()
	defer hc.mu.Unlock()
	headersCopy := map[string]string{}
	for key, val := range hc.headers {
		headersCopy[key] = val
	}
	return headersCopy
}

// SetDefaultHeaders allows you to set the headers that will be included with every request by default.
func (hc *HTTPClient) SetDefaultHeaders(headers map[string]string) {
	if hc == nil {
		return
	}
	// ensure we don't get a collision if two or more goroutines try to write concurrently
	hc.mu.Lock()
	defer hc.mu.Unlock()
	newHeaders := map[string]string{}
	if len(headers) > 0 {
		for key, val := range headers {
			newHeaders[key] = val
		}
	}
	hc.headers = newHeaders
}
