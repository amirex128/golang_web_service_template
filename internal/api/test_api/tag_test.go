package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var tagID *string

func TestCreateTag(t *testing.T) {
	recorder := callApi([]byte(`
{
  "name": "عنوان تگ",
  "slug": "slug-tag"
}
	`),
		"/api/v1/user/tag/create",
		"POST")
	tagID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}
func TestAddTag(t *testing.T) {
	recorder := callApi([]byte(`
{
  "name": "عنوان تگ",
  "slug": "slug-tag"
}
	`),
		"/api/v1/user/tag/add",
		"POST")
	tagID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestShowTag(t *testing.T) {
	assert.NotNilf(t, tagID, "tag id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/tag/show/"+*tagID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexTag(t *testing.T) {
	recorder := callApi([]byte(``),
		"/api/v1/user/tag/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteTag(t *testing.T) {
	assert.NotNilf(t, tagID, "tag id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/tag/delete/"+*tagID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
