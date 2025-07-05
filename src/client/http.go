package client

type HTTPClient struct {
}

// NewHTTPClient builds and returns a new HTTPClient pointer instance as well as any error that may have occurred
// during the creation.
func NewHTTPClient() (*HTTPClient, error) {
	return &HTTPClient{}, nil
}
