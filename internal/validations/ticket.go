package validations

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
)

func IndexTicket(c *gin.Context) (DTOs.IndexTicket, error) {
	var dto DTOs.IndexTicket
	tags := ValidationTags{
		"Page": {
			"min": "حداقل شماره صفحه باید 1 باشد",
		},
		"PageSize": {
			"min": "حداقل سایز صفحه باید 1 باشد",
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
func CreateTicket(c *gin.Context) (DTOs.CreateTicket, error) {
	var dto DTOs.CreateTicket
	tags := ValidationTags{
		"Page": {
			"min": "حداقل شماره صفحه باید 1 باشد",
		},
		"PageSize": {
			"min": "حداقل سایز صفحه باید 1 باشد",
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
