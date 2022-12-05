package validations

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"strings"
)

func CreateShop(c *gin.Context) (DTOs.CreateShop, error) {
	var dto DTOs.CreateShop
	tags := ValidationTags{
		"Name": {
			"required": "نام فروشگاه را وارد کنید",
		},
		"EnglishName": {
			"required": "نام فروشگاه را وارد کنید",
		},
		"Type": {
			"required": "نوع فروشگاه را وارد کنید",
		},
		"SocialAddress": {
			"required": "آدرس شبکه اجتماعی را وارد کنید",
		},
		"Mobile": {
			"numeric":   "شماره موبایل باید عدد باشد",
			"statswith": "شماره موبایل باید با 09 شروع شود",
		},
		"Email": {
			"email": "ایمیل وارد شده نا درست میباشد",
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

func IndexShop(c *gin.Context) (DTOs.IndexShop, error) {
	var dto DTOs.IndexShop
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

func UpdateShop(c *gin.Context) (DTOs.UpdateShop, error) {
	var dto DTOs.UpdateShop
	tags := ValidationTags{
		"Mobile": {
			"numeric":   "شماره موبایل باید عدد باشد",
			"statswith": "شماره موبایل باید با 09 شروع شود",
		},
		"Email": {
			"email": "ایمیل وارد شده نا درست میباشد",
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

func CheckSocial(c *gin.Context) (DTOs.CheckSocial, error) {
	var dto DTOs.CheckSocial
	tags := ValidationTags{
		"SocialAddress": {
			"required": "آدرس شبکه اجتماعی را وارد کنید",
		},
		"Type": {
			"required": "نوع فروشگاه را وارد کنید",
		},
		"ShopID": {
			"required": "شناسه فروشگاه را وارد کنید",
			"numeric":  "شناسه فروشگاه باید عدد باشد",
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

func SendPriceShop(c *gin.Context) (DTOs.SendPriceShop, error) {
	var dto DTOs.SendPriceShop
	tags := ValidationTags{
		"SendPriceShop": {
			"required": "مبلغ ارسال را وارد کنید",
			"numeric":  "مبلغ ارسال باید عدد باشد",
		},
		"ShopID": {
			"required": "شناسه فروشگاه را وارد کنید",
			"numeric":  "شناسه فروشگاه باید عدد باشد",
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
		dto.ShopID = utils.StringToUint64(id)
	}
	return dto, nil
}

func SelectThemeShop(c *gin.Context) (DTOs.SelectThemeShop, error) {
	var dto DTOs.SelectThemeShop
	tags := ValidationTags{
		"ThemeID": {
			"required": "شناسه قالب را وارد کنید",
			"numeric":  "شناسه قالب باید عددی باشد",
		},
		"ShopID": {
			"required": "شناسه فروشگاه را وارد کنید",
			"numeric":  "شناسه فروشگاه باید عدد باشد",
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
func DeleteShop(c *gin.Context) (DTOs.DeleteShop, error) {
	var dto DTOs.DeleteShop
	tags := ValidationTags{
		"NewShopID": {
			"numeric": "شناسه فروشگاه جدید باید عدد باشد",
		},
		"ProductBehave": {
			"required": "رفتار حذف فروشگاه را وارد کنید",
			"oneof":    "رفتار حذف فروشگاه نا درست میباشد",
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
