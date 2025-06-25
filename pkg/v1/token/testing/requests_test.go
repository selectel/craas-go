package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/craas-go/pkg/testutils"
	"github.com/selectel/craas-go/pkg/v1/client"
	"github.com/selectel/craas-go/pkg/v1/token"
)

func TestCreateToken(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v1/token",
		RawResponse: testCreateTokenResponseRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	actual, httpResponse, err := token.Create(ctx, testClient, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Post method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedCreateTokenResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedCreateTokenResponse, actual)
	}
}

func TestGetToken(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v1/token/" + testTokenID,
		RawResponse: testGetTokenResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	actual, httpResponse, err := token.Get(ctx, testClient, testTokenID)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Get method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedGetTokenResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetTokenResponse, actual)
	}
}

func TestRevokeToken(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/api/v1/token/" + testTokenID,
		Method:   http.MethodDelete,
		Status:   http.StatusNoContent,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	httpResponse, err := token.Revoke(ctx, testClient, testTokenID)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Delete method")
	}
	if httpResponse.StatusCode != http.StatusNoContent {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusNoContent, httpResponse.StatusCode)
	}
}

func TestRefreshToken(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v1/token/" + testTokenID + "/refresh",
		RawResponse: testRefreshTokenResponseRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	actual, httpResponse, err := token.Refresh(ctx, testClient, testTokenID)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Post method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedRefreshTokenResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedRefreshTokenResponse, actual)
	}
}
