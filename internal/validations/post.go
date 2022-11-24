package validations

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"strings"
)

func CreatePost(c *gin.Context) (DTOs.CreatePost, error) {
	var dto DTOs.CreatePost
	tags := ValidationTags{
		"Name": {
			"required": "عنوان الزامی است",
		},
		"Body": {
			"required": "متن الزامی است",
		},
		"GalleryID": {
			"required": "تصویر الزامی است",
		},
		"Slug": {
			"required": "نامک الزامی است",
		},
		"CategoryID": {
			"required": "دسته بندی الزامی است",
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
func UpdatePost(c *gin.Context) (DTOs.UpdatePost, error) {
	var dto DTOs.UpdatePost
	tags := ValidationTags{
		"Name": {
			"required": "عنوان الزامی است",
		},
		"Body": {
			"required": "متن الزامی است",
		},
		"Slug": {
			"required": "نامک الزامی است",
		},
		"CategoryID": {
			"required": "دسته بندی الزامی است",
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
	if dto.ID == 0 && id != "" {
		dto.ID = utils.StringToUint64(id)
	}
	return dto, nil
}

func IndexPost(c *gin.Context) (DTOs.IndexPost, error) {
	var dto DTOs.IndexPost
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
