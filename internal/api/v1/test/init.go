package test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/amirex128/selloora_backend/internal/api"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp/v3"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	host   string
	token  string
	runner *gin.Engine
)

func init() {
	host = "http://localhost:8585/api/v1/"
	token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFtaXJleDEyOEBnbWFpbC5jb20iLCJleHAiOjUyNjgzMTg5OTYsImV4cGlyZV9hdCI6IiIsImZpcnN0bmFtZSI6Itin2YXbjNixIiwiaWQiOjEsImxhc3RuYW1lIjoi2LTbjNix2K_ZhNuMIiwibW9iaWxlIjoiMDkwMjQ4MDk3NTAiLCJvcmlnX2lhdCI6MTY2ODMyMjU5Niwic3RhdHVzIjoiIn0.x7BKuxw288cm1JsskGRD178UPmNz-xRwkWHtb0WsU74"
	gin.SetMode(gin.TestMode)
	runner = api.Runner()
	ctx, _ := context.WithCancel(context.Background())
	models.Initialize(ctx)

}

func callApi(body []byte, api string, method string, recorder *httptest.ResponseRecorder) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, host+api, bytes.NewBuffer(body))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	runner.ServeHTTP(recorder, req)
	return recorder
}

func parse(t *testing.T, recorder *httptest.ResponseRecorder) {
	var response map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		assert.Errorf(t, err, "error in unmarshal response")
	}
	pp.Println("------------------Error Response-------------------")
	pp.Println(response)
}
