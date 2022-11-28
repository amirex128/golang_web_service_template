package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var domainID *string

func TestCreateDomain(t *testing.T) {
	recorder := callApi([]byte(`
{
  "name": "example.com",
  "shop_id": 1,
  "type": "domain"
}
	`),
		"/api/v1/user/domain/create",
		"POST")
	domainID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestShowDomain(t *testing.T) {
	assert.NotNilf(t, domainID, "domain id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/domain/show/"+*domainID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexDomain(t *testing.T) {
	recorder := callApi([]byte(``),
		"/api/v1/user/domain/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteDomain(t *testing.T) {
	assert.NotNilf(t, domainID, "domain id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/domain/delete/"+*domainID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
