package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

const (
	ResourceURLToken             = "token"
	ResourceURLRefresh           = "refresh"
	ResourceURLRegistries        = "registries"
	ResourceURLRepositories      = "repositories"
	ResourceURLGarbageCollection = "garbage-collection"
	ResourceURLSize              = "size"
	ResourceURLImages            = "images"
	ResourceURLTags              = "tags"
)

const (
	// appName represents an application name.
	appName = "craas-go"

	// appVersion is a version of the application.
	appVersion = "0.1.0"

	// userAgent contains a basic user agent that will be used in queries.
	userAgent = appName + "/" + appVersion

	// defaultHTTPTimeout represents the default timeout (in seconds) for HTTP requests.
	defaultHTTPTimeout = 120

	// defaultDialTimeout represents the default timeout (in seconds) for HTTP connection establishments.
	defaultDialTimeout = 60

	// defaultKeepaliveTimeout represents the default keep-alive period for an active network connection.
	defaultKeepaliveTimeout = 60

	// defaultMaxIdleConns represents the maximum number of idle (keep-alive) connections.
	defaultMaxIdleConns = 100

	// defaultIdleConnTimeout represents the maximum amount of time an idle (keep-alive) connection will remain
	// idle before closing itself.
	defaultIdleConnTimeout = 100

	// defaultTLSHandshakeTimeout represents the default timeout (in seconds) for TLS handshake.
	defaultTLSHandshakeTimeout = 60

	// defaultExpectContinueTimeout represents the default amount of time to wait for a server's first
	// response headers.
	defaultExpectContinueTimeout = 1
)

const errGotHTTPStatusCodeFmt = "craas-go: got the %d status code from the server"

// ServiceClient stores details that are needed to work with Selectel CRaaS API.
type ServiceClient struct {
	// HTTPClient represents an initialized HTTP client that will be used to do requests.
	HTTPClient *http.Client

	// Token is a client authentication token.
	Token string

	// Endpoint represents an endpoint that will be used in all requests.
	Endpoint string

	// UserAgent contains user agent that will be used in all requests.
	UserAgent string
}

// NewCRaaSClientV1 initializes a new CRaaS client for the V1 API.
func NewCRaaSClientV1(token, endpoint string) *ServiceClient {
	return &ServiceClient{
		HTTPClient: newHTTPClient(),
		Token:      token,
		Endpoint:   endpoint,
		UserAgent:  userAgent,
	}
}

// NewCRaaSClientV1WithCustomHTTP initializes a new CRaaS client for the V1 API using custom HTTP client.
// If custom HTTP client is nil - default HTTP client will be used.
func NewCRaaSClientV1WithCustomHTTP(customHTTPClient *http.Client, tokenID, endpoint string) *ServiceClient {
	if customHTTPClient == nil {
		customHTTPClient = newHTTPClient()
	}

	return &ServiceClient{
		HTTPClient: customHTTPClient,
		Token:      tokenID,
		Endpoint:   endpoint,
		UserAgent:  userAgent,
	}
}

// newHTTPClient returns a reference to an initialized and configured HTTP client.
func newHTTPClient() *http.Client {
	return &http.Client{
		Timeout:   defaultHTTPTimeout * time.Second,
		Transport: newHTTPTransport(),
	}
}

// newHTTPTransport returns a reference to an initialized and configured HTTP transport.
func newHTTPTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   defaultDialTimeout * time.Second,
			KeepAlive: defaultKeepaliveTimeout * time.Second,
		}).DialContext,
		MaxIdleConns:          defaultMaxIdleConns,
		IdleConnTimeout:       defaultIdleConnTimeout * time.Second,
		TLSHandshakeTimeout:   defaultTLSHandshakeTimeout * time.Second,
		ExpectContinueTimeout: defaultExpectContinueTimeout * time.Second,
	}
}

// DoRequest performs the HTTP request with the current ServiceClient's HTTPClient.
// Authentication and optional headers will be added automatically.
func (client *ServiceClient) DoRequest(ctx context.Context, method, path string, body io.Reader) (*ResponseResult, error) {
	// Prepare an HTTP request with the provided context.
	request, err := http.NewRequestWithContext(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", client.UserAgent)
	request.Header.Set("X-Auth-Token", client.Token)
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request = request.WithContext(ctx)

	// Send the HTTP request and populate the ResponseResult.
	response, err := client.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseResult := &ResponseResult{
		Response:    response,
		ErrNotFound: nil,
		ErrGeneric:  nil,
		Err:         nil,
	}

	// Check status code and populate custom error body with extended error message if it's possible.
	if response.StatusCode >= http.StatusBadRequest {
		err = responseResult.extractErr()
	}
	if err != nil {
		return nil, err
	}

	return responseResult, nil
}

// ResponseResult represents a result of an HTTP request.
// It embeds standard http.Response and adds custom API error representations.
type ResponseResult struct {
	*http.Response

	*ErrNotFound

	*ErrGeneric

	// Err contains an error that can be provided to a caller.
	Err error
}

// ErrNotFound represents 'not found' error of an HTTP response.
type ErrNotFound struct {
	Error struct {
		// Object ID.
		ID string `json:"id"`

		// Message of the error.
		Message string `json:"message"`
	} `json:"error"`
}

// ErrGeneric represents a generic error of an HTTP response.
type ErrGeneric struct {
	Error string `json:"error"`
}

// ExtractResult allows to provide an object into which ResponseResult body will be extracted.
func (result *ResponseResult) ExtractResult(to interface{}) error {
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return err
	}
	defer result.Body.Close()

	return json.Unmarshal(body, to)
}

// ExtractRaw extracts ResponseResult body into the slice of bytes without unmarshalling.
func (result *ResponseResult) ExtractRaw() ([]byte, error) {
	bytes, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	return bytes, nil
}

// extractErr populates an error message and error structure in the ResponseResult body.
func (result *ResponseResult) extractErr() error {
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return err
	}
	defer result.Body.Close()

	if len(body) == 0 {
		result.Err = fmt.Errorf(errGotHTTPStatusCodeFmt, result.StatusCode)

		return nil
	}
	if result.StatusCode == http.StatusNotFound {
		_ = json.Unmarshal(body, &result.ErrNotFound)
	} else {
		_ = json.Unmarshal(body, &result.ErrGeneric)
	}
	if result.ErrNotFound == nil && result.ErrGeneric == nil {
		result.Err = fmt.Errorf(errGotHTTPStatusCodeFmt, result.StatusCode)

		return nil
	}

	result.Err = fmt.Errorf(errGotHTTPStatusCodeFmt+": %s", result.StatusCode, string(body))

	return nil
}
