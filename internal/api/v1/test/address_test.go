package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
		"user/address/create",
		"POST")

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.NotContains(t, recorder.Body.String(), "error")
	parseErr(recorder)

}

func TestDeleteAddress(t *testing.T) {
	id := getID(callApi([]byte(`

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
		"user/address/create",
		"POST"))
	if !assert.NotNilf(t, id, "Create model failed for get id") {
		return
	}

	recorder := httptest.NewRecorder()
	callApi([]byte(``),
		"user/address/delete/"+*id,
		"POST")

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.NotContains(t, recorder.Body.String(), "error")
	parseErr(recorder)
}

func TestUpdateAddress(t *testing.T) {
	id := getID(callApi([]byte(`

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
		"user/address/create",
		"POST"))
	if !assert.NotNilf(t, id, "Create model failed for get id") {
		return
	}

	recorder := httptest.NewRecorder()
	callApi([]byte(`

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
		"user/address/update/"+*id,
		"POST")

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.NotContains(t, recorder.Body.String(), "error")
	parseErr(recorder)
}

func TestIndexAddress(t *testing.T) {

	recorder := httptest.NewRecorder()
	callApi([]byte(``),
		"user/address/list",
		"GET")

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.NotContains(t, recorder.Body.String(), "error")
	parseErr(recorder)
}

func TestShowAddress(t *testing.T) {
	id := getID(callApi([]byte(`

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
		"user/address/create",
		"POST"))
	if !assert.NotNilf(t, id, "Create model failed for get id") {
		return
	}

	recorder := httptest.NewRecorder()
	callApi([]byte(``),
		"user/address/show/"+*id,
		"GET")

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.NotContains(t, recorder.Body.String(), "error")
	parseErr(recorder)
}
