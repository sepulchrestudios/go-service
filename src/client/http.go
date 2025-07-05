package client

// HTTPMethod is a type used to represent the HTTP method that should be used when sending a request.
type HTTPMethod int

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

// NoRequestOverrideHeaders is an empty map that represents no override headers being provided in a request. This
// primarily exists for readability as well as allowing consuming logic to use the same method call but conditionally
// choose whether to supply headers.
var NoRequestOverrideHeaders = map[string]string{}

// NoRequestBody is an empty byte slice that represents no body data being provided in a request. This primarily exists
// for readability as well as allowing consuming logic to use the same method call but conditionally choose whether to
// supply body data.
var NoRequestBody = []byte{}

// HTTPAwareClient is an interface that represents a client used for performing HTTP operations.
type HTTPAwareClient interface {
	// Send sends an HTTP request using the specified method and URL. Returns a byte slice representing the response,
	// if any, as well as any error that was generated during the request process.
	Send(method HTTPMethod, url string) ([]byte, error)

	// SendWithBody performs the same operation as Send but allows for the inclusion of request body data.
	SendWithBody(method HTTPMethod, url string, body []byte) ([]byte, error)

	// SendWithHeaders sends an HTTP request using the specified method, URL, and override headers. The "override headers"
	// are intended to function as a map where headers for *this specific request* will be replaced if they already exist
	// or added if they are not already present. Returns a byte slice representing the response, if any, as well as any
	// error that was generated during the request process.
	SendWithHeaders(method HTTPMethod, url string, overrideHeaders map[string]string) ([]byte, error)

	// SendWithBody performs the same operation as SendWithHeaders but allows for the inclusion of request body data.
	SendWithHeadersAndBody(method HTTPMethod, url string, overrideHeaders map[string]string, body []byte) ([]byte, error)
}
