package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var themeID *string

func TestCreateTheme(t *testing.T) {
	recorder := callApi([]byte(`
{
 "theme": "ادرس کامل",
 "city_id": 1,
 "full_name": "نام گیرنده",
 "lat": "35.5",
 "long": "36.5",
 "mobile": "09024809750",
 "postal_code": "1111111111",
 "province_id": 1,
 "title": "عنوان"
}
	`),
		"user/theme/create",
		"POST")
	themeID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdateTheme(t *testing.T) {
	assert.NotNilf(t, themeID, "theme id is nil")
	recorder := callApi([]byte(`
{
  "theme": "ادرس کامل",
  "city_id": 1,
  "full_name": "نام گیرنده",
  "id": 1,
  "lat": "35.123456",
  "long": "35.123456",
  "mobile": "09024809750",
  "postal_code": "1111111111",
  "province_id": 1,
  "title": "عنوان"
}
	`),
		"user/theme/update/"+*themeID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowTheme(t *testing.T) {
	assert.NotNilf(t, themeID, "theme id is nil")
	recorder := callApi([]byte(``),
		"user/theme/show/"+*themeID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexTheme(t *testing.T) {
	recorder := callApi([]byte(``),
		"user/theme/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteTheme(t *testing.T) {
	assert.NotNilf(t, themeID, "theme id is nil")
	recorder := callApi([]byte(``),
		"user/theme/delete/"+*themeID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
