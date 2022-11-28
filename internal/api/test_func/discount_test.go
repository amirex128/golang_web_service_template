package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var discountID *string

func TestCreateDiscount(t *testing.T) {
	recorder := callApi([]byte(`
{
  "amount": 0,
  "code": "qwer",
  "count": 10,
  "ended_at": "2025-01-01 00:00:00",
  "percent": 50,
  "product_ids": [
    1
  ],
  "started_at": "2021-01-01 00:00:00",
  "status": true,
  "type": "percent"
}
	`),
		"/api/v1/user/discount/create",
		"POST")
	discountID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdateDiscount(t *testing.T) {
	assert.NotNilf(t, discountID, "discountID is nil")
	recorder := callApi([]byte(`
{
  "amount": 0,
  "code": "qwer",
  "count": 10,
  "ended_at": "2025-01-01 00:00:00",
  "id": 1,
  "percent": 40,
  "product_ids": [
    2
  ],
  "started_at": "2021-01-01 00:00:00",
  "status": true,
  "type": "percent"
}
	`),
		"/api/v1/user/discount/update/"+*discountID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowDiscount(t *testing.T) {
	assert.NotNilf(t, discountID, "discountID is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/discount/show/"+*discountID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestCheckDiscount(t *testing.T) {
	assert.NotNilf(t, discountID, "discountID is nil")
	recorder := callApi([]byte(`
{
  "code": "qwer",
  "product_ids": [
    {
      "count": 10,
      "product_id": 1
    }
  ],
  "user_id": 1
}
	`),
		"/api/v1/customer/discount/check",
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestIndexDiscount(t *testing.T) {
	recorder := callApi([]byte(``),
		"/api/v1/user/discount/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteDiscount(t *testing.T) {
	assert.NotNilf(t, discountID, "discountID is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/discount/delete/"+*discountID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
