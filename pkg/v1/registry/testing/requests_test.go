package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/craas-go/pkg/testutils"
	"github.com/selectel/craas-go/pkg/v1/client"
	"github.com/selectel/craas-go/pkg/v1/registry"
)

func TestCreate(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v1/registries",
		RawResponse: testCreateRegistryResponseRaw,
		RawRequest:  testCreateRegistryRequestRaw,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	actual, httpResponse, err := registry.Create(ctx, testClient, "test-registry")
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Post method")
	}
	if httpResponse.StatusCode != http.StatusCreated {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedCreateRegistryResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedCreateRegistryResponse, actual)
	}
}

func TestGet(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v1/registries/" + testRegistryID,
		RawResponse: testGetRegistryResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	actual, httpResponse, err := registry.Get(ctx, testClient, testRegistryID)
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
	if !reflect.DeepEqual(expectedGetRegistryResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetRegistryResponse, actual)
	}
}

func TestGetWithUnknownStatus(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v1/registries/" + testRegistryID,
		RawResponse: testGetRegistryWithUnknownStatusResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	actual, httpResponse, err := registry.Get(ctx, testClient, testRegistryID)
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
	if !reflect.DeepEqual(expectedGetRegistryWithUnknownStatusResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetRegistryWithUnknownStatusResponse, actual)
	}
}

func TestList(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v1/registries",
		RawResponse: testListRegistriesResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	actual, httpResponse, err := registry.List(ctx, testClient)
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
	if !reflect.DeepEqual(expectedListRegistriesResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedListRegistriesResponse, actual)
	}
}

func TestDelete(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/api/v1/registries/" + testRegistryID,
		Method:   http.MethodDelete,
		Status:   http.StatusNoContent,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	httpResponse, err := registry.Delete(ctx, testClient, testRegistryID)
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
