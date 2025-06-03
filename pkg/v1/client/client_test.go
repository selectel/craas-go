package client_test

import (
	"net/http"
	"testing"
	"time"

	v1 "github.com/selectel/craas-go/pkg"
	"github.com/selectel/craas-go/pkg/svc"
	"github.com/selectel/craas-go/pkg/testutils"
)

const userAgent = "craas-go/0.1.0"

func TestNewCRaaSClientV1(t *testing.T) {
	tokenID := "fakeID"
	endpoint := "http://example.org"
	expected := &svc.ServiceClient{
		Token:     tokenID,
		Endpoint:  endpoint,
		UserAgent: userAgent,
	}

	actual := v1.NewCRaaSClientV1(tokenID, endpoint)

	if expected.Token != actual.Token {
		t.Errorf("expected Endpoint %s, but got %s", expected.Endpoint, actual.Endpoint)
	}
	if expected.Endpoint != actual.Endpoint {
		t.Errorf("expected Token %s, but got %s", expected.Token, actual.Token)
	}
	if expected.UserAgent != actual.UserAgent {
		t.Errorf("expected UserAgent %s, but got %s", expected.UserAgent, actual.UserAgent)
	}
	if actual.HTTPClient == nil {
		t.Errorf("expected initialized HTTPClient but it's nil")
	}
}

func TestNewCRaaSClientV1WithCustomHTTP(t *testing.T) {
	tokenID := testutils.TokenID
	endpoint := "http://example.org"
	expected := &svc.ServiceClient{
		Token:     tokenID,
		Endpoint:  endpoint,
		UserAgent: userAgent,
	}
	customHTTPClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	actual := v1.NewCRaaSClientV1WithCustomHTTP(customHTTPClient, tokenID, endpoint)

	if expected.Token != actual.Token {
		t.Errorf("expected Endpoint %s, but got %s", expected.Endpoint, actual.Endpoint)
	}
	if expected.Endpoint != actual.Endpoint {
		t.Errorf("expected Token %s, but got %s", expected.Token, actual.Token)
	}
	if expected.UserAgent != actual.UserAgent {
		t.Errorf("expected UserAgent %s, but got %s", expected.UserAgent, actual.UserAgent)
	}
	if actual.HTTPClient == nil {
		t.Errorf("expected initialized HTTPClient but it's nil")
	}
}
