package v1

import (
	"net"
	"net/http"
	"time"

	"github.com/selectel/craas-go/pkg/svc"
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

// NewCRaaSClientV1 initializes a new CRaaS client for the V1 API.
//
// Deprecated: Use v1 or v2 client constructors instead.
func NewCRaaSClientV1(token, endpoint string) *svc.Request {
	return &svc.Request{
		HTTPClient: newHTTPClient(),
		Token:      token,
		Endpoint:   endpoint,
		UserAgent:  userAgent,
	}
}

// NewCRaaSClientV1WithCustomHTTP initializes a new CRaaS client for the V1 API using custom HTTP client.
// If custom HTTP client is nil - default HTTP client will be used.
//
// Deprecated: Use v1 or v2 client constructors instead.
func NewCRaaSClientV1WithCustomHTTP(customHTTPClient *http.Client, tokenID, endpoint string) *svc.Request {
	if customHTTPClient == nil {
		customHTTPClient = newHTTPClient()
	}

	return &svc.Request{
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
