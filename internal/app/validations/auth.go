package validations

import (
	"errors"
	"github.com/amirex128/selloora_backend/internal/app/DTOs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestLoginRegister(c *gin.Context) (*DTOs.RequestLoginRegister, error) {
	var register = new(DTOs.RequestLoginRegister)
	tags := ValidationTags{
		"Mobile": {
			"required":   "شماره موبایل الزامی میباشد",
			"min":        "شماره موبایل باید 11 رقم باشد",
			"max":        "شماره موبایل باید 11 رقم باشد",
			"startswith": "شماره موبایل باید با 09 شروع شود",
		},
	}
	err := c.Bind(&register)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مقادیر ارسال شده نا درست میباشد",
			"type":    "validation",
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

func Verify(c *gin.Context) (DTOs.Verify, error) {
	var login DTOs.Verify
	tags := ValidationTags{
		"Mobile": {
			"required":   "شماره موبایل الزامی میباشد",
			"min":        "شماره موبایل باید 11 رقم باشد",
			"max":        "شماره موبایل باید 11 رقم باشد",
			"startswith": "شماره موبایل باید با 09 شروع شود",
		},
		"VerifyCode": {
			"min": "کد تایید باید حداقل 4 رقم باشد",
			"max": "کد تایید باید حداکثر 4 رقم باشد",
		},
	}
	err := c.Bind(&login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مقادیر ارسال شده نا درست میباشد",
			"type":    "validation",
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

func ChangePassword(c *gin.Context) (DTOs.ChangePassword, error) {
	var login DTOs.ChangePassword
	tags := ValidationTags{
		"Password": {
			"required": "رمز عبور الزامی میباشد",
			"min":      "رمز عبور باید حداقل 6 رقم باشد",
			"max":      "رمز عبور باید حداکثر 20 رقم باشد",
		},
		"AgainPassword": {
			"required": "تکرار رمز عبور الزامی میباشد",
			"min":      "تکرار رمز عبور باید حداقل 6 رقم باشد",
			"max":      "تکرار رمز عبور باید حداکثر 20 رقم باشد",
			"eqfield":  "رمز عبور و تکرار آن باید یکسان باشد",
		},
	}
	err := c.Bind(&login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مقادیر ارسال شده نا درست میباشد",
			"type":    "validation",
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

func ForgetPassword(c *gin.Context) (DTOs.ForgetPassword, error) {
	var login DTOs.ForgetPassword
	tags := ValidationTags{
		"Mobile": {
			"required":   "شماره موبایل الزامی میباشد",
			"min":        "شماره موبایل باید 11 رقم باشد",
			"max":        "شماره موبایل باید 11 رقم باشد",
			"startswith": "شماره موبایل باید با 09 شروع شود",
		},
	}
	err := c.Bind(&login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مقادیر ارسال شده نا درست میباشد",
			"type":    "validation",
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
