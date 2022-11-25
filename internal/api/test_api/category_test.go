package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var categoryID *string

func TestCreateCategory(t *testing.T) {
	recorder := callApi([]byte(`

		{
		  "description": "توضیحات دسته بندی",
		  "equivalent": "آموزش,یادگیری",
		  "gallery_id": 1,
		  "name": "نام دسته بندی",
		  "parent_id": 0,
		  "type": "product"
		}

	`),
		"user/category/create",
		"POST")
	categoryID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdateCategory(t *testing.T) {
	assert.NotNilf(t, categoryID, "categoryID is nil")
	recorder := callApi([]byte(`
		{
		  "description": "توضیحات دسته بندی",
		  "equivalent": "آموزش,یادگیری",
		  "gallery_id": 1,
		  "id": 1,
		  "name": "نام دسته بندی",
		  "parent_id": 0,
		  "sort": 1
		}
	`),
		"user/category/update/"+*categoryID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowCategory(t *testing.T) {
	assert.NotNilf(t, categoryID, "categoryID is nil")
	recorder := callApi([]byte(``),
		"user/category/show/"+*categoryID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexCategory(t *testing.T) {
	recorder := callApi([]byte(``),
		"user/category/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteCategory(t *testing.T) {
	assert.NotNilf(t, categoryID, "categoryID is nil")
	recorder := callApi([]byte(``),
		"user/category/delete/"+*categoryID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
