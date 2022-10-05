package validations

import (
	"backend/internal/app/DTOs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateShop(c *gin.Context) (DTOs.CreateShop, error) {
	var dto DTOs.CreateShop
	tags := ValidationTags{
		"Name": {
			"required": "نام فروشگاه را وارد کنید",
		},
		"Type": {
			"required": "نوع فروشگاه را وارد کنید",
		},
		"SocialAddress": {
			"required": "آدرس شبکه اجتماعی را وارد کنید",
		},
		"Mobile": {
			"numeric":   "شماره موبایل باید عدد باشد",
			"statswith": "شماره موبایل باید با 09 شروع شود",
		},
		"Email": {
			"email": "ایمیل وارد شده نا درست میباشد",
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

func UpdateShop(c *gin.Context) (DTOs.UpdateShop, error) {
	var dto DTOs.UpdateShop
	tags := ValidationTags{
		"Mobile": {
			"numeric":   "شماره موبایل باید عدد باشد",
			"statswith": "شماره موبایل باید با 09 شروع شود",
		},
		"Email": {
			"email": "ایمیل وارد شده نا درست میباشد",
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

func CheckSocial(c *gin.Context) (DTOs.CheckSocial, error) {
	var dto DTOs.CheckSocial
	tags := ValidationTags{
		"SocialAddress": {
			"required": "آدرس شبکه اجتماعی را وارد کنید",
		},
		"Type": {
			"required": "نوع فروشگاه را وارد کنید",
		},
		"ShopID": {
			"required": "شناسه فروشگاه را وارد کنید",
			"numeric":  "شناسه فروشگاه باید عدد باشد",
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

func SendPrice(c *gin.Context) (DTOs.SendPrice, error) {
	var dto DTOs.SendPrice
	tags := ValidationTags{
		"SendPrice": {
			"required": "مبلغ ارسال را وارد کنید",
			"numeric":  "مبلغ ارسال باید عدد باشد",
		},
		"ShopID": {
			"required": "شناسه فروشگاه را وارد کنید",
			"numeric":  "شناسه فروشگاه باید عدد باشد",
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

func DeleteShop(c *gin.Context) (DTOs.DeleteShop, error) {
	var dto DTOs.DeleteShop
	tags := ValidationTags{
		"NewShopID": {
			"numeric": "شناسه فروشگاه جدید باید عدد باشد",
		},
		"ProductBehave": {
			"required": "رفتار محصولات را وارد کنید",
			"oneof":    "رفتار محصولات نا درست میباشد",
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
