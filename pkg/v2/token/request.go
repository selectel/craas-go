package tokenv2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/selectel/craas-go/pkg/svc"
	v2 "github.com/selectel/craas-go/pkg/v2"
	clientv2 "github.com/selectel/craas-go/pkg/v2/client"
)

// Create method token.
func Create(ctx context.Context, client *clientv2.ServiceClient, tkn *TokenV2, dockerCfg *bool) (*TokenV2, *svc.ResponseResult, error) {
	val := url.Values{}
	url := strings.Join([]string{client.GetEndpoint(), v2.ResourceURLToken}, "/")
	if dockerCfg != nil {
		val.Add("docker-config", strconv.FormatBool(*dockerCfg))
		url = fmt.Sprintf("%s?%s", url, val.Encode())
	}
	reqBody, err := json.Marshal(tkn)
	if err != nil {
		return nil, nil, err
	}
	responseResult, err := client.DoRequest(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if responseResult.Err != nil {
		return nil, responseResult, err
	}

	// Extract token from the response body.
	var token TokenV2
	err = responseResult.ExtractResult(&token)
	if err != nil {
		return nil, responseResult, err
	}

	return &token, responseResult, nil
}

// List returns a list tokens.
func List(ctx context.Context, client *svc.ServiceClient, opts Opts) (*TokensV2, *svc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v2.ResourceURLToken}, "/")
	urlWithQuery := fmt.Sprintf("%s?%s", url, makeQueryString(opts))
	responseResult, err := client.DoRequest(ctx, http.MethodGet, urlWithQuery, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a tokens from the response body.
	var token TokensV2
	err = responseResult.ExtractResult(&token)
	if err != nil {
		return nil, responseResult, err
	}

	return &token, responseResult, nil
}

// Get returns a token by ID.
func GetByID(ctx context.Context, client *svc.ServiceClient, tokenID string) (*TokenV2, *svc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v2.ResourceURLToken, tokenID}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract a token from the response body.
	var token TokenV2
	err = responseResult.ExtractResult(&token)
	if err != nil {
		return nil, responseResult, err
	}
	token.Token = tokenID

	return &token, responseResult, nil
}

// Revoke revokes a token by its ID.
func Revoke(ctx context.Context, client *svc.ServiceClient, tokenID string) (*svc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v2.ResourceURLToken, tokenID, v2.ResourceURLRevoke}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		return responseResult, responseResult.Err
	}

	return responseResult, nil
}

// Refresh refresh a token by its ID.
func Refresh(ctx context.Context, client *svc.ServiceClient, tokenID string, exp Exp) (*TokenV2, *svc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v2.ResourceURLToken, tokenID, v2.ResourceURLRefresh}, "/")
	reqBody, err := json.Marshal(exp)
	if err != nil {
		return nil, nil, err
	}
	responseResult, err := client.DoRequest(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	var token TokenV2
	err = responseResult.ExtractResult(&token)
	if err != nil {
		return nil, responseResult, err
	}

	return &token, responseResult, nil
}

// Regenerate regenerate a token by its ID.
func Regenerate(ctx context.Context, client *svc.ServiceClient, tokenID string, exp Exp) (*TokenV2, *svc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v2.ResourceURLToken, tokenID, v2.ResourceURLRegenerate}, "/")
	reqBody, err := json.Marshal(exp)
	if err != nil {
		return nil, nil, err
	}
	responseResult, err := client.DoRequest(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	var token TokenV2
	err = responseResult.ExtractResult(&token)
	if err != nil {
		return nil, responseResult, err
	}

	return &token, responseResult, nil
}

// Delete delete a token by its ID.
func Delete(ctx context.Context, client *svc.ServiceClient, tokenID string) (*svc.ResponseResult, error) {
	url := strings.Join([]string{client.Endpoint, v2.ResourceURLToken, tokenID}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	if responseResult.Err != nil {
		return responseResult, responseResult.Err
	}

	return responseResult, nil
}

// Patch patch a token by its ID.
func Patch(ctx context.Context, client *svc.ServiceClient, tokenID string, name string, sc Scope) (*TokenV2, *svc.ResponseResult, error) {
	var token TokenV2
	if name != "" {
		token.Name = name
	}
	token.Scope = sc
	reqBody, err := json.Marshal(token)
	if err != nil {
		return nil, nil, err
	}
	url := strings.Join([]string{client.Endpoint, v2.ResourceURLToken, tokenID}, "/")
	responseResult, err := client.DoRequest(ctx, http.MethodPatch, url, bytes.NewReader(reqBody))
	if err != nil {
		return nil, nil, err
	}
	if responseResult.Err != nil {
		return nil, responseResult, responseResult.Err
	}

	// Extract token from the response body.
	var tokenResult TokenV2
	err = responseResult.ExtractResult(&tokenResult)
	if err != nil {
		return nil, responseResult, err
	}

	return &tokenResult, responseResult, nil
}
