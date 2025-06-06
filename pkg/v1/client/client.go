package client

import (
	"context"
	"io"
	"log"
	"strings"

	"github.com/selectel/craas-go/pkg/svc"
)

type ServiceClient struct {
	requests *svc.Request
}

// NewCRaaSClientV1 initializes a new CRaaS client for the V1 API.
func NewCRaaSClientV1(token, endpoint string) *ServiceClient {
	if strings.HasSuffix(endpoint, "v2") {
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
