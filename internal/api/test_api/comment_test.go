package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var commentID *string

func TestCreateComment(t *testing.T) {

	recorder := callApi([]byte(`
{
  "body": "متن نظر",
  "email": "amirex128@gmail.com",
  "name": "نام",
  "post_id": 1
}
	`),
		"/api/v1/comment/create",
		"POST")
	commentID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestApproveComment(t *testing.T) {

	assert.NotNilf(t, commentID, "commentID is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/comment/approve/"+*commentID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexComment(t *testing.T) {

	recorder := callApi([]byte(``),
		"/api/v1/user/comment/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeleteComment(t *testing.T) {
	assert.NotNilf(t, commentID, "commentID is nil")
	recorder := callApi([]byte(``),
		"/api/v1/user/comment/delete/"+*commentID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
