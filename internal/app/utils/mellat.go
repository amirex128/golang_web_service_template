package utils

import (
	"fmt"
	"github.com/tiaguinho/gosoap"
	"log"
	"net/http"
	"time"
)

const (
	MellatMerchantId = "134753151"
	MellatUsername   = "user134753151"
	MellatPassword   = "41011575"

	MellatPayUrl     = "https://bpm.shaparak.ir/pgwchannel/startpay.mellat"
	MellatPayTestUrl = "https://sandbox.banktest.ir/mellat/bpm.shaparak.ir/pgwchannel/startpay.mellat"

	MellatRequestUrl     = "https://bpm.shaparak.ir/pgwchannel/services/pgw?wsdl"
	MellatRequestTestUrl = "https://sandbox.banktest.ir/mellat/bpm.shaparak.ir/pgwchannel/services/pgw?wsdl"
)

func MellatPayRequest() {
	httpClient := &http.Client{
		Timeout: 1500 * time.Millisecond,
	}
	soap, err := gosoap.SoapClient(MellatPayTestUrl, httpClient)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}

	params := gosoap.Params{
		"terminalId":     MellatMerchantId,
		"userName":       MellatUsername,
		"userPassword":   MellatPassword,
		"orderId":        1,
		"amount":         1000,
		"localDate":      time.Now().Format("20060102"),
		"localTime":      time.Now().Format("150405"),
		"additionalData": "",
		"callBackUrl":    "http://localhost:8585/api/v1/mellat/verify",
		"payerId":        0,
	}

	res, err := soap.Call("bpPayRequest", params)
	if err != nil {
		log.Fatalf("Call error: %s", err)
	}

	fmt.Println(res)

}
func MellatVerify() {

}
