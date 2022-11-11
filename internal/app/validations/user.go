package validations

import (
	"backend/internal/app/DTOs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateUser(c *gin.Context) (DTOs.UpdateUser, error) {
	var dto DTOs.UpdateUser
	tags := ValidationTags{
		"Gender": {
			"required": "جنسیت الزامی است",
			"oneof":    "جنسیت باید مرد یا زن باشد",
		},
		"FirstName": {
			"required": "نام الزامی است",
		},
		"LastName": {
			"required": "نام خانوادگی الزامی است",
		},
		"Email": {
			"required": "ایمیل الزامی است",
			"email":    "ایمیل نا درست میباشد",
		},
		"Mobile": {
			"required":   "موبایل الزامی است",
			"numeric":    "موبایل باید عددی باشد",
			"startswith": "موبایل باید با 09 شروع شود",
		},
		"CardNumber": {
			"required": "شماره کارت الزامی است",
			"numeric":  "شماره کارت باید عددی باشد",
		},
		"Shaba": {
			"required": "شماره شبا الزامی است",
			"numeric":  "شماره شبا باید عددی باشد",
		},
		"Password": {
			"min": "رمز عبور باید حداقل 6 رقم باشد",
			"max": "رمز عبور باید حداکثر 20 رقم باشد",
		},
		"AgainPassword": {
			"min":     "تکرار رمز عبور باید حداقل 6 رقم باشد",
			"max":     "تکرار رمز عبور باید حداکثر 20 رقم باشد",
			"eqfield": "رمز عبور و تکرار آن باید یکسان باشد",
		},
	}
	err := c.Bind(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مقادیر ارسال شده نا درست میباشد",
			"type":    "validation",
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
