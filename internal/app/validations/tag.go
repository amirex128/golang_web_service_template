package validations

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTag(c *gin.Context) (DTOs.CreateTag, error) {
	var dto DTOs.CreateTag
	tags := ValidationTags{
		"Name": {
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

func IndexTag(c *gin.Context) (DTOs.IndexTag, error) {
	var dto DTOs.IndexTag
	tags := ValidationTags{}
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

	dto.Page = utils.StringToUint32(c.Query("page"))
	dto.PageSize = utils.StringToUint32(c.Query("page_size"))
	dto.Search = c.Query("search")
	dto.Sort = c.Query("sort")
	return dto, nil
}
