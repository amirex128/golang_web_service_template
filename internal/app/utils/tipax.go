package utils

import (
	"github.com/amirex128/selloora_backend/internal/app/DTOs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	tipaxusername = "amirex128@gmail.com"
	tipaxpassword = "q6766581Amirex"
	tipaxApi      = "i5HoeTtqh2fWzmwmgHN6au+SupQ/OpwMCIWnxmDp1jM="
)

var (
	tipaxToken = ""
)

type j map[string]interface{}

func getToken() (string, error) {
	if tipaxToken == "" {
		body := j{
			"username": tipaxusername,
			"password": tipaxpassword,
			"apiKey":   tipaxApi,
		}
		res := &struct {
			AccessToken  string `json:"accessToken"`
			RefreshToken string `json:"refreshToken"`
		}{}

		client := resty.New()
		resp, err := client.R().
			EnableTrace().
			SetBody(body).
			SetResult(res).
			Post("http://212.33.205.90:262/api/v2/Account/token")

		if err != nil {
			return "", err
		}

		if resp.Status() != "200" {
			return "", nil
		}
		tipaxToken = res.AccessToken
		return res.AccessToken, nil
	} else {
		return tipaxToken, nil
	}
}

func GetCitiesTipax() error {
	token, err := getToken()
	if err != nil {
		return err
	}
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		Get("http://212.33.205.90:262/api/v2/Cities")

	if err != nil {
		return nil
	}

	if resp.Status() != "200" {
		return nil
	}

	return nil
}

func GetProvinceTipax() error {
	token, err := getToken()
	if err != nil {
		return err
	}
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		Get("http://212.33.205.90:262/api/v2/States")

	if err != nil {
		return nil
	}

	if resp.Status() != "200" {
		return nil
	}

	return nil
}

func TipaxSendOrderRequest(*gin.Context) error {
	body := j{
		"packages": []j{
			{
				"origin": j{
					"cityId":      0,
					"fullAddress": "string",
					"no":          "string",
					"floor":       "string",
					"unit":        "string",
					"postalCode":  "string",
					"description": "string",
					"latitude":    "string",
					"longitude":   "string",
					"beneficiary": j{
						"phone":    "string",
						"fullName": "string",
						"mobile":   "string",
					},
				},
				"destination": j{
					"cityId":      0,
					"fullAddress": "string",
					"no":          "string",
					"floor":       "string",
					"unit":        "string",
					"postalCode":  "string",
					"latitude":    "string",
					"longitude":   "string",
					"description": "string",
					"beneficiary": j{
						"phone":    "string",
						"fullName": "string",
						"mobile":   "string",
					},
				},
				"weight":             0,
				"packageValue":       0,
				"length":             0,
				"width":              0,
				"height":             0,
				"packingId":          0,
				"packageContentId":   0,
				"packType":           10,
				"description":        "string",
				"serviceId":          0,
				"enableLabelPrivacy": true,
				"paymentType":        10,
				"pickupType":         10,
				"distributionType":   10,
			},
		},
	}
	res := &struct {
		TrackingCodes string `json:"trackingCodes"`
		OrderId       string `json:"orderId"`
	}{}

	token, err := getToken()
	if err != nil {
		return err
	}
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetBody(body).
		SetResult(res).
		SetAuthToken(token).
		Post("/api/v2/Orders")

	if err != nil {
		return nil
	}

	if resp.Status() != "200" {
		return nil
	}

	return nil
}

// محاسبه هزینه ارسال
func GetPriceOrder(orderID int) error {
	token, err := getToken()
	if err != nil {
		return err
	}
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		Get(fmt.Sprintf("%s%d", "/api/v2/Orders/GetOrderAmountByOrderId/", orderID))

	if err != nil {
		return nil
	}

	if resp.Status() != "200" {
		return nil
	}

	return nil
}

// دریافت هزینه بسته بندی
func GetPackingPrice() error {
	token, err := getToken()
	if err != nil {
		return err
	}
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		Get("/api/v2/PackingPrices")

	if err != nil {
		return nil
	}

	if resp.Status() != "200" {
		return nil
	}

	return nil
}

// کنسل کردن سفارش ثبت شده
func CancelOrder(trackingCode string) error {
	token, err := getToken()
	if err != nil {
		return err
	}
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		Put(fmt.Sprintf("%s%d", "/api/v2/Parcels/Cancel/", trackingCode))

	if err != nil {
		return nil
	}

	if resp.Status() != "200" {
		return nil
	}

	return nil
}

// محاسبه هزینه ارسال
func CalculateSendPriceTipax(dto DTOs.CalculateOrder) error {
	body := j{
		"packageInputs": []j{
			{
				"origin": j{
					"cityId":      0,
					"fullAddress": "string",
					"no":          "string",
					"floor":       "string",
					"unit":        "string",
					"postalCode":  "string",
					"description": "string",
					"latitude":    "string",
					"longitude":   "string",
					"beneficiary": j{
						"phone":    "string",
						"fullName": "string",
						"mobile":   "string",
					},
				},
				"destination": j{
					"cityId":      0,
					"fullAddress": "string",
					"no":          "string",
					"floor":       "string",
					"unit":        "string",
					"postalCode":  "string",
					"latitude":    "string",
					"longitude":   "string",
					"description": "string",
					"beneficiary": j{
						"phone":    "string",
						"fullName": "string",
						"mobile":   "string",
					},
				},
				"weight":             0,
				"packageValue":       0,
				"length":             0,
				"width":              0,
				"height":             0,
				"packingId":          0,
				"packageContentId":   0,
				"packType":           10,
				"description":        "string",
				"serviceId":          0,
				"enableLabelPrivacy": true,
				"paymentType":        10,
				"pickupType":         10,
				"distributionType":   10,
			},
		},
	}
	type regularRate struct {
		ServiceId            int    `json:"serviceId"`
		ServiceTitle         string `json:"serviceTitle"`
		TransportationCost   int    `json:"transportationCost"`
		Compensation         int    `json:"compensation"`
		NoticeOfDelivery     int    `json:"noticeOfDelivery"`
		Tracking             int    `json:"tracking"`
		Packing              int    `json:"packing"`
		BeforeTaxAndDiscount int    `json:"beforeTaxAndDiscount"`
		Tax                  int    `json:"tax"`
		FinalPrice           int    `json:"finalPrice"`
		Discount             int    `json:"discount"`
	}
	res := &struct {
		PackNum            string `json:"packNum"`
		OriginCity         string `json:"originCity"`
		DestiationCity     string `json:"destiationCity"`
		RegularRate        regularRate
		RegularPlusRate    regularRate
		ExpressRate        regularRate
		SameDayExpressRate regularRate
		AirExpressRate     regularRate
	}{}

	token, err := getToken()
	if err != nil {
		return err
	}
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetBody(body).
		SetResult(res).
		SetAuthToken(token).
		Post("/api/v2/Pricing")

	if err != nil {
		return nil
	}

	if resp.Status() != "200" {
		return nil
	}

	return nil
}

func TrackingOrder(string) {

}
