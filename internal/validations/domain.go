package validations

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
)

func CreateDomain(c *gin.Context) (DTOs.CreateDomain, error) {
	var dto DTOs.CreateDomain
	tags := ValidationTags{
		"title": {
			"required": "عنوان الزامی است",
		},
		"body": {
			"required": "متن الزامی است",
		},
		"slug": {
			"required": "نامک الزامی است",
		},
		"shop_id": {
			"required": "فروشگاه الزامی است",
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

func IndexDomain(c *gin.Context) (DTOs.IndexDomain, error) {
	var dto DTOs.IndexDomain
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
	return dto, nil
}
