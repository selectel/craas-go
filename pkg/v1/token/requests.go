package token

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	v1 "github.com/selectel/craas-go/pkg"
	"github.com/selectel/craas-go/pkg/svc"
)

// Create request a creation of a new token.
// The token is valid for 12 hours by default and could be set with CreateOpts.
func Create(ctx context.Context, client *svc.ServiceClient, opts *CreateOpts) (*Token, *svc.ResponseResult, error) {
	if opts == nil {
		opts = &CreateOpts{
			TokenTTL: TTL12Hours,
		}
	}
	if err := validateTokenTTL(opts.TokenTTL); err != nil {
		return nil, nil, err
	}

	url := strings.Join([]string{client.Endpoint, v1.ResourceURLToken}, "/")
	urlWithQuery := fmt.Sprintf("%s?ttl=%s", url, opts.TokenTTL)
	responseResult, err := client.DoRequest(ctx, http.MethodPost, urlWithQuery, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract token from the response body.
	var token Token
	err = responseResult.ExtractResult(&token)
	if err != nil {
		return nil, responseResult, err
	}

	return &token, responseResult, nil
}

// validateTokenTTL validates token TTL parameter.
func validateTokenTTL(ttl TTL) error {
	switch ttl {
	case TTL12Hours, TTL1Year:
		return nil
	}

	return fmt.Errorf("%w: %s", ErrInvalidTokenTTL, ttl)
}

// Get returns a single token by its ID.
func Get(ctx context.Context, client *svc.ServiceClient, tokenID string) (*Token, *svc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLToken, tokenID}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a token from the response body.
	var token Token
	err = responseResult.ExtractResult(&token)
	if err != nil {
		return nil, responseResult, err
	}
	token.Token = tokenID

	return &token, responseResult, nil
}

// Revoke revokes a token by its ID.
func Revoke(ctx context.Context, client *svc.ServiceClient, tokenID string) (*svc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLToken, tokenID}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		return responseResult, responseResult.Err
	}

	return responseResult, nil
}

// Refresh refreshes a token by its ID.
func Refresh(ctx context.Context, client *svc.ServiceClient, tokenID string) (*Token, *svc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v1.ResourceURLToken, tokenID, v1.ResourceURLRefresh}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a token from the response body.
	var token Token
	err = responseResult.ExtractResult(&token)
	if err != nil {
		return nil, responseResult, err
	}
	token.Token = tokenID

	return &token, responseResult, nil
}
