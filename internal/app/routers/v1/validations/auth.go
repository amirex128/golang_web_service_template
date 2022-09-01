package validations

import (
	"backend/internal/app/DTOs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) (*DTOs.Register, error) {
	var register = new(DTOs.Register)
	tags := ValidationTags{
		"Password": {
			"min":      "رمز عبور باید حداقل 8 کاراکتر باشد",
			"max":      "رمز عبور باید حداکثر 20 کاراکتر باشد",
			"required": "رمز عبور باید وارد شود",
		},
		"Mobile": {
			"required":   "شماره موبایل الزامی میباشد",
			"min":        "شماره موبایل باید 11 رقم باشد",
			"max":        "شماره موبایل باید 11 رقم باشد",
			"startswith": "شماره موبایل باید با 09 شروع شود",
		},
	}
	err := c.BindJSON(&register)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مقادیر ارسال شده نا درست میباشد",
			"error":   err.Error(),
		})
		return register, errors.New("validation error")
	}

	err = validate.Struct(register)
	err = validateTags(tags, err, c)
	if err != nil {
		return register, err
	}
	return register, nil
}

func Login(c *gin.Context) (DTOs.Login, error) {
	var login DTOs.Login
	tags := ValidationTags{
		"Password": {
			"min":      "رمز عبور باید حداقل 8 کاراکتر باشد",
			"max":      "رمز عبور باید حداکثر 20 کاراکتر باشد",
			"required": "رمز عبور باید وارد شود",
		},
		"Mobile": {
			"required":   "شماره موبایل الزامی میباشد",
			"min":        "شماره موبایل باید 11 رقم باشد",
			"max":        "شماره موبایل باید 11 رقم باشد",
			"startswith": "شماره موبایل باید با 09 شروع شود",
		},
	}
	err := c.BindJSON(&login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مقادیر ارسال شده نا درست میباشد",
			"error":   err.Error(),
		})
		return login, errors.New("validation error")
	}

	err = validate.Struct(login)
	err = validateTags(tags, err, c)
	if err != nil {
		return login, err
	}
	return login, nil
}
