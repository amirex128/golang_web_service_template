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
	"github.com/spf13/viper"
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"strings"
)

var (
	host   string
	token  string
	runner *gin.Engine
)

func init() {
	viper.AutomaticEnv()

	host = "http://localhost:8585"
	token = viper.GetString("token")
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
		pp.Fatalln("body has contain error")
		pp.Fatalln("------------------Stack-Trace------------------")
		debug.PrintStack()
		pp.Fatalln("------------------Error-Message------------------")
		pp.Fatalln(err)
		return nil
	}
	return response
}
func parseErr(recorder *httptest.ResponseRecorder) {
	if http.StatusOK != recorder.Code || strings.Contains(recorder.Body.String(), "error") {
		res := parse(recorder)
		pp.Fatalln("error when parsing body")
		pp.Fatalln("------------------Stack-Trace------------------")
		debug.PrintStack()
		pp.Fatalln("------------------Error-Message------------------")
		pp.Fatalln(res)
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
			pp.Fatal("body is nil for get id from response")
			pp.Fatalln("------------------Stack-Trace------------------")
			debug.PrintStack()
			return nil
		}
	}

	var id string
	if res, ok := body["data"].(map[string]interface{}); ok {
		if resID, ok := res["id"]; ok {
			id = fmt.Sprintf("%v", uint64(resID.(float64)))
		} else {
			pp.Fatal("id is invalid")
			pp.Fatalln("------------------Stack-Trace------------------")
			debug.PrintStack()
			return nil
		}
	} else {
		return nil
	}
	return &id
}
