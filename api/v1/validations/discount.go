package validations

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckDiscount(c *gin.Context) (DTOs.CheckDiscount, error) {
	var dto DTOs.CheckDiscount
	tags := ValidationTags{
		"Code": {
			"required": "کد تخفیف الزامی است",
		},
		"ProductIDs": {
			"required": "شناسه محصولات الزامی است",
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

func CreateDiscount(c *gin.Context) (DTOs.CreateDiscount, error) {
	var dto DTOs.CreateDiscount
	tags := ValidationTags{
		"Code": {
			"required": "کد تخفیف الزامی است",
		},
		"StartedAt": {
			"required": "تاریخ شروع تخفیف الزامی است",
		},
		"EndedAt": {
			"required": "تاریخ پایان تخفیف الزامی است",
		},
		"Count": {
			"required": "تعداد تخفیف الزامی است",
			"numeric":  "تعداد تخفیف باید عددی باشد",
		},
		"Type": {
			"required": "نوع تخفیف الزامی است",
		},
		"Amount": {
			"required": "مقدار تخفیف الزامی است",
			"numeric":  "مقدار تخفیف باید عددی باشد",
		},
		"Percent": {
			"required": "درصد تخفیف الزامی است",
			"numeric":  "درصد تخفیف باید عددی باشد",
		},
		"ProductIDs": {
			"required": "شناسه محصولات الزامی است",
		},
		"Status": {
			"required": "وضعیت تخفیف الزامی است",
			"numeric":  "وضعیت تخفیف باید عددی باشد",
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
func UpdateDiscount(c *gin.Context) (DTOs.UpdateDiscount, error) {
	var dto DTOs.UpdateDiscount
	tags := ValidationTags{
		"DiscountID": {
			"required": "شناسه تخفیف الزامی است",
			"numeric":  "شناسه تخفیف باید عددی باشد",
		},
		"Code": {
			"required": "کد تخفیف الزامی است",
		},
		"StartedAt": {
			"required": "تاریخ شروع تخفیف الزامی است",
		},
		"EndedAt": {
			"required": "تاریخ پایان تخفیف الزامی است",
		},
		"Count": {
			"required": "تعداد تخفیف الزامی است",
			"numeric":  "تعداد تخفیف باید عددی باشد",
		},
		"Type": {
			"required": "نوع تخفیف الزامی است",
		},
		"Amount": {
			"required": "مقدار تخفیف الزامی است",
			"numeric":  "مقدار تخفیف باید عددی باشد",
		},
		"Percent": {
			"required": "درصد تخفیف الزامی است",
			"numeric":  "درصد تخفیف باید عددی باشد",
		},
		"ProductIDs": {
			"required": "شناسه محصولات الزامی است",
		},
		"Status": {
			"required": "وضعیت تخفیف الزامی است",
			"numeric":  "وضعیت تخفیف باید عددی باشد",
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
func IndexDiscount(c *gin.Context) (DTOs.IndexDiscount, error) {
	var dto DTOs.IndexDiscount
	tags := ValidationTags{
		"UserID": {
			"required": "شناسه فروشگاه الزامی است",
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
	dto.Page = utils.StringToUint32(c.Query("page"))
	dto.PageSize = utils.StringToUint32(c.Query("page_size"))
	dto.Search = c.Query("search")
	dto.Sort = c.Query("sort")
	return dto, nil
}
