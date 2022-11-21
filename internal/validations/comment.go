package validations

import (
	"errors"
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment(c *gin.Context) (DTOs.CreateComment, error) {
	var dto DTOs.CreateComment
	tags := ValidationTags{
		"PostID": {
			"required": "شناسه پست الزامی میباشد",
			"numeric":  "شناسه پست باید عددی باشد",
		},
		"Name": {
			"required": "عنوان نظر الزامی میباشد",
			"min":      "عنوان نظر باید حداقل 3 کاراکتر باشد",
			"max":      "عنوان نظر باید حداکثر 255 کاراکتر باشد",
		},
		"Body": {
			"required": "متن نظر الزامی میباشد",
			"min":      "متن نظر باید حداقل 3 کاراکتر باشد",
		},
		"Email": {
			"required": "ایمیل الزامی میباشد",
			"email":    "ایمیل نا درست میباشد",
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

func IndexComment(c *gin.Context) (DTOs.IndexComment, error) {
	var dto DTOs.IndexComment
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
