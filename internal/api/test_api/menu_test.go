package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var menuID *string

func TestCreateMenu(t *testing.T) {
	recorder := callApi([]byte(`
{
  "link": "https://example.selloora.com/page/test",
  "name": "خانه",
  "parent_id": 0,
  "position": "top",
  "shop_id": 1
}
	`),
		"user/menu/create",
		"POST")
	menuID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdateMenu(t *testing.T) {
	assert.NotNilf(t, menuID, "menu id is nil")
	recorder := callApi([]byte(`
{
  "id": 1,
  "link": "https://example.selloora.com/page/test",
  "name": "خانه",
  "parent_id": 0,
  "position": "top",
  "sort": 1
}
	`),
		"user/menu/update/"+*menuID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowMenu(t *testing.T) {
	assert.NotNilf(t, menuID, "menu id is nil")
	recorder := callApi([]byte(``),
		"user/menu/show/"+*menuID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexMenu(t *testing.T) {
	recorder := callApi([]byte(``),
		"user/menu/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteMenu(t *testing.T) {
	assert.NotNilf(t, menuID, "menu id is nil")
	recorder := callApi([]byte(``),
		"user/menu/delete/"+*menuID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
