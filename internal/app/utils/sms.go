package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	username = "09024809750"
	password = "#DB4Z"
	from     = "50004001809750"
)

func makeRequest(jsonData map[string]string, op string, c *gin.Context) error {

	jsonValue, _ := json.Marshal(jsonData)
	response, err := http.Post("https://rest.payamak-panel.com/api/SendSMS/"+op, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در ارسال پیامک لطفا مجدد تلاش کنید",
		})
		return errors.New("sms faild")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	return nil
}

func SendSMS(c *gin.Context, to string, text string, isFlash bool) error {

	jsonData := map[string]string{
		"username": username,
		"password": password,
		"to":       to,
		"from":     from,
		"text":     text,
		"isFlash":  strconv.FormatBool(isFlash),
	}
	err := makeRequest(jsonData, "SendSMS", c)
	if err != nil {
		return err
	}
	return nil
}

func SendByBaseNumber(c *gin.Context, text string, to string, bodyId int64) error {

	jsonData := map[string]string{
		"username": username,
		"password": password,
		"text":     text,
		"to":       to,
		"bodyId":   strconv.FormatInt(bodyId, 10),
	}
	err := makeRequest(jsonData, "BaseServiceNumber", c)
	if err != nil {
		return err
	}
	return nil
}
