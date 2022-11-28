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
  "body": "متن پیام",
  "gallery_id": 1,
  "guest_mobile": "09024809750",
  "guest_name": "امیر",
  "parent_id": 1,
  "title": "عنوان"
}
	`),
		"/api/v1/user/ticket/create",
		"POST")
	ticketID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestShowTicket(t *testing.T) {
	assert.NotNilf(t, ticketID, "ticket id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/ticket/show/"+*ticketID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexTicket(t *testing.T) {
	recorder := callApi([]byte(``),
		"/api/v1/user/ticket/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteTicket(t *testing.T) {
	assert.NotNilf(t, ticketID, "ticket id is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/ticket/delete/"+*ticketID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
