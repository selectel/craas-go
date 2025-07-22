package svc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/selectel/craas-go/pkg/testutils"
)

const (
	userAgent = "agent"
	token     = "token"
)

func newFakeClient(endpoint string) *Request {
	return &Request{
		Token:      token,
		Endpoint:   endpoint,
		UserAgent:  userAgent,
		HTTPClient: &http.Client{},
	}
}

func TestDoGetRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, "response")

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := newFakeClient(endpoint)

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("got %d response status, want 200", response.StatusCode)
	}
}

func TestDoPostRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, "response")

		if r.Method != http.MethodPost {
			t.Errorf("got %s method, want POST", r.Method)
		}

		_, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unable to read the request body: %v", err)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := newFakeClient(endpoint)

	requestBody, err := json.Marshal(&struct {
		ID string `json:"id"`
	}{
		ID: "uuid",
	})
	if err != nil {
		t.Fatalf("can't marshal JSON: %v", err)
	}

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodPost, endpoint, bytes.NewReader(requestBody))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusOK {
		t.Fatalf("got %d response status, want 200", response.StatusCode)
	}
}

func TestDoErrNotFoundRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"error":{"id":"9fb12d6e-0da2-4db1-a076-414059cfb448","message":"Registry not found"}}`)

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := newFakeClient(endpoint)

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("got %d response status, want 404", response.StatusCode)
	}

	if response.ErrNotFound.Error.Message != "Registry not found" {
		t.Fatalf("got %s error message, want 'Registry not found'", response.ErrNotFound.Error.Message)
	}

	if response.ErrNotFound.Error.ID != "9fb12d6e-0da2-4db1-a076-414059cfb448" {
		t.Fatalf("got %s object id, want '9fb12d6e-0da2-4db1-a076-414059cfb448'", response.ErrNotFound.Error.ID)
	}
}

func TestDoErrGenericRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `{"error":"id value is invalid"}`)

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := newFakeClient(endpoint)

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("got %d response status, want 400", response.StatusCode)
	}

	if response.ErrGeneric.Error != "id value is invalid" {
		t.Fatalf("got %s error message, want 'id value is invalid'", response.ErrGeneric.Error)
	}
}

func TestDoErrNoContentRequest(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, "") // write no content in the response body.

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := newFakeClient(endpoint)

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusBadGateway {
		t.Fatalf("got %d response status, want 502", response.StatusCode)
	}

	if response.Err.Error() != "craas-go: got the 502 status code from the server" {
		t.Fatalf("got %s error message, want 'craas-go: got the 502 status code from the server'", response.Err.Error())
	}
}

func TestDoErrRequestUnmarshalError(t *testing.T) {
	testEnv := testutils.SetupTestEnv()
	defer testEnv.TearDownTestEnv()
	testEnv.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{") // write invalid json in the response body.

		if r.Method != http.MethodGet {
			t.Errorf("got %s method, want GET", r.Method)
		}
	})

	endpoint := testEnv.Server.URL + "/"
	client := newFakeClient(endpoint)

	ctx := context.Background()
	response, err := client.DoRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if response.Body == nil {
		t.Fatal("response body is empty")
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("got %d response status, want 500", response.StatusCode)
	}

	if response.Err.Error() != "craas-go: got the 500 status code from the server" {
		t.Fatalf("got %s error message, want 'craas-go: got the 500 status code from the server'", response.Err.Error())
	}
}
