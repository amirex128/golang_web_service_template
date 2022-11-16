package validations

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
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
func IndexAddress(c *gin.Context) (DTOs.IndexAddress, error) {
	var dto DTOs.IndexAddress
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
