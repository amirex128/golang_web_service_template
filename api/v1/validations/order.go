package validations

import (
	"backend/internal/app/DTOs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) (DTOs.CreateOrder, error) {
	var dto DTOs.CreateOrder
	tags := ValidationTags{
		"OrderItem": {
			"required": "محصولات الزامی است",
		},
		"CustomerID": {
			"required": "شناسه مشتری الزامی است",
		},
		"DiscountID": {
			"numeric": "شناسه تخفیف باید عددی باشد",
		},
		"ProductID": {
			"required": "شناسه محصول الزامی است",
			"numeric":  "شناسه محصول باید عددی باشد",
		},
		"OptionID": {
			"numeric": "شناسه گزینه باید عددی باشد",
		},
		"Count": {
			"required": "تعداد الزامی است",
		},
		"VerifyCode": {
			"required": "کد تایید الزامی است",
			"numeric":  "کد تایید باید عددی باشد",
			"min":      "کد تایید باید 4 رقمی باشد",
			"max":      "کد تایید باید 4 رقمی باشد",
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
