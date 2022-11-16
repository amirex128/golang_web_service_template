package api

import (
	"github.com/amirex128/selloora_backend/internal/pkg/framework/array"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/assert"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Call helper for api calls
func Call(ctx context.Context, method, url string, headers map[string]string, timeout time.Duration, pl interface{}, cookies []*http.Cookie) ([]byte, http.Header, int, error) {
	d, err := json.Marshal(pl)
	assert.Nil(err)
	var b io.Reader
	b = bytes.NewReader(d)
	method = strings.ToUpper(method)
	if array.StringInArray(method, "GET", "DELETE") {
		b = nil
	}
	r, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, nil, 0, errors.New("error while creating request")
	}
	for i := range headers {
		r.Header.Set(i, headers[i])
	}
	for i := range cookies {
		r.AddCookie(cookies[i])
	}
	nCtx, _ := context.WithTimeout(ctx, timeout)
	resp, err := http.DefaultClient.Do(r.WithContext(nCtx))
	if err != nil {
		return nil, nil, 0, errors.New("error in return response")
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()
	if err != nil {
		return nil, nil, resp.StatusCode, errors.New("error in reading response")
	}
	return data, resp.Header, resp.StatusCode, nil
}
