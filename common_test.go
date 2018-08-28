package triveutil

import (
	"errors"
	"net/http"
	"testing"
)

func TestStatusResponse(t *testing.T) {
	response := StatusResponse(http.StatusInternalServerError)
	expected := "{\"status\":500,\"message\":\"Internal Server Error\"}"
	if response != expected {
		t.Errorf("Status response was incorrect, got %s, want: %s", response, expected)
	}
}

func TestClientError(t *testing.T) {
	response, err := ClientError(http.StatusBadRequest)
	if err != nil {
		t.Errorf("ClientError produced an error")
	}
	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("Status code was incorrect, got %d, want: %d", response.StatusCode, http.StatusBadRequest)
	}
	expected := "{\"status\":400,\"message\":\"Bad Request\"}"
	if response.Body != expected {
		t.Errorf("Response body was incorrect, got %s, want: %s", response.Body, expected)
	}
}

func TestServerError(t *testing.T) {
	response, err := ServerError(errors.New("test server error"))
	if err != nil {
		t.Errorf("ServerError produced an error")
	}
	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Status code was incorrect, got %d, want: %d", response.StatusCode, http.StatusInternalServerError)
	}
	expected := "{\"status\":500,\"message\":\"Internal Server Error\"}"
	if response.Body != expected {
		t.Errorf("Response body was incorrect, got %s, want: %s", response.Body, expected)
	}
}
