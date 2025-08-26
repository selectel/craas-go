package testing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/craas-go/pkg/testutils"
	"github.com/selectel/craas-go/pkg/v2/client"
	tokenV2 "github.com/selectel/craas-go/pkg/v2/token"
)

const apiV2 = "/api/v2"

func TestCreateToken(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v2/tokens",
		RawResponse: testCreateTokenResponseRaw,
		Method:      http.MethodPost,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient, err := client.NewCRaaSClientV2(testutils.TokenID, testEnv.Server.URL+apiV2)
	if err != nil {
		t.Errorf("got error %s", err)
	}
	actual, response, err := tokenV2.Create(ctx, testClient, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatal("expected an HTTP response from the POST method")
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if !reflect.DeepEqual(expectedCreateTokenResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedCreateTokenResponse, actual)
	}
}

func TestGetListToken(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v2/tokens",
		RawResponse: testListTokensResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient, err := client.NewCRaaSClientV2(testutils.TokenID, testEnv.Server.URL+apiV2)
	if err != nil {
		t.Errorf("got error %s", err)
	}
	dig := new(int)
	*dig = 1
	opts := tokenV2.Opts{
		Limit: dig,
	}

	actual, response, err := tokenV2.List(ctx, testClient, opts)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if response == nil {
		t.Fatal("expected an HTTP response from the GET method")
	}
	if !reflect.DeepEqual(expectedListTokenResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedListTokenResponse, actual)
	}
}

func TestDeleteToken(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/api/v2/tokens/%s", testTokenID),
		RawResponse: testCreateTokenResponseRaw,
		Method:      http.MethodDelete,
		Status:      http.StatusNoContent,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient, err := client.NewCRaaSClientV2(testutils.TokenID, testEnv.Server.URL+apiV2)
	if err != nil {
		t.Errorf("got error %s", err)
	}
	actual, err := tokenV2.Delete(ctx, testClient, testTokenID)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if actual.StatusCode != http.StatusNoContent {
		t.Fatalf("expected %d status in the HTTP response, but got %d", http.StatusNoContent, actual.StatusCode)
	}
}

func TestPatchToken(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf("/api/v2/tokens/%s", testTokenID),
		RawResponse: testPatchTokenResponseRaw,
		Method:      http.MethodPatch,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient, err := client.NewCRaaSClientV2(testutils.TokenID, testEnv.Server.URL+apiV2)
	if err != nil {
		t.Errorf("got error %s", err)
	}
	scope := tokenV2.Scope{
		ModeRW:        true,
		RegistryIDs:   []string{"888af692-c646-4b76-a234-81ca9b5bcafe", "6303699d-c2cd-40b1-8428-9dcd6cc3d00d"},
		AllRegistries: false,
	}
	exp := tokenV2.Exp{
		ExpiresAt: expiresAt,
	}
	actual, response, err := tokenV2.Patch(ctx, testClient, testTokenID, "token", scope, exp)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d", http.StatusOK, response.StatusCode)
	}
	if !reflect.DeepEqual(expectedPatchTokenResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedPatchTokenResponse, actual)
	}
}
