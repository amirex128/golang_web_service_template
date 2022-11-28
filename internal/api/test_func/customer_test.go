package test_api

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var customerID *string
var verifyCode string

func TestRequestCreateLoginCustomer(t *testing.T) {
	recorder := callApi([]byte(`
{
  "mobile": "09024809750",
  "shop_id": 1
}
	`),
		"/api/v1/customer/login/register",
		"POST")
	customerID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestShowCustomer(t *testing.T) {
	assert.NotNilf(t, customerID, "customer id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/customer/show/"+*customerID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	var body map[string]map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	verifyCode = body["customer"]["verify_code"].(string)
	assert.Nilf(t, err, "error in unmarshal body: %s", err)

	parseErr(recorder)
}

func TestVerifyCreateLoginCustomer(t *testing.T) {
	assert.NotNilf(t, customerID, "customer id is nil")
	assert.NotEmptyf(t, verifyCode, "verify code is empty")
	recorder := callApi([]byte(fmt.Sprintf(`
{
  "address": "آدرس",
  "city_id": 1,
  "full_name": "نام",
  "mobile": "09024809750",
  "postal_code": "9111111111",
  "province_id": 1,
  "shop_id": 1,
  "verify_code": "%s"
}
	`, verifyCode)),
		"customer/verify",
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexCustomer(t *testing.T) {
	recorder := callApi([]byte(``),
		"/api/v1/user/customer/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteCustomer(t *testing.T) {
	assert.NotNilf(t, customerID, "customer id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/customer/delete/"+*customerID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
