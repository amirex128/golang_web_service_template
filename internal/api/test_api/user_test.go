package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var userID *string

func TestUpdateUser(t *testing.T) {
	assert.NotNilf(t, userID, "user id is nil")
	recorder := callApi([]byte(`
{
  "again_password": "123456789",
  "cart_number": "6037998125410760",
  "email": "amirex128@gmail.com",
  "firstname": "امیر",
  "gender": "man",
  "id": 1,
  "lastname": "شیردل",
  "mobile": "09024809750",
  "password": "123456789",
  "shaba": "IR820540102680020817909002"
}
	`),
		"/api/v1/user/user/update/"+*userID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
