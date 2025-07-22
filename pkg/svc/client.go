package svc

import (
	"net"
	"net/http"
	"time"
)

const (
	// appName represents an application name.
	appName = "craas-go"

	// appVersion is a version of the application.
	appVersion = "0.2.0"

	// userAgent contains a basic user agent that will be used in queries.
	UserAgent = appName + "/" + appVersion

	// defaultHTTPTimeout represents the default timeout (in seconds) for HTTP requests.
	DefaultHTTPTimeout = 120

	// defaultDialTimeout represents the default timeout (in seconds) for HTTP connection establishments.
	DefaultDialTimeout = 60

	// defaultKeepaliveTimeout represents the default keep-alive period for an active network connection.
	DefaultKeepaliveTimeout = 60

	// defaultMaxIdleConns represents the maximum number of idle (keep-alive) connections.
	DefaultMaxIdleConns = 100

	// defaultIdleConnTimeout represents the maximum amount of time an idle (keep-alive) connection will remain
	// idle before closing itself.
	DefaultIdleConnTimeout = 100

	// defaultTLSHandshakeTimeout represents the default timeout (in seconds) for TLS handshake.
	DefaultTLSHandshakeTimeout = 60

	// defaultExpectContinueTimeout represents the default amount of time to wait for a server's first
	// response headers.
	DefaultExpectContinueTimeout = 1
)

// newHTTPClient returns a reference to an initialized and configured HTTP client.
func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout:   DefaultHTTPTimeout * time.Second,
		Transport: NewHTTPTransport(),
	}
}

// newHTTPTransport returns a reference to an initialized and configured HTTP transport.
func NewHTTPTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   DefaultDialTimeout * time.Second,
			KeepAlive: DefaultKeepaliveTimeout * time.Second,
		}).DialContext,
		MaxIdleConns:          DefaultMaxIdleConns,
		IdleConnTimeout:       DefaultIdleConnTimeout * time.Second,
		TLSHandshakeTimeout:   DefaultTLSHandshakeTimeout * time.Second,
		ExpectContinueTimeout: DefaultExpectContinueTimeout * time.Second,
	}
}
