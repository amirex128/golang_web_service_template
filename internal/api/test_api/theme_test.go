package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestIndexTheme(t *testing.T) {
	recorder := callApi([]byte(``),
		"/api/v1/user/theme/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
