package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var shopID *string

func TestCreateShop(t *testing.T) {
	recorder := callApi([]byte(`
{
 "shop": "ادرس کامل",
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
		"user/shop/create",
		"POST")
	shopID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdateShop(t *testing.T) {
	assert.NotNilf(t, shopID, "shop id is nil")
	recorder := callApi([]byte(`
{
  "shop": "ادرس کامل",
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
		"user/shop/update/"+*shopID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowShop(t *testing.T) {
	assert.NotNilf(t, shopID, "shop id is nil")
	recorder := callApi([]byte(``),
		"user/shop/show/"+*shopID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexShop(t *testing.T) {
	recorder := callApi([]byte(``),
		"user/shop/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteShop(t *testing.T) {
	assert.NotNilf(t, shopID, "shop id is nil")
	recorder := callApi([]byte(``),
		"user/shop/delete/"+*shopID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
