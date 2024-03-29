package validations

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"strings"
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
		return dto, errorx.New("مقادیر ارسال شده نا درست میباشد", "validation", err)
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
		return dto, errorx.New("مقادیر ارسال شده نا درست میباشد", "validation", err)
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
		return dto, errorx.New("مقادیر ارسال شده نا درست میباشد", "validation", err)
	}

	err = validate.Struct(dto)
	err = validateTags(tags, err, c)
	if err != nil {
		return dto, err
	}
	id := strings.Replace(c.Param("id"), "/", "", -1)
	if id != "" {
		dto.ID = utils.StringToUint64(id)
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
		return dto, errorx.New("مقادیر ارسال شده نا درست میباشد", "validation", err)
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
