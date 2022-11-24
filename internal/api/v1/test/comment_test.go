package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateComment(t *testing.T) {

	recorder := httptest.NewRecorder()
	callApi([]byte(`

		{
		 "address": "ادرس کامل",
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
		"user/address/create",
		"POST", recorder)

	if !assert.Equal(t, http.StatusOK, recorder.Code) || !assert.NotContains(t, recorder.Body.String(), "error") {
		parse(t, recorder)
		return
	}

}

func TestDeleteComment(t *testing.T) {

	recorder := httptest.NewRecorder()
	callApi([]byte(``),
		"user/address/delete/1",
		"POST", recorder)

	if !assert.Equal(t, http.StatusOK, recorder.Code) || !assert.NotContains(t, recorder.Body.String(), "error") {
		parse(t, recorder)
		return
	}
}

func TestUpdateComment(t *testing.T) {

	recorder := httptest.NewRecorder()
	callApi([]byte(`

		{
		  "address": "ادرس کامل",
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
		"user/address/update",
		"POST", recorder)

	if !assert.Equal(t, http.StatusOK, recorder.Code) || !assert.NotContains(t, recorder.Body.String(), "error") {
		parse(t, recorder)
		return
	}
}

func TestIndexComment(t *testing.T) {

	recorder := httptest.NewRecorder()
	callApi([]byte(``),
		"user/address/list",
		"GET", recorder)

	if !assert.Equal(t, http.StatusOK, recorder.Code) || !assert.NotContains(t, recorder.Body.String(), "error") {
		parse(t, recorder)
		return
	}
}

func TestShowComment(t *testing.T) {

	recorder := httptest.NewRecorder()
	callApi([]byte(``),
		"user/address/show/2",
		"GET", recorder)

	if !assert.Equal(t, http.StatusOK, recorder.Code) || !assert.NotContains(t, recorder.Body.String(), "error") {
		parse(t, recorder)
		return
	}
}
