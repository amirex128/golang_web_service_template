package validations

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSlider(c *gin.Context) (DTOs.CreateSlider, error) {
	var dto DTOs.CreateSlider
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
func UpdateSlider(c *gin.Context) (DTOs.UpdateSlider, error) {
	var dto DTOs.UpdateSlider
	tags := ValidationTags{
		"Name": {
			"required": "عنوان الزامی است",
		},
		"Body": {
			"required": "متن الزامی است",
		},
		"Slug": {
			"required": "نامک الزامی است",
		},
		"CategoryID": {
			"required": "دسته بندی الزامی است",
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

func IndexSlider(c *gin.Context) (DTOs.IndexSlider, error) {
	var dto DTOs.IndexSlider
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
	dto.ShopID = utils.StringToUint64(c.Query("shop_id"))
	dto.WithoutPagination = c.Query("without_pagination") == "true"
	dto.Page = utils.StringToUint32(c.Query("page"))
	dto.PageSize = utils.StringToUint32(c.Query("page_size"))
	dto.Search = c.Query("search")
	dto.Sort = c.Query("sort")
	return dto, nil
}
