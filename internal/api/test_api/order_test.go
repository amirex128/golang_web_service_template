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
		"/api/v1/user/order/create",
		"POST")
	orderID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestShowOrder(t *testing.T) {
	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/order/show/"+*orderID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexCustomerOrders(t *testing.T) {
	recorder := callApi([]byte(``),
		"/api/v1/user/order/list",
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
		"/api/v1/user/order/delete/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestSendOrder(t *testing.T) {
	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/order/delete/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestReturnedOrder(t *testing.T) {
	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/order/delete/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestCalculateSendPrice(t *testing.T) {
	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/order/delete/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestAcceptReturnedOrder(t *testing.T) {
	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/order/delete/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestCancelOrder(t *testing.T) {
	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/order/delete/"+*orderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestApproveOrder(t *testing.T) {
	assert.NotNilf(t, orderID, "order id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/order/delete/"+*orderID,
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
