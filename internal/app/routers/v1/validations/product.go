package validations

import (
	"backend/internal/app/DTOs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexProduct(c *gin.Context) (DTOs.IndexProduct, error) {
	var dto DTOs.IndexProduct
	tags := ValidationTags{
		"Page": {
			"min": "حداقل شماره صفحه باید 1 باشد",
		},
		"PageSize": {
			"min": "حداقل سایز صفحه باید 1 باشد",
		},
	}
	err := c.BindJSON(&dto)
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
func CreateProduct(c *gin.Context) (DTOs.CreateProduct, error) {
	var dto DTOs.CreateProduct
	tags := ValidationTags{}
	err := c.BindJSON(&dto)
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
