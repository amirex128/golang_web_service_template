package validations

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
)

func CreateGallery(c *gin.Context) (DTOs.CreateGallery, error) {
	var dto DTOs.CreateGallery
	tags := ValidationTags{
		"File": {
			"required": "فایل الزامی است",
		},
		"OwnerID": {
			"required": "شناسه مالک الزامی است",
		},
		"OwnerType": {
			"required": "نوع مالک الزامی است",
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
