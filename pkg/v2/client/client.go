package client

import (
	"context"
	"io"
	"strings"

	"github.com/selectel/craas-go/pkg/svc"
)

type ServiceClient struct {
	requests *svc.Request
}

// NewCRaaSClientV2 initializes a new CRaaS client for the V2 API.
func NewCRaaSClientV2(token, endpoint string) (*ServiceClient, error) {
	if strings.HasSuffix(endpoint, "v1") {
		return nil, svc.ErrEndpointVersionMismatch
	}

	return &ServiceClient{
		requests: &svc.Request{
			HTTPClient: svc.NewHTTPClient(),
			Token:      token,
			Endpoint:   endpoint,
			UserAgent:  svc.UserAgent,
		},
	}, nil
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
