package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var ticketID *string

func TestCreateTicket(t *testing.T) {
	recorder := callApi([]byte(`
{
 "ticket": "ادرس کامل",
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
		"user/ticket/create",
		"POST")
	ticketID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdateTicket(t *testing.T) {
	assert.NotNilf(t, ticketID, "ticket id is nil")
	recorder := callApi([]byte(`
{
  "ticket": "ادرس کامل",
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
		"user/ticket/update/"+*ticketID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowTicket(t *testing.T) {
	assert.NotNilf(t, ticketID, "ticket id is nil")
	recorder := callApi([]byte(``),
		"user/ticket/show/"+*ticketID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexTicket(t *testing.T) {
	recorder := callApi([]byte(``),
		"user/ticket/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteTicket(t *testing.T) {
	assert.NotNilf(t, ticketID, "ticket id is nil")
	recorder := callApi([]byte(``),
		"user/ticket/delete/"+*ticketID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
