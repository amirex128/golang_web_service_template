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
  "description": "توضیحات فروشگاه",
  "email": "amirex128@gmail.com",
  "gallery_id": 1,
  "instagram_id": "amirex_dev",
  "mobile": "09024809750",
  "name": "فروشگاه امیر",
  "phone": "05136643278",
  "send_price": 20000,
  "social_address": "amirex_dev",
  "telegram_id": "amirex128",
  "theme_id": 1,
  "type": "instagram",
  "website": "https://amirshirdel.ir",
  "whatsapp_id": "amirex128"
}
	`),
		"/api/v1/user/shop/create",
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
  "description": "توضیحات فروشگاه",
  "email": "amirex128@gmail.com",
  "gallery_id": 1,
  "id": 1,
  "instagram_id": "amirex_dev",
  "mobile": "09024809750",
  "name": "فروشگاه امیر",
  "phone": "05136643278",
  "send_price": 25000,
  "social_address": "amirex_dev",
  "telegram_id": "amirex128",
  "theme_id": 1,
  "type": "instagram",
  "website": "https://amirshirdel.ir",
  "whatsapp_id": "amirex128"
}
	`),
		"/api/v1/user/shop/update/"+*shopID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestSendPriceShop(t *testing.T) {
	assert.NotNilf(t, shopID, "shop id is nil")
	recorder := callApi([]byte(`
{
  "send_price": 30000,
  "shop_id": 1
}
	`),
		"/api/v1/user/shop/send-price/"+*shopID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestGetInstagramPost(t *testing.T) {
	assert.NotNilf(t, shopID, "shop id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/shop/instagram",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowShop(t *testing.T) {
	assert.NotNilf(t, shopID, "shop id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/shop/show/"+*shopID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexShop(t *testing.T) {
	recorder := callApi([]byte(``),
		"/api/v1/user/shop/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteShop(t *testing.T) {
	assert.NotNilf(t, shopID, "shop id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/shop/delete/"+*shopID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
