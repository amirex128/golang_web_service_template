package validations

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) (DTOs.CreateOrder, error) {
	var dto DTOs.CreateOrder
	tags := ValidationTags{
		"OrderItems": {
			"required": "محصولات الزامی است",
		},
		"CustomerID": {
			"required": "شناسه مشتری الزامی است",
		},
		"DiscountID": {
			"numeric": "شناسه تخفیف باید عددی باشد",
		},
		"ProductID": {
			"required": "شناسه محصول الزامی است",
			"numeric":  "شناسه محصول باید عددی باشد",
		},
		"OptionID": {
			"numeric": "شناسه گزینه باید عددی باشد",
		},
		"Count": {
			"required": "تعداد الزامی است",
		},
		"VerifyCode": {
			"required": "کد تایید الزامی است",
			"numeric":  "کد تایید باید عددی باشد",
			"min":      "کد تایید باید 4 رقمی باشد",
			"max":      "کد تایید باید 4 رقمی باشد",
		},
	}
	err := c.Bind(&dto)
	if err != nil {
		return dto, errorx.New("مقادیر ارسال شده نا درست میباشد", "validation", err)
	}

	err = validate.Struct(dto)
	err = validateTags(tags, err, c)
	if err != nil {
		return dto, err
	}
	return dto, nil
}

func SendOrder(c *gin.Context) (DTOs.SendOrder, error) {
	var dto DTOs.SendOrder
	tags := ValidationTags{
		"OrderID": {
			"required": "شناسه سفارش الزامی است",
			"numeric":  "شناسه سفارش باید عددی باشد",
		},
		"Courier": {
			"required": "نام پیک الزامی است",
			"oneof":    "نام پیک نامعتبر است",
		},
		"PackageSize": {
			"required": "اندازه بسته الزامی است",
		},
		"Weight": {
			"required": "وزن الزامی است",
			"numeric":  "وزن باید عددی باشد",
		},
		"Value": {
			"required": "ارزش الزامی است",
			"numeric":  "ارزش باید عددی باشد",
		},
		"AddressID": {
			"required": "شناسه آدرس الزامی است",
			"numeric":  "شناسه آدرس باید عددی باشد",
		},
	}
	err := c.Bind(&dto)
	if err != nil {
		return dto, errorx.New("مقادیر ارسال شده نا درست میباشد", "validation", err)
	}

	err = validate.Struct(dto)
	err = validateTags(tags, err, c)
	if err != nil {
		return dto, err
	}
	return dto, nil
}

func CalculateOrder(c *gin.Context) (DTOs.CalculateOrder, error) {
	var dto DTOs.CalculateOrder
	tags := ValidationTags{
		"PackageSize": {
			"required": "اندازه بسته الزامی است",
		},
		"Weight": {
			"required": "وزن الزامی است",
			"numeric":  "وزن باید عددی باشد",
		},
		"Value": {
			"required": "ارزش الزامی است",
			"numeric":  "ارزش باید عددی باشد",
		},
		"AddressID": {
			"required": "شناسه آدرس الزامی است",
			"numeric":  "شناسه آدرس باید عددی باشد",
		},
	}
	err := c.Bind(&dto)
	if err != nil {
		return dto, errorx.New("مقادیر ارسال شده نا درست میباشد", "validation", err)
	}

	err = validate.Struct(dto)
	err = validateTags(tags, err, c)
	if err != nil {
		return dto, err
	}
	return dto, nil
}
