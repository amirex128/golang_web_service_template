package validations

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
)

func RequestCreateLoginCustomer(c *gin.Context) (DTOs.RequestCreateLoginCustomer, error) {
	var dto DTOs.RequestCreateLoginCustomer
	tags := ValidationTags{
		"Mobile": {
			"required":   "شماره موبایل الزامی است",
			"min":        "شماره موبایل باید 11 رقم باشد",
			"max":        "شماره موبایل باید 11 رقم باشد",
			"startswith": "شماره موبایل باید با 09 شروع شود",
		},
		"ShopID": {
			"required": "شناسه فروشگاه الزامی است",
			"numeric":  "شناسه فروشگاه باید عددی باشد",
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

func CreateUpdateCustomer(c *gin.Context) (DTOs.CreateUpdateCustomer, error) {
	var dto DTOs.CreateUpdateCustomer
	tags := ValidationTags{
		"Mobile": {
			"min":        "شماره موبایل باید 11 رقم باشد",
			"max":        "شماره موبایل باید 11 رقم باشد",
			"startswith": "شماره موبایل باید با 09 شروع شود",
		},
		"VerifyCode": {
			"min": "کد تایید باید 4 رقم باشد",
			"max": "کد تایید باید 4 رقم باشد",
		},
		"ProvinceID": {
			"numeric": "شناسه استان باید عددی باشد",
		},
		"CityID": {
			"numeric": "شناسه شهر باید عددی باشد",
		},
		"PostalCode": {
			"startswith": "کد پستی باید با ۹ شروع شود",
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

func IndexOrderCustomer(c *gin.Context) (DTOs.IndexOrderCustomer, error) {
	var dto DTOs.IndexOrderCustomer
	tags := ValidationTags{
		"Mobile": {
			"required":   "شماره موبایل الزامی است",
			"min":        "شماره موبایل باید 11 رقم باشد",
			"max":        "شماره موبایل باید 11 رقم باشد",
			"startswith": "شماره موبایل باید با 09 شروع شود",
		},
		"verifyCode": {
			"required": "کد تایید الزامی است",
			"min":      "کد تایید باید 4 رقم باشد",
			"max":      "کد تایید باید 4 رقم باشد",
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

func IndexCustomer(c *gin.Context) (DTOs.IndexCustomer, error) {
	var dto DTOs.IndexCustomer
	tags := ValidationTags{}
	err := c.Bind(&dto)
	if err != nil {
		return dto, errorx.New("مقادیر ارسال شده نا درست میباشد", "validation", err)
	}

	err = validate.Struct(dto)
	err = validateTags(tags, err, c)
	if err != nil {
		return dto, err
	}
	dto.ShopID = utils.StringToUint64(c.Query("shop_id"))

	dto.Page = utils.StringToUint32(c.Query("page"))
	dto.PageSize = utils.StringToUint32(c.Query("page_size"))
	dto.Search = c.Query("search")
	dto.Sort = c.Query("sort")
	if dto.Search == "" {
		dto.Search = c.Param("search")
	}
	return dto, nil
}
