package validations

import (
	"backend/internal/app/DTOs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTag(c *gin.Context) (DTOs.CreateTag, error) {
	var dto DTOs.CreateTag
	tags := ValidationTags{
		"Title": {
			"required": "عنوان الزامی است",
			"min":      "عنوان باید حداقل 3 کاراکتر باشد",
			"max":      "عنوان باید حداکثر 40 کاراکتر باشد",
		},
		"Slug": {
			"required": "نامک الزامی است",
			"min":      "نامک باید حداقل 3 کاراکتر باشد",
			"max":      "نامک باید حداکثر 40 کاراکتر باشد",
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

func AddTag(c *gin.Context) (DTOs.AddTag, error) {
	var dto DTOs.AddTag
	tags := ValidationTags{
		"PostID": {
			"required": "پست الزامی است",
			"numeric":  "پست باید عددی باشد",
		},
		"TagID": {
			"required": "تگ الزامی است",
			"numeric":  "تگ باید عددی باشد",
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

func IndexTag(c *gin.Context) (DTOs.IndexTag, error) {
	var dto DTOs.IndexTag
	tags := ValidationTags{}
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