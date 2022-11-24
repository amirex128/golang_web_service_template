package validations

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"strings"
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
		return dto, errorx.New("مقادیر ارسال شده نا درست میباشد", "validation", err)
	}

	err = validate.Struct(dto)
	err = validateTags(tags, err, c)
	if err != nil {
		return dto, err
	}
	id := strings.Replace(c.Param("id"), "/", "", -1)
	if dto.ID == 0 && id != "" {
		dto.ID = utils.StringToUint64(id)
	}
	return dto, nil
}
