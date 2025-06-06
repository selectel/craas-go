package testing

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/selectel/craas-go/pkg/testutils"
	"github.com/selectel/craas-go/pkg/v1/client"
	"github.com/selectel/craas-go/pkg/v1/repository"
)

func TestListRepositories(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v1/registries/" + testRegistryID + "/repositories",
		RawResponse: testListRepositoriesResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	actual, httpResponse, err := repository.ListRepositories(ctx, testClient, testRegistryID)
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
	if !reflect.DeepEqual(expectedListRepositoriesResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedListRepositoriesResponse, actual)
	}
}

func TestGetRepository(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v1/registries/" + testRegistryID + "/repositories/" + testRepositoryName,
		RawResponse: testGetRepositoryResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	actual, httpResponse, err := repository.GetRepository(ctx, testClient, testRegistryID, testRepositoryName)
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
	if !reflect.DeepEqual(expectedGetRepositoryResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedGetRepositoryResponse, actual)
	}
}

func TestDeleteRepository(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/api/v1/registries/" + testRegistryID + "/repositories/" + testRepositoryName,
		Method:   http.MethodDelete,
		Status:   http.StatusNoContent,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	httpResponse, err := repository.DeleteRepository(ctx, testClient, testRegistryID, testRepositoryName)
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
			http.StatusOK, httpResponse.StatusCode)
	}
}

func TestListImages(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v1/registries/" + testRegistryID + "/repositories/" + testRepositoryName + "/images",
		RawResponse: testListImagesResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	actual, httpResponse, err := repository.ListImages(ctx, testClient, testRegistryID, testRepositoryName)
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
	if !reflect.DeepEqual(expectedListImagesResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedListImagesResponse, actual)
	}
}

func TestListTags(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v1/registries/" + testRegistryID + "/repositories/" + testRepositoryName + "/tags",
		RawResponse: testListTagsResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	actual, httpResponse, err := repository.ListTags(ctx, testClient, testRegistryID, testRepositoryName)
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
	if !reflect.DeepEqual(expectedListTagsResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedListTagsResponse, actual)
	}
}

func TestListImageLayers(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         "/api/v1/registries/" + testRegistryID + "/repositories/" + testRepositoryName + "/" + testImageDigest,
		RawResponse: testListImageLayersResponseRaw,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	actual, httpResponse, err := repository.ListImageLayers(ctx, testClient, testRegistryID, testRepositoryName, testImageDigest)
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
	if !reflect.DeepEqual(expectedListImageLayersResponse, actual) {
		t.Fatalf("expected %#v, but got %#v", expectedListImageLayersResponse, actual)
	}
}

func TestDeleteImageManifest(t *testing.T) {
	endpointCalled := false
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	testutils.HandleReqWithoutBody(t, &testutils.HandleReqOpts{
		Mux:      testEnv.Mux,
		URL:      "/api/v1/registries/" + testRegistryID + "/repositories/" + testRepositoryName + "/" + testImageDigest,
		Method:   http.MethodDelete,
		Status:   http.StatusNoContent,
		CallFlag: &endpointCalled,
	})

	ctx := context.Background()
	testClient := client.NewCRaaSClientV1(testutils.TokenID, testEnv.Server.URL+"/api/v1")

	httpResponse, err := repository.DeleteImageManifest(ctx, testClient, testRegistryID, testRepositoryName, testImageDigest)
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
			http.StatusOK, httpResponse.StatusCode)
	}
}
