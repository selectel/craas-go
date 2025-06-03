package registry

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	v1 "github.com/selectel/craas-go/pkg"
	"github.com/selectel/craas-go/pkg/svc"
)

var (
	ErrRegistryNameEmpty = errors.New("registry name is empty")
	ErrRegistryIDEmpty   = errors.New("registry id is empty")
)

// Create creates a new registry with the specified options.
// Registry name is a required parameter.
func Create(ctx context.Context, client *svc.ServiceClient, name string) (*Registry, *svc.ResponseResult, error) {
	if name == "" {
		return nil, nil, ErrRegistryNameEmpty
	}

	requestBody, err := json.Marshal(CreateOpts{Name: name})
	if err != nil {
		return nil, nil, err
	}

	url := strings.Join([]string{client.Endpoint, v1.ResourceURLRegistries}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodPost, url, bytes.NewReader(requestBody))
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a registry from the response body.
	var registry Registry
	err = responseResult.ExtractResult(&registry)
	if err != nil {
		return nil, responseResult, err
	}

	return &registry, responseResult, nil
}

// List returns a list of all registries.
func List(ctx context.Context, client *svc.ServiceClient) ([]*Registry, *svc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLRegistries}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract registries from the response body.
	registries := make([]*Registry, 0)

	err = responseResult.ExtractResult(&registries)
	if err != nil {
		return nil, responseResult, err
	}

	return registries, responseResult, nil
}

// Get returns a single registry by its id.
// Registry ID is a required parameter.
func Get(ctx context.Context, client *svc.ServiceClient, registryID string) (*Registry, *svc.ResponseResult, error) {
	if registryID == "" {
		return nil, nil, ErrRegistryIDEmpty
	}

	url := strings.Join([]string{client.Endpoint, v1.ResourceURLRegistries, registryID}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a registry from the response body.
	var registry Registry
	err = responseResult.ExtractResult(&registry)
	if err != nil {
		return nil, responseResult, err
	}

	return &registry, responseResult, nil
}

// Delete deletes a registry by its id.
// Registry ID is a required parameter.
func Delete(ctx context.Context, client *svc.ServiceClient, registryID string) (*svc.ResponseResult, error) {
	if registryID == "" {
		return nil, ErrRegistryIDEmpty
	}

	url := strings.Join([]string{client.Endpoint, v1.ResourceURLRegistries, registryID}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		return responseResult, responseResult.Err
	}

	return responseResult, nil
}
