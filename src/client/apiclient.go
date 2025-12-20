package client

import "context"

// HTTPAPIClient represents a struct that provides a standard way to invoke API calls on services over HTTP.
type HTTPAPIClient struct {
	httpclient HTTPAwareClient
}

// NewHTTPAPIClient builds and returns a new HTTPAPIClient pointer instance as well as any error that may have occurred
// during creation.
func NewHTTPAPIClient(httpclient HTTPAwareClient) (*HTTPAPIClient, error) {
	if httpclient == nil {
		return nil, ErrHTTPClientCannotBeNil
	}
	return &HTTPAPIClient{
		httpclient: httpclient,
	}, nil
}

// Send sends an HTTP request using the specified method and URL. Returns a byte slice representing the response,
// if any, as well as any error that was generated during the request process.
func (hc *HTTPAPIClient) Send(ctx context.Context, method HTTPMethod, url string) ([]byte, error) {
	if hc == nil {
		return []byte{}, ErrHTTPClientCannotBeNil
	}
	return hc.httpclient.Send(ctx, method, url)
}

// SendWithBody performs the same operation as Send but allows for the inclusion of request body data.
func (hc *HTTPAPIClient) SendWithBody(ctx context.Context, method HTTPMethod, url string, body []byte) ([]byte, error) {
	if hc == nil {
		return []byte{}, ErrHTTPClientCannotBeNil
	}
	return hc.httpclient.SendWithBody(ctx, method, url, body)
}

// SendWithHeaders sends an HTTP request using the specified method, URL, and override headers. The "override headers"
// are intended to function as a map where headers for *this specific request* will be replaced if they already exist
// or added if they are not already present. Returns a byte slice representing the response, if any, as well as any
// error that was generated during the request process.
func (hc *HTTPAPIClient) SendWithHeaders(
	ctx context.Context, method HTTPMethod, url string, overrideHeaders map[string]string,
) ([]byte, error) {
	if hc == nil {
		return []byte{}, ErrHTTPClientCannotBeNil
	}
	return hc.httpclient.SendWithHeaders(ctx, method, url, overrideHeaders)
}

// SendWithHeadersAndBody performs the same operation as SendWithHeaders but allows for the inclusion of request
// body data.
func (hc *HTTPAPIClient) SendWithHeadersAndBody(
	ctx context.Context, method HTTPMethod, url string, overrideHeaders map[string]string, body []byte,
) ([]byte, error) {
	if hc == nil {
		return []byte{}, ErrHTTPClientCannotBeNil
	}
	return hc.httpclient.SendWithHeadersAndBody(ctx, method, url, overrideHeaders, body)
}

// GetDefaultHeaders returns a map containing the headers that are included with every request by default.
func (hc *HTTPAPIClient) GetDefaultHeaders() map[string]string {
	if hc == nil {
		return map[string]string{}
	}
	return hc.httpclient.GetDefaultHeaders()
}

// SetDefaultHeaders allows you to set the headers that will be included with every request by default.
func (hc *HTTPAPIClient) SetDefaultHeaders(headers map[string]string) {
	if hc == nil {
		return
	}
	hc.httpclient.SetDefaultHeaders(headers)
}
