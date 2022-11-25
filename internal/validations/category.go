package validations

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"strings"
)

func IndexCategory(c *gin.Context) (DTOs.IndexCategory, error) {
	var dto DTOs.IndexCategory
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
func CreateCategory(c *gin.Context) (DTOs.CreateCategory, error) {
	var dto DTOs.CreateCategory
	tags := ValidationTags{
		"Name": {
			"required": "نام دسته بندی الزامی میباشد",
		},
		"Type": {
			"required": "نوع دسته بندی الزامی میباشد",
		},
		"ParentID": {
			"numeric": "شناسه والد باید عددی باشد",
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

func UpdateCategory(c *gin.Context) (DTOs.UpdateCategory, error) {
	var dto DTOs.UpdateCategory
	tags := ValidationTags{
		"ParentID": {
			"numeric": "شناسه والد باید عددی باشد",
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
