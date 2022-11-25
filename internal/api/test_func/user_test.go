package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var userID *string

func TestCreateUser(t *testing.T) {
	recorder := callApi([]byte(`
{
 "user": "ادرس کامل",
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
		"user/user/create",
		"POST")
	userID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdateUser(t *testing.T) {
	assert.NotNilf(t, userID, "user id is nil")
	recorder := callApi([]byte(`
{
  "user": "ادرس کامل",
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
		"user/user/update/"+*userID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowUser(t *testing.T) {
	assert.NotNilf(t, userID, "user id is nil")
	recorder := callApi([]byte(``),
		"user/user/show/"+*userID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexUser(t *testing.T) {
	recorder := callApi([]byte(``),
		"user/user/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteUser(t *testing.T) {
	assert.NotNilf(t, userID, "user id is nil")
	recorder := callApi([]byte(``),
		"user/user/delete/"+*userID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
