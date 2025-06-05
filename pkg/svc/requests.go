package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ServiceClient stores details that are needed to work with Selectel CRaaS API.
type Requests struct {
	// HTTPClient represents an initialized HTTP client that will be used to do requests.
	HTTPClient *http.Client

	// Token is a client authentication token.
	Token string

	// Endpoint represents an endpoint that will be used in all requests.
	Endpoint string

	// UserAgent contains user agent that will be used in all requests.
	UserAgent string
}

const errGotHTTPStatusCodeFmt = "craas-go: got the %d status code from the server"

// DoRequest performs the HTTP request with the current ServiceClient's HTTPClient.
// Authentication and optional headers will be added automatically.
func (client *Requests) DoRequest(ctx context.Context, method, path string, body io.Reader) (*ResponseResult, error) {
	// Prepare an HTTP request with the provided context.
	request, err := http.NewRequestWithContext(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", client.UserAgent)
	request.Header.Set("X-Auth-Token", client.Token)
	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	request = request.WithContext(ctx)

	// Send the HTTP request and populate the ResponseResult.
	response, err := client.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	responseResult := &ResponseResult{
		Response:    response,
		ErrNotFound: nil,
		ErrGeneric:  nil,
		Err:         nil,
	}

	// Check status code and populate custom error body with extended error message if it's possible.
	if response.StatusCode >= http.StatusBadRequest {
		err = responseResult.extractErr()
	}
	if err != nil {
		return nil, err
	}

	return responseResult, nil
}

// ResponseResult represents a result of an HTTP request.
// It embeds standard http.Response and adds custom API error representations.
type ResponseResult struct {
	*http.Response

	*ErrNotFound

	*ErrGeneric

	// Err contains an error that can be provided to a caller.
	Err error
}

// ErrNotFound represents 'not found' error of an HTTP response.
type ErrNotFound struct {
	Error struct {
		// Object ID.
		ID string `json:"id"`

		// Message of the error.
		Message string `json:"message"`
	} `json:"error"`
}

// ErrGeneric represents a generic error of an HTTP response.
type ErrGeneric struct {
	Error string `json:"error"`
}

// ExtractResult allows to provide an object into which ResponseResult body will be extracted.
func (result *ResponseResult) ExtractResult(to interface{}) error {
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return err
	}
	defer result.Body.Close()

	return json.Unmarshal(body, to)
}

// ExtractRaw extracts ResponseResult body into the slice of bytes without unmarshalling.
func (result *ResponseResult) ExtractRaw() ([]byte, error) {
	bytes, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	return bytes, nil
}

// extractErr populates an error message and error structure in the ResponseResult body.
func (result *ResponseResult) extractErr() error {
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return err
	}
	defer result.Body.Close()

	if len(body) == 0 {
		result.Err = fmt.Errorf(errGotHTTPStatusCodeFmt, result.StatusCode)

		return nil
	}
	if result.StatusCode == http.StatusNotFound {
		_ = json.Unmarshal(body, &result.ErrNotFound)
	} else {
		_ = json.Unmarshal(body, &result.ErrGeneric)
	}
	if result.ErrNotFound == nil && result.ErrGeneric == nil {
		result.Err = fmt.Errorf(errGotHTTPStatusCodeFmt, result.StatusCode)

		return nil
	}

	result.Err = fmt.Errorf(errGotHTTPStatusCodeFmt+": %s", result.StatusCode, string(body))

	return nil
}
