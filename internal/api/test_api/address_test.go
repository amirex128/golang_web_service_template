package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var addressID *string

func TestCreateAddress(t *testing.T) {

	recorder := callApi([]byte(`
{
 "address": "ادرس کامل",
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
		"/api/v1/user/address/create",
		"POST")
	addressID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdateAddress(t *testing.T) {

	assert.NotNilf(t, addressID, "address id is nil")
	recorder := callApi([]byte(`
{
  "address": "ادرس کامل",
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
		"/api/v1/user/address/update/"+*addressID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowAddress(t *testing.T) {

	assert.NotNilf(t, addressID, "address id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/address/show/"+*addressID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexAddress(t *testing.T) {

	recorder := callApi([]byte(``),
		"/api/v1/user/address/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteAddress(t *testing.T) {
	assert.NotNilf(t, addressID, "address id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/address/delete/"+*addressID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
