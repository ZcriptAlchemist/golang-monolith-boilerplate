package core

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntializeRouter(t *testing.T) {
	router, _ := IntializeRouter()

	t.Run("Health Endpoint", func(t *testing.T) {
		// Create a test HTTP request to the /health endpoint
		req, _ := http.NewRequest(http.MethodGet, "/health", nil)

		// Use httptest to simulate the request
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert the response status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Assert the response body
		expected := `{"message":"Application up and running successfully","status":"OK"}`
		assert.JSONEq(t, expected, w.Body.String())
	})
}
