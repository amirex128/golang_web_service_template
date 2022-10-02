package validations

import (
	"backend/internal/app/DTOs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexCategory(c *gin.Context) (DTOs.IndexCategory, error) {
	var dto DTOs.IndexCategory
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
