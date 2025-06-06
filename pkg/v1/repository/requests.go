package repository

import (
	"context"
	"errors"
	"net/http"
	"strings"

	v1 "github.com/selectel/craas-go/pkg"
	"github.com/selectel/craas-go/pkg/svc"
	"github.com/selectel/craas-go/pkg/v1/client"
	"github.com/selectel/craas-go/pkg/v1/registry"
)

var (
	ErrRepositoryNameEmpty = errors.New("repository name is empty")
	ErrImageNameEmpty      = errors.New("image name is empty")
)

// ListRepositories returns a list of all repositories for the specified registry.
func ListRepositories(ctx context.Context, client *client.ServiceClient, registryID string) ([]*Repository, *svc.ResponseResult, error) {
	if registryID == "" {
		return nil, nil, registry.ErrRegistryIDEmpty
	}

	url := strings.Join([]string{
		client.Endpoint(), v1.ResourceURLRegistries, registryID, v1.ResourceURLRepositories,
	}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract repositories from the response body.
	repositories := make([]*Repository, 0)
	err = responseResult.ExtractResult(&repositories)
	if err != nil {
		return nil, responseResult, err
	}

	return repositories, responseResult, nil
}

// GetRepository returns a single repository by its name.
func GetRepository(ctx context.Context, client *client.ServiceClient, registryID, repositoryName string) (*Repository, *svc.ResponseResult, error) {
	if registryID == "" {
		return nil, nil, registry.ErrRegistryIDEmpty
	}
	if repositoryName == "" {
		return nil, nil, ErrRepositoryNameEmpty
	}

	url := strings.Join([]string{
		client.Endpoint(), v1.ResourceURLRegistries, registryID, v1.ResourceURLRepositories, repositoryName,
	}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a repository from the response body.
	var repository Repository
	err = responseResult.ExtractResult(&repository)
	if err != nil {
		return nil, responseResult, err
	}

	return &repository, responseResult, nil
}

// DeleteRepository deletes a repository by its name.
func DeleteRepository(ctx context.Context, client *client.ServiceClient, registryID, repositoryName string) (*svc.ResponseResult, error) {
	if registryID == "" {
		return nil, registry.ErrRegistryIDEmpty
	}
	if repositoryName == "" {
		return nil, ErrRepositoryNameEmpty
	}

	url := strings.Join([]string{
		client.Endpoint(), v1.ResourceURLRegistries, registryID, v1.ResourceURLRepositories, repositoryName,
	}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		return responseResult, responseResult.Err
	}

	return responseResult, nil
}

// ListImages returns a list of all images for the specified repository.
func ListImages(ctx context.Context, client *client.ServiceClient, registryID, repositoryName string) ([]*Image, *svc.ResponseResult, error) {
	if registryID == "" {
		return nil, nil, registry.ErrRegistryIDEmpty
	}
	if repositoryName == "" {
		return nil, nil, ErrRepositoryNameEmpty
	}

	url := strings.Join([]string{
		client.Endpoint(), v1.ResourceURLRegistries, registryID, v1.ResourceURLRepositories, repositoryName, v1.ResourceURLImages,
	}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract images from the response body.
	images := make([]*Image, 0)
	err = responseResult.ExtractResult(&images)
	if err != nil {
		return nil, responseResult, err
	}

	return images, responseResult, nil
}

// ListTags returns a list of all tags for the specified repository.
func ListTags(ctx context.Context, client *client.ServiceClient, registryID, repositoryName string) ([]string, *svc.ResponseResult, error) {
	if registryID == "" {
		return nil, nil, registry.ErrRegistryIDEmpty
	}
	if repositoryName == "" {
		return nil, nil, ErrRepositoryNameEmpty
	}

	url := strings.Join([]string{
		client.Endpoint(), v1.ResourceURLRegistries, registryID, v1.ResourceURLRepositories, repositoryName, v1.ResourceURLTags,
	}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract tags from the response body.
	tags := make([]string, 0)
	err = responseResult.ExtractResult(&tags)
	if err != nil {
		return nil, responseResult, err
	}

	return tags, responseResult, nil
}

// ListImageLayers returns a list of all layers for the specified image.
// `image` could be represented as a tag or a digest.
func ListImageLayers(
	ctx context.Context,
	client *client.ServiceClient,
	registryID, repository, image string,
) ([]*Layer, *svc.ResponseResult, error) {
	if registryID == "" {
		return nil, nil, registry.ErrRegistryIDEmpty
	}
	if repository == "" {
		return nil, nil, ErrRepositoryNameEmpty
	}
	if image == "" {
		return nil, nil, ErrImageNameEmpty
	}

	url := strings.Join([]string{
		client.Endpoint(), v1.ResourceURLRegistries, registryID, v1.ResourceURLRepositories, repository, image,
	}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract layers from the response body.
	layers := make([]*Layer, 0)
	err = responseResult.ExtractResult(&layers)
	if err != nil {
		return nil, responseResult, err
	}

	return layers, responseResult, nil
}

// DeleteImageManifest deletes an image manifest by its name.
func DeleteImageManifest(ctx context.Context, client *client.ServiceClient, registryID, repository, image string) (*svc.ResponseResult, error) {
	if registryID == "" {
		return nil, registry.ErrRegistryIDEmpty
	}
	if repository == "" {
		return nil, ErrRepositoryNameEmpty
	}
	if image == "" {
		return nil, ErrImageNameEmpty
	}

	url := strings.Join([]string{
		client.Endpoint(), v1.ResourceURLRegistries, registryID, v1.ResourceURLRepositories, repository, image,
	}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		return responseResult, responseResult.Err
	}

	return responseResult, nil
}
