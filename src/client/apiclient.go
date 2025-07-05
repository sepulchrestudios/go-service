package client

// HTTPAPIClient represents a struct that provides a standard way to invoke API calls on services over HTTP.
type HTTPAPIClient struct {
	httpclient *HTTPAwareClient
}

// NewHTTPAPIClient builds and returns a new HTTPAPIClient pointer instance as well as any error that may have occurred
// during creation.
func NewHTTPAPIClient(httpclient *HTTPAwareClient) (*HTTPAPIClient, error) {
	if httpclient == nil {
		return nil, ErrHTTPClientCannotBeNil
	}
	return &HTTPAPIClient{
		httpclient: httpclient,
	}, nil
}
