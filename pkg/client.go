package v1

import (
	"net"
	"net/http"
	"time"

	"github.com/selectel/craas-go/pkg/svc"
	clientv1 "github.com/selectel/craas-go/pkg/v1/client"
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

// NewCRaaSClientV1 initializes a new CRaaS client for the V1 API.
//
// Deprecated: Use v1 or v2 client constructors instead.
func NewCRaaSClientV1(token, endpoint string) *clientv1.ServiceClient {
	client, err := clientv1.NewCRaaSClientV1(token, endpoint)
	if err != nil {
		panic(err)
	}

	return client
}

// NewCRaaSClientV1WithCustomHTTP initializes a new CRaaS client for the V1 API using custom HTTP client.
// If custom HTTP client is nil - default HTTP client will be used.
//
// Deprecated: Use v1 or v2 client constructors instead.
func NewCRaaSClientV1WithCustomHTTP(customHTTPClient *http.Client, tokenID, endpoint string) *clientv1.ServiceClient {
	if customHTTPClient == nil {
		customHTTPClient = newHTTPClient()
	}

	client, err := clientv1.NewCRaaSClientV1WithCustomHTTP(customHTTPClient, tokenID, endpoint)
	if err != nil {
		panic(err)
	}

	return client
}

// newHTTPClient returns a reference to an initialized and configured HTTP client.
func newHTTPClient() *http.Client {
	return &http.Client{
		Timeout:   svc.DefaultHTTPTimeout * time.Second,
		Transport: newHTTPTransport(),
	}
}

// newHTTPTransport returns a reference to an initialized and configured HTTP transport.
func newHTTPTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   svc.DefaultDialTimeout * time.Second,
			KeepAlive: svc.DefaultKeepaliveTimeout * time.Second,
		}).DialContext,
		MaxIdleConns:          svc.DefaultMaxIdleConns,
		IdleConnTimeout:       svc.DefaultIdleConnTimeout * time.Second,
		TLSHandshakeTimeout:   svc.DefaultTLSHandshakeTimeout * time.Second,
		ExpectContinueTimeout: svc.DefaultExpectContinueTimeout * time.Second,
	}
}
