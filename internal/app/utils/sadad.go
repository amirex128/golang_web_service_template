package utils

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)
import "github.com/go-resty/resty/v2"

const (
	SadadMerchantId  = "108"
	SadadTerminalId  = "ATS32431"
	SadadTerminalKey = "MGNhNmRlOTllZjA0OWFkYTA3ZTMzZmE2"

	SadadPayUrl     = "https://sadad.shaparak.ir/VPG/Purchase"
	SadadPayTestUrl = "https://sandbox.banktest.ir/melli/sadad.shaparak.ir/VPG/Purchase"

	SadadVerifyUrl     = "https://sadad.shaparak.ir/VPG/api/v0/Advice/Verify"
	SadadVerifyTestUrl = "https://sandbox.banktest.ir/melli/sadad.shaparak.ir/VPG/api/v0/Advice/Verify"

	SadadRequestUrl     = "https://sadad.shaparak.ir/VPG/api/v0/Request/PaymentRequest"
	SadadRequestTestUrl = "https://sandbox.banktest.ir/melli/sadad.shaparak.ir/VPG/api/v0/Request/PaymentRequest"
)

type SadadRequestResult struct {
	ResCode     int    `json:"ResCode"`
	Token       string `json:"Token"`
	Description string `json:"Description"`
}
type SadadVerifyResult struct {
	ResCode       int    `json:"ResCode"`
	Amount        int    `json:"Amount"`
	Description   string `json:"Description"`
	RetrivalRefNo string `json:"RetrivalRefNo"`
	SystemTraceNo string `json:"SystemTraceNo"`
	OrderId       int    `json:"OrderId"`
}

func SadadPayRequest(c *gin.Context, orderID int, amount float32) error {
	body := map[string]interface{}{
		"MerchantId":    SadadMerchantId,
		"TerminalId":    SadadTerminalId,
		"Amount":        amount,
		"OrderId":       orderID,
		"LocalDateTime": "09/29/2022 7:12:29 am",
		"ReturnUrl":     "http://localhost:8585/api/v1/sadad/verify",
		"SignData": func() string {
			desEncrypt, _ := tripleDesECBEncrypt([]byte(fmt.Sprintf("%s;%d;%d", SadadTerminalId, orderID, int(amount))), []byte("123456789012345678901234"))
			return desEncrypt
		}(),
	}
	res := &SadadRequestResult{}
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetBody(body).
		SetResult(res).
		Post(SadadRequestTestUrl)
	if err != nil {
		return err
	}
	fmt.Println(resp, res)
	location := url.URL{Path: fmt.Sprintf("%s?Token=%s", SadadPayTestUrl, res.Token)}
	c.Redirect(http.StatusFound, location.RequestURI())
	return nil
}
func SadadVerify(c *gin.Context, orderID int, amount float32, resCode int, token string) error {
	body := map[string]interface{}{
		"Token": token,
		"SignData": func() string {
			desEncrypt, _ := tripleDesECBEncrypt([]byte(fmt.Sprintf("%s;%d;%d", SadadTerminalId, orderID, int(amount))), []byte("123456789012345678901234"))
			return desEncrypt
		}(),
	}
	client := resty.New()
	res := &SadadRequestResult{}
	resp, err := client.R().
		EnableTrace().
		SetBody(body).
		SetResult(res).
		Post(SadadVerifyTestUrl)
	if err != nil {
		return err
	}
	fmt.Println(resp, res)
	return nil

}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func tripleDesECBEncrypt(src, key []byte) (string, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	origData := PKCS5Padding(src, bs)
	if len(origData)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(origData))
	dst := out
	for len(origData) > 0 {
		block.Encrypt(dst, origData[:bs])
		origData = origData[bs:]
		dst = dst[bs:]
	}

	outString := base64.StdEncoding.EncodeToString(out)
	return outString, nil
}

func tripleDesECBDecrypt(src, key []byte) (string, error) {
	src, err := base64.StdEncoding.DecodeString(string(src))
	if err != nil {
		return "", err
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out), nil
}
