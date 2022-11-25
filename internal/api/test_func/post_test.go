package test_api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var postID *string

func TestCreatePost(t *testing.T) {
	recorder := callApi([]byte(`
{
  "body": "<p>متن مقاله</p>",
  "category_id": 1,
  "gallery_id": 1,
  "slug": "amoozesh-barnamenevisi",
  "title": "آموزش برنامه نویسی"
}
	`),
		"user/post/create",
		"POST")
	postID = getID(recorder)
	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)

}

func TestUpdatePost(t *testing.T) {
	assert.NotNilf(t, postID, "post id is nil")
	recorder := callApi([]byte(`
{
  "body": "<p>متن مقاله</p>",
  "category_id": 1,
  "gallery_id": 1,
  "id": 1,
  "slug": "amoozesh-barnamenevisi",
  "title": "آموزش برنامه نویسی"
}
	`),
		"user/post/update/"+*postID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestShowPost(t *testing.T) {
	assert.NotNilf(t, postID, "post id is nil")
	recorder := callApi([]byte(``),
		"user/post/show/"+*postID,
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestIndexPost(t *testing.T) {
	recorder := callApi([]byte(``),
		"user/post/list",
		"GET")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}

func TestDeletePost(t *testing.T) {
	assert.NotNilf(t, postID, "post id is nil")
	recorder := callApi([]byte(``),
		"user/post/delete/"+*postID,
		"POST")

	assert.Equalf(t, http.StatusOK, recorder.Code, "status code is not ok")
	assert.NotContainsf(t, recorder.Body.String(), "error", "error found in response")

	parseErr(recorder)
}
