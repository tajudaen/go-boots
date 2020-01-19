package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMockContext(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "localhost:123/something", nil)
	assert.Nil(t, err)
	request.Header = http.Header{"X-Mock": {"true"}}
	response := httptest.NewRecorder()
	c := GetMockedContext(request, response)

	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "123", c.Request.URL.Port())
	assert.EqualValues(t, "/something", c.Request.URL.Path)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
	assert.EqualValues(t, 1, len(c.Request.Header))
	assert.EqualValues(t, "true", c.GetHeader("X-Mock"))
}
