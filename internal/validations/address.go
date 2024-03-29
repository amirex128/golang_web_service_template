package validations

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"strings"
)

func CreateAddress(c *gin.Context) (DTOs.CreateAddress, error) {
	var dto DTOs.CreateAddress
	tags := ValidationTags{
		"Name": {
			"required": "عنوان آدرس الزامی میباشد",
		},
		"ProvinceID": {
			"required": "استان الزامی میباشد",
			"numeric":  "استان باید عددی باشد",
		},
		"CityID": {
			"required": "شهر الزامی میباشد",
			"numeric":  "شهر باید عددی باشد",
		},
		"Address": {
			"required": "آدرس الزامی میباشد",
		},
		"PostalCode": {
			"required": "کد پستی الزامی میباشد",
			"numeric":  "کد پستی باید عددی باشد",
		},
		"Mobile": {
			"required":   "شماره موبایل الزامی میباشد",
			"numeric":    "شماره موبایل باید عددی باشد",
			"startswith": "شماره موبایل باید با 09 شروع شود",
			"min":        "شماره موبایل باید 11 رقم باشد",
			"max":        "شماره موبایل باید 11 رقم باشد",
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

func UpdateAddress(c *gin.Context) (DTOs.UpdateAddress, error) {
	var dto DTOs.UpdateAddress
	tags := ValidationTags{
		"Name": {
			"required": "عنوان آدرس الزامی میباشد",
		},
		"ProvinceID": {
			"required": "استان الزامی میباشد",
			"numeric":  "استان باید عددی باشد",
		},
		"CityID": {
			"required": "شهر الزامی میباشد",
			"numeric":  "شهر باید عددی باشد",
		},
		"Address": {
			"required": "آدرس الزامی میباشد",
		},
		"PostalCode": {
			"required": "کد پستی الزامی میباشد",
			"numeric":  "کد پستی باید عددی باشد",
		},
		"Mobile": {
			"required":   "شماره موبایل الزامی میباشد",
			"numeric":    "شماره موبایل باید عددی باشد",
			"startswith": "شماره موبایل باید با 09 شروع شود",
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
func IndexAddress(c *gin.Context) (DTOs.IndexAddress, error) {
	var dto DTOs.IndexAddress
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

	dto.Page = utils.StringToUint32(c.Query("page"))
	dto.PageSize = utils.StringToUint32(c.Query("page_size"))
	dto.Search = c.Query("search")
	dto.Sort = c.Query("sort")
	return dto, nil
}
