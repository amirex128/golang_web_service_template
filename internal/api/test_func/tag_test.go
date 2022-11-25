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
 "tag": "ادرس کامل",
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
		"user/tag/create",
		"POST")
	tagID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdateTag(t *testing.T) {
	assert.NotNilf(t, tagID, "tag id is nil")
	recorder := callApi([]byte(`
{
  "tag": "ادرس کامل",
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
		"user/tag/update/"+*tagID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowTag(t *testing.T) {
	assert.NotNilf(t, tagID, "tag id is nil")
	recorder := callApi([]byte(``),
		"user/tag/show/"+*tagID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexTag(t *testing.T) {
	recorder := callApi([]byte(``),
		"user/tag/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteTag(t *testing.T) {
	assert.NotNilf(t, tagID, "tag id is nil")
	recorder := callApi([]byte(``),
		"user/tag/delete/"+*tagID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
