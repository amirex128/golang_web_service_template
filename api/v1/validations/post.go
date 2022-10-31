package validations

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(c *gin.Context) (DTOs.CreatePost, error) {
	var dto DTOs.CreatePost
	tags := ValidationTags{
		"Title": {
			"required": "عنوان الزامی است",
		},
		"Body": {
			"required": "متن الزامی است",
		},
		"Thumbnail": {
			"required": "تصویر الزامی است",
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
func UpdatePost(c *gin.Context) (DTOs.UpdatePost, error) {
	var dto DTOs.UpdatePost
	tags := ValidationTags{
		"Title": {
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

func IndexPost(c *gin.Context) (DTOs.IndexPost, error) {
	var dto DTOs.IndexPost
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
