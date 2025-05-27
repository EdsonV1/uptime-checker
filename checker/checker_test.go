package checker

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckUrl_Up(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	result := CheckUrl(server.URL)
	assert.Equal(t, "UP", result)
}

func TestCheckUrl_Warning(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer server.Close()

	result := CheckUrl(server.URL)
	assert.True(t, strings.HasPrefix(result, "WARNING"))
	assert.Contains(t, result, "500")
}

func TestCheckUrl_Down(t *testing.T) {
	result := CheckUrl("http://localhost:9999")
	assert.True(t, strings.HasPrefix(result, "DOWN"))
}
