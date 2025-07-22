package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/craas-go/pkg/testutils"
	"github.com/selectel/craas-go/pkg/v1/client"
	"github.com/selectel/craas-go/pkg/v1/gc"
)

func TestStartGarbageCollection(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/api/v1/registries/" + testRegistryID + "/garbage-collection",
		Method:   http.MethodPost,
		Status:   http.StatusCreated,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	testClient, err := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")
	if err != nil {
		t.Fatal(err)
	}
	httpResponse, err := gc.StartGarbageCollection(ctx, testClient, testRegistryID, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Delete method")
	}
	if httpResponse.StatusCode != http.StatusCreated {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusCreated, httpResponse.StatusCode)
	}
}

func TestGetGarbageSize(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v1/registries/" + testRegistryID + "/garbage-collection/size",
		RawResponse: testGetGarbageSizeResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient, err := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")
	if err != nil {
		t.Fatal(err)
	}
	actual, httpResponse, err := gc.GetGarbageSize(ctx, testClient, testRegistryID)
	if err != nil {
		t.Fatal(err)
	}
	if !endpointCalled {
		t.Fatal("endpoint wasn't called")
	}
	if httpResponse == nil {
		t.Fatal("expected an HTTP response from the Delete method")
	}
	if httpResponse.StatusCode != http.StatusOK {
		t.Fatalf("expected %d status in the HTTP response, but got %d",
			http.StatusOK, httpResponse.StatusCode)
	}
	if !reflect.DeepEqual(expectedGetGarbageSizeResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetGarbageSizeResponse, actual)
	}
}
