package validations

import (
	"backend/internal/app/DTOs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestCustomer(c *gin.Context) (DTOs.RequestCustomer, error) {
	var dto DTOs.RequestCustomer
	tags := ValidationTags{
		"Mobile": {
			"required":   "شماره موبایل الزامی است",
			"min":        "شماره موبایل باید 11 رقم باشد",
			"max":        "شماره موبایل باید 11 رقم باشد",
			"startswith": "شماره موبایل باید با 09 شروع شود",
		},
	}
	err := c.Bind(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مقادیر ارسال شده نا درست میباشد",
			"error":   err.Error(),
		})
		return dto, errors.New("validation error")
	}

	err = validate.Struct(dto)
	err = validateTags(tags, err, c)
	if err != nil {
		return dto, err
	}
	return dto, nil
}
func VerifyCustomer(c *gin.Context) (DTOs.VerifyCustomer, error) {
	var dto DTOs.VerifyCustomer
	tags := ValidationTags{
		"Mobile": {
			"required":   "شماره موبایل الزامی است",
			"min":        "شماره موبایل باید 11 رقم باشد",
			"max":        "شماره موبایل باید 11 رقم باشد",
			"startswith": "شماره موبایل باید با 09 شروع شود",
		},
		"VerifyCode": {
			"required": "کد تایید الزامی است",
			"min":      "کد تایید باید 4 رقم باشد",
			"max":      "کد تایید باید 4 رقم باشد",
		},
	}
	err := c.Bind(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مقادیر ارسال شده نا درست میباشد",
			"error":   err.Error(),
		})
		return dto, errors.New("validation error")
	}

	err = validate.Struct(dto)
	err = validateTags(tags, err, c)
	if err != nil {
		return dto, err
	}
	return dto, nil
}
func UpdateCustomer(c *gin.Context) (DTOs.UpdateCustomer, error) {
	var dto DTOs.UpdateCustomer
	tags := ValidationTags{
		"VerifyCode": {
			"required": "کد تایید الزامی است",
			"min":      "کد تایید باید 4 رقم باشد",
			"max":      "کد تایید باید 4 رقم باشد",
		},
		"FullName": {
			"required": "نام و نام خانوادگی الزامی است",
		},
		"ProvinceID": {
			"required": "استان الزامی است",
			"numeric":  "استان نامعتبر است",
		},
		"CityID": {
			"required": "شهر الزامی است",
			"numeric":  "شهر نامعتبر است",
		},
		"Address": {
			"required": "آدرس الزامی است",
		},
		"PostalCode": {
			"required": "کد پستی الزامی است",
			"numeric":  "کد پستی نامعتبر است",
		},
	}
	err := c.Bind(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مقادیر ارسال شده نا درست میباشد",
			"error":   err.Error(),
		})
		return dto, errors.New("validation error")
	}

	err = validate.Struct(dto)
	err = validateTags(tags, err, c)
	if err != nil {
		return dto, err
	}
	return dto, nil
}
