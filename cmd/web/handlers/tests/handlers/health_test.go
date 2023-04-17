package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pedidosya/@project_name@/cmd/web/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHealthCheck(t *testing.T) {
	r := getRouter()
	e := getEngine()
	r.GET("/health", specHandler(e, handlers.HandleHealthCheck))
	expected := gin.H{
		"env": "",
		"name": "@project_name@",
		"version": "",
	}
	w := performRequest(r, "GET", "/health")

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["message"]

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expected["message"], value)
}
