package gc

import (
	"context"
	"errors"
	"net/http"
	"strings"

	v1 "github.com/selectel/craas-go/pkg"
)

var ErrRegistryIDEmpty = errors.New("registry id is empty")

// StartGarbageCollection starts a garbage collection.
// Registry ID is a required parameter.
func StartGarbageCollection(ctx context.Context, client *v1.ServiceClient, registryID string, opts *StartGCOpts) (*v1.ResponseResult, error) {
	if registryID == "" {
		return nil, ErrRegistryIDEmpty
	}
	if opts == nil {
		opts = &StartGCOpts{}
	}

	url := strings.Join([]string{
		client.Endpoint, v1.ResourceURLRegistries, registryID, v1.ResourceURLGarbageCollection,
	}, "/")
	if opts.DeleteUntagged {
		url += "?delete-untagged=true"
	}
	responseResult, err := client.DoRequest(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		return responseResult, responseResult.Err
	}

	return responseResult, nil
}

func GetGarbageSize(ctx context.Context, client *v1.ServiceClient, registryID string) (*GarbageSize, *v1.ResponseResult, error) {
	if registryID == "" {
		return nil, nil, ErrRegistryIDEmpty
	}

	url := strings.Join([]string{
		client.Endpoint, v1.ResourceURLRegistries, registryID, v1.ResourceURLGarbageCollection, v1.ResourceURLSize,
	}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	var size GarbageSize
	err = responseResult.ExtractResult(&size)
	if err != nil {
		return nil, responseResult, err
	}

	return &size, responseResult, nil
}
