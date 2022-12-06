package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestBlogLanding(t *testing.T) {

	recorder := callApi([]byte(``),
		"/blog",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	parseErr(recorder)
}

func TestCategoryLanding(t *testing.T) {

	recorder := callApi([]byte(``),
		"/category/1",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	parseErr(recorder)
}

func TestDetailsLanding(t *testing.T) {

	recorder := callApi([]byte(``),
		"/blog/test",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	parseErr(recorder)
}

func TestIndexLanding(t *testing.T) {

	recorder := callApi([]byte(``),
		"/",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	parseErr(recorder)
}

func TestPageLanding(t *testing.T) {

	recorder := callApi([]byte(``),
		"/page/test",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	parseErr(recorder)
}

func TestSearchLanding(t *testing.T) {

	recorder := callApi([]byte(``),
		"/search/test",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	parseErr(recorder)
}

func TestTagLanding(t *testing.T) {

	recorder := callApi([]byte(``),
		"/tag/test",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	parseErr(recorder)
}
