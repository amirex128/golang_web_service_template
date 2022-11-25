package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var pageID *string

func TestCreatePage(t *testing.T) {
	recorder := callApi([]byte(`
{
  "body": "<p>متن صفحه درباره ما</p>",
  "shop_id": 1,
  "slug": "about-us",
  "title": "درباره ما",
  "type": "normal"
}
	`),
		"user/page/create",
		"POST")
	pageID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdatePage(t *testing.T) {
	assert.NotNilf(t, pageID, "page id is nil")
	recorder := callApi([]byte(`
{
  "body": "<p>متن صفحه تماس با ما</p>",
  "id": 1,
  "slug": "contact-us",
  "title": "تماس با ما",
  "type": "normal"
}
	`),
		"user/page/update/"+*pageID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowPage(t *testing.T) {
	assert.NotNilf(t, pageID, "page id is nil")
	recorder := callApi([]byte(``),
		"user/page/show/"+*pageID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexPage(t *testing.T) {
	recorder := callApi([]byte(``),
		"user/page/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeletePage(t *testing.T) {
	assert.NotNilf(t, pageID, "page id is nil")
	recorder := callApi([]byte(``),
		"user/page/delete/"+*pageID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
