package validations

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
)

func CreateTag(c *gin.Context) (DTOs.CreateTag, error) {
	var dto DTOs.CreateTag
	tags := ValidationTags{
		"Name": {
			"required": "عنوان الزامی است",
			"min":      "عنوان باید حداقل 3 کاراکتر باشد",
			"max":      "عنوان باید حداکثر 40 کاراکتر باشد",
		},
		"Slug": {
			"required": "نامک الزامی است",
			"min":      "نامک باید حداقل 3 کاراکتر باشد",
			"max":      "نامک باید حداکثر 40 کاراکتر باشد",
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

func AddTag(c *gin.Context) (DTOs.AddTag, error) {
	var dto DTOs.AddTag
	tags := ValidationTags{
		"PostID": {
			"required": "پست الزامی است",
			"numeric":  "پست باید عددی باشد",
		},
		"TagID": {
			"required": "تگ الزامی است",
			"numeric":  "تگ باید عددی باشد",
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

func IndexTag(c *gin.Context) (DTOs.IndexTag, error) {
	var dto DTOs.IndexTag
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
