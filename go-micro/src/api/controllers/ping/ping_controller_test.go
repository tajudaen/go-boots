package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tajud99n/go-micro/src/api/utils/test"
)

func TestConstant(t *testing.T) {
	assert.EqualValues(t, "ping", ping)
}

func TestPing(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	c := test.GetMockedContext(request, response)

	Ping(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "ping", response.Body.String())
}
