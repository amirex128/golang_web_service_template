package validations

import (
	"backend/internal/app/DTOs"
	"errors"
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
		"Title": {
			"required": "عنوان نظر الزامی میباشد",
			"min":      "عنوان نظر باید حداقل 3 کاراکتر باشد",
			"max":      "عنوان نظر باید حداکثر 255 کاراکتر باشد",
		},
		"Body": {
			"required": "متن نظر الزامی میباشد",
			"min":      "متن نظر باید حداقل 3 کاراکتر باشد",
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

func IndexComment(c *gin.Context) (DTOs.IndexComment, error) {
	var dto DTOs.IndexComment
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
