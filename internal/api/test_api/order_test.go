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
  "customer_id": 1,
  "description": "توضیحات ارسال سفارش",
  "discount_code": "asdf",
  "order_items": [
    {
      "count": 10,
      "option_id": 1,
      "product_id": 1
    }
  ],
  "shop_id": 1,
  "user_id": 1,
  "verify_code": "1524"
}
	`),
		"/api/v1/customer/order/create",
		"POST")
	orderID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestShowOrder(t *testing.T) {

	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/customer/order/show/"+*orderID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexCustomerOrders(t *testing.T) {

	recorder := callApi([]byte(``),
		"/api/v1/customer/order/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexOrder(t *testing.T) {

	recorder := callApi([]byte(``),
		"/api/v1/user/order/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestTrackingOrder(t *testing.T) {

	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/customer/order/tracking/"+*orderID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestSendOrder(t *testing.T) {

	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(`
{
  "address_id": 1,
  "courier": "tipax",
  "order_id": 1,
  "package_size": "10x10x10",
  "value": 10000,
  "weight": 1000
}
`),
		"/api/v1/user/order/send/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestReturnedOrder(t *testing.T) {

	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/customer/order/returned/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestAcceptReturnedOrder(t *testing.T) {

	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/order/returned/accept/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestCancelOrder(t *testing.T) {
	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/customer/order/cancel/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestApproveOrder(t *testing.T) {
	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/order/approve/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteOrder(t *testing.T) {
	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/order/delete/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
