package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var orderID *string

func TestCreateOrder(t *testing.T) {
	recorder := callApi([]byte(`
{
 "order": "ادرس کامل",
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
		"user/order/create",
		"POST")
	orderID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdateOrder(t *testing.T) {
	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(`
{
  "order": "ادرس کامل",
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
		"user/order/update/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowOrder(t *testing.T) {
	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"user/order/show/"+*orderID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexOrder(t *testing.T) {
	recorder := callApi([]byte(``),
		"user/order/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteOrder(t *testing.T) {
	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"user/order/delete/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
