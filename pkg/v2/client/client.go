package client

import (
	"context"
	"io"
	"strings"

	"github.com/selectel/craas-go/pkg/svc"
)

const clientVersion = "v2"

// ServiceClient stores details that are needed to work with Selectel CRaaS API.
type ServiceClient struct {
	requests *svc.Requests
}

func NewCRaaSClientV2(token, endpoint string) (*ServiceClient, error) {
	if !strings.HasSuffix(endpoint, "v2") {
		return nil, svc.ErrEndpointVersionMismatch
	}

	requests := &svc.Requests{
		HTTPClient: svc.NewHTTPClient(),
		Token:      token,
		Endpoint:   endpoint,
		UserAgent:  svc.UserAgent,
	}

	return &ServiceClient{
		requests: requests,
	}, nil
}

func (sc *ServiceClient) DoRequest(ctx context.Context, method, path string, body io.Reader) (*svc.ResponseResult, error) {
	return sc.requests.DoRequest(ctx, method, path, body)
}

// GetToken returns the token from the service client's requests
func (sc *ServiceClient) GetToken() string {
	return sc.requests.Token
}

// GetEndpoint returns the endpoint from the service client's requests
func (sc *ServiceClient) GetEndpoint() string {
	return sc.requests.Endpoint
}
