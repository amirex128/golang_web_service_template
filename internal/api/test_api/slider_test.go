package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var sliderID *string

func TestCreateSlider(t *testing.T) {
	recorder := callApi([]byte(`
{
  "description": "توضیحات اسلایدر",
  "gallery_id": 1,
  "link": "https://google.com",
  "position": "top",
  "shop_id": 1,
  "title": "عنوان اسلایدر"
}
	`),
		"user/slider/create",
		"POST")
	sliderID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdateSlider(t *testing.T) {
	assert.NotNilf(t, sliderID, "slider id is nil")
	recorder := callApi([]byte(`
{
  "description": "توضیحات اسلایدر",
  "gallery_id": 1,
  "id": 1,
  "link": "https://google.com",
  "position": "top",
  "shop_id": 1,
  "sort": 1,
  "title": "عنوان اسلایدر"
}
	`),
		"user/slider/update/"+*sliderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowSlider(t *testing.T) {
	assert.NotNilf(t, sliderID, "slider id is nil")
	recorder := callApi([]byte(``),
		"user/slider/show/"+*sliderID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexSlider(t *testing.T) {
	recorder := callApi([]byte(``),
		"user/slider/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteSlider(t *testing.T) {
	assert.NotNilf(t, sliderID, "slider id is nil")
	recorder := callApi([]byte(``),
		"user/slider/delete/"+*sliderID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code,"status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error","error found in response")

	parseErr(recorder)
}
