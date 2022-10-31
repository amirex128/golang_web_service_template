package validations

import (
	"backend/internal/app/DTOs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGallery(c *gin.Context) (DTOs.CreateGallery, error) {
	var dto DTOs.CreateGallery
	tags := ValidationTags{
		"File": {
			"required": "فایل الزامی است",
		},
		"OwnerID": {
			"required": "شناسه مالک الزامی است",
		},
		"OwnerType": {
			"required": "نوع مالک الزامی است",
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
