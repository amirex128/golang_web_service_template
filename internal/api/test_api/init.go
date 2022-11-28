package test_api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/amirex128/selloora_backend/internal/api"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp/v3"
	"net/http"
	"net/http/httptest"
	"strings"
)

var (
	host   string
	token  string
	runner *gin.Engine
)

func init() {
	host = "http://localhost:8585"
	token = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFtaXJleDEyOEBnbWFpbC5jb20iLCJleHAiOjUyNjgzMTg5OTYsImV4cGlyZV9hdCI6IiIsImZpcnN0bmFtZSI6Itin2YXbjNixIiwiaWQiOjEsImxhc3RuYW1lIjoi2LTbjNix2K_ZhNuMIiwibW9iaWxlIjoiMDkwMjQ4MDk3NTAiLCJvcmlnX2lhdCI6MTY2ODMyMjU5Niwic3RhdHVzIjoiIn0.x7BKuxw288cm1JsskGRD178UPmNz-xRwkWHtb0WsU74"
	gin.SetMode(gin.TestMode)
	runner = api.Runner()
	ctx, _ := context.WithCancel(context.Background())
	models.Initialize(ctx)

}

func callApi(body []byte, api string, method string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(method, host+api, bytes.NewBuffer(body))
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	runner.ServeHTTP(recorder, req)
	return recorder
}

func parse(recorder *httptest.ResponseRecorder) map[string]interface{} {
	var response map[string]interface{}
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		return nil
	}
	return response
}
func parseErr(recorder *httptest.ResponseRecorder) {
	if http.StatusOK != recorder.Code || strings.Contains(recorder.Body.String(), "error") {
		res := parse(recorder)
		pp.Println(res)
	}
}

func getID(recorder *httptest.ResponseRecorder) *string {
	var body map[string]interface{}
	if http.StatusOK != recorder.Code || strings.Contains(recorder.Body.String(), "error") {
		parse(recorder)
		return nil
	} else {
		body = parse(recorder)
		if body == nil {
			return nil
		}
	}

	var id string
	if res, ok := body["data"].(map[string]interface{}); ok {
		if resID, ok := res["id"]; ok {
			id = fmt.Sprintf("%v", uint64(resID.(float64)))
		} else {
			return nil
		}
	} else {
		return nil
	}
	return &id
}
