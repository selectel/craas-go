package client

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/selectel/craas-go/pkg/svc"
)

type ServiceClient struct {
	requests *svc.Request
}

// NewCRaaSClientV1 initializes a new CRaaS client for the V1 API.
func NewCRaaSClientV1(token, endpoint string) *ServiceClient {
	if !strings.HasSuffix(endpoint, "v1") {
		log.Fatalf("can't use client V2 with V1 endpoint")
	}

	return &ServiceClient{
		requests: &svc.Request{
			HTTPClient: svc.NewHTTPClient(),
			Token:      token,
			Endpoint:   endpoint,
			UserAgent:  svc.UserAgent,
		},
	}
}

func (s *ServiceClient) DoRequest(ctx context.Context, method string, path string, body io.Reader) (*svc.ResponseResult, error) {
	return s.requests.DoRequest(ctx, method, path, body)
}

func (s *ServiceClient) Token() string {
	return s.requests.Token
}

func (s *ServiceClient) Endpoint() string {
	return s.requests.Endpoint
}

func (s *ServiceClient) UserAgent() string {
	return s.requests.UserAgent
}

// NewCRaaSClientV2WithCustomHTTP initializes a new CRaaS client for the V1 API using custom HTTP client.
// If custom HTTP client is nil - default HTTP client will be used.
//
// vDeprecated: Use just v1.NewCRaaSClientV1 client constructors instead.
func NewCRaaSClientV1WithCustomHTTP(customHTTPClient *http.Client, token, endpoint string) (*ServiceClient, error) {
	if customHTTPClient == nil {
		customHTTPClient = newHTTPClient()
	}

	if strings.HasSuffix(endpoint, "v2") {
		return nil, svc.ErrEndpointVersionMismatch
	}

	return &ServiceClient{
		requests: &svc.Request{
			HTTPClient: customHTTPClient,
			Token:      token,
			Endpoint:   endpoint,
			UserAgent:  svc.UserAgent,
		},
	}, nil
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
