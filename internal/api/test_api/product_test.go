package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var productID *string

func TestCreateProduct(t *testing.T) {
	recorder := callApi([]byte(`
{
  "categoryID": 1,
  "description": "توضیحات محصول",
  "endedAt": "2025-01-01 00:00:00",
  "gallery_ids": [
    1
  ],
  "manufacturer": "سامسونگ",
  "name": "گوشی موبایل گلگسی نوت ۱۰",
  "optionId": 1,
  "optionItemID": 1,
  "price": 1000000,
  "quantity": 10,
  "shopID": 0,
  "startedAt": "2020-01-01 00:00:00"
}
	`),
		"user/product/create",
		"POST")
	productID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdateProduct(t *testing.T) {
	assert.NotNilf(t, productID, "product id is nil")
	recorder := callApi([]byte(`
{
  "active": true,
  "categoryID": 1,
  "description": "توضیحات محصول",
  "endedAt": "2025-01-01 00:00:00",
  "gallery_ids": [
    0
  ],
  "id": 1,
  "manufacturer": "سامسونگ",
  "name": "گوشی موبایل گلگسی نوت ۱۰",
  "optionId": 1,
  "optionItemID": 1,
  "price": 1000000,
  "quantity": 10,
  "shop_id": 2,
  "startedAt": "2020-01-01 00:00:00"
}
	`),
		"user/product/update/"+*productID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowProduct(t *testing.T) {
	assert.NotNilf(t, productID, "product id is nil")
	recorder := callApi([]byte(``),
		"user/product/show/"+*productID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexProduct(t *testing.T) {
	recorder := callApi([]byte(``),
		"user/product/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteProduct(t *testing.T) {
	assert.NotNilf(t, productID, "product id is nil")
	recorder := callApi([]byte(``),
		"user/product/delete/"+*productID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
