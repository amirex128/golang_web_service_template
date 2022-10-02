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
		"Title": {
			"required": "عنوان الزامی است",
			"min":      "عنوان باید حداقل 3 کاراکتر باشد",
			"max":      "عنوان باید حداکثر 40 کاراکتر باشد",
		},
		"Type": {
			"required": "نوع فروشگاه الزامی است",
			"oneof":    "نوع فروشگاه نامعتبر است",
		},
		"Social": {
			"required": "شبکه اجتماعی الزامی است",
		},
		"CategoryID": {
			"required": "دسته بندی الزامی است",
			"numeric":  "دسته بندی نامعتبر است",
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
		"Name": {
			"required": "عنوان الزامی است",
		},
		"CategoryID": {
			"numeric": "دسته بندی نامعتبر است",
		},
		"Mobile": {
			"startswith": "شماره موبایل نامعتبر است",
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
