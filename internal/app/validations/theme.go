package validations

import (
	"github.com/amirex128/selloora_backend/internal/app/DTOs"
	"github.com/amirex128/selloora_backend/internal/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTheme(c *gin.Context) (DTOs.CreateTheme, error) {
	var dto DTOs.CreateTheme
	tags := ValidationTags{
		"title": {
			"required": "عنوان الزامی است",
		},
		"body": {
			"required": "متن الزامی است",
		},
		"slug": {
			"required": "نامک الزامی است",
		},
		"shop_id": {
			"required": "فروشگاه الزامی است",
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

func IndexTheme(c *gin.Context) (DTOs.IndexTheme, error) {
	var dto DTOs.IndexTheme
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
