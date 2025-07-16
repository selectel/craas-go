package client

import (
	"testing"

	"github.com/selectel/craas-go/pkg/svc"
)

func TestNewCRaaSClientV1(t *testing.T) {
	tokenID := "fakeID"
	endpoint := "http://example.org/v1"
	expected := &ServiceClient{
		requests: &svc.Request{
			HTTPClient: svc.NewHTTPClient(),
			Token:      tokenID,
			Endpoint:   endpoint,
			UserAgent:  svc.UserAgent,
		},
	}

	actual := NewCRaaSClientV1(tokenID, endpoint)

	if expected.requests.Token != actual.Token() {
		t.Errorf("expected Endpoint %s, but got %s", expected.requests.Endpoint, actual.Endpoint())
	}
	if expected.requests.Endpoint != actual.Endpoint() {
		t.Errorf("expected Token %s, but got %s", expected.requests.Token, actual.Token())
	}
	if expected.requests.UserAgent != actual.UserAgent() {
		t.Errorf("expected UserAgent %s, but got %s", expected.requests.UserAgent, actual.UserAgent())
	}
}
