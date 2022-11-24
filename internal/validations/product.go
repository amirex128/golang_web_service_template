package validations

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strings"
)

func IndexProduct(c *gin.Context) (DTOs.IndexProduct, error) {
	var dto DTOs.IndexProduct
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
	dto.ShopID = utils.StringToUint64(c.Query("shop_id"))

	dto.Page = utils.StringToUint32(c.Query("page"))
	dto.PageSize = utils.StringToUint32(c.Query("page_size"))
	dto.Search = c.Query("search")
	dto.Sort = c.Query("sort")
	return dto, nil
}
func CreateProduct(c *gin.Context) (DTOs.CreateProduct, error) {
	var dto DTOs.CreateProduct
	tags := ValidationTags{
		"ManufacturerId": {
			"numeric": "شناسه تولید کننده باید عددی باشد",
		},
		"Description": {
			"required": "توضیحات محصول الزامی است",
			"min":      "توضیحات محصول باید حداقل 3 کاراکتر باشد",
			"max":      "توضیحات محصول باید حداکثر 1000 کاراکتر باشد",
		},
		"Name": {
			"required": "نام محصول الزامی است",
			"min":      "نام محصول باید حداقل 3 کاراکتر باشد",
			"max":      "نام محصول باید حداکثر 100 کاراکتر باشد",
		},
		"ShortDescription": {
			"required": "توضیحات کوتاه محصول الزامی است",
			"min":      "توضیحات کوتاه محصول باید حداقل 3 کاراکتر باشد",
			"max":      "توضیحات کوتاه محصول باید حداکثر 300 کاراکتر باشد",
		},
		"Quantity": {
			"numeric": "تعداد محصول باید عددی باشد",
		},
		"Price": {
			"numeric": "قیمت محصول باید عددی باشد",
		},
		"FreeSend": {
			"numeric": "وضعیت ارسال رایگان باید عددی باشد",
		},
		"Weight": {
			"numeric": "وزن محصول باید عددی باشد",
		},
		"Height": {
			"numeric": "ارتفاع محصول باید عددی باشد",
		},
		"Width": {
			"numeric": "عرض محصول باید عددی باشد",
		},
		"StartedAt": {
			"datetime": "تاریخ شروع فروش محصول باید تاریخ باشد",
		},
		"EndedAt": {
			"datetime": "تاریخ پایان فروش محصول باید تاریخ باشد",
		},
		"DeliveryTime": {
			"numeric": "زمان تحویل محصول باید عددی باشد",
		},
		"OptionId": {
			"numeric": "شناسه گزینه محصول باید عددی باشد",
		},
		"OptionItemID": {
			"numeric": "شناسه آیتم گزینه محصول باید عددی باشد",
		},
	}
	err := c.BindWith(&dto, binding.FormMultipart)
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
func UpdateProduct(c *gin.Context) (DTOs.UpdateProduct, error) {
	var dto DTOs.UpdateProduct
	dto.ID = utils.StringToUint64(c.Param("id"))

	tags := ValidationTags{
		"ID": {
			"numeric": "شناسه محصول باید عددی باشد",
		},
		"ManufacturerId": {
			"numeric": "شناسه تولید کننده باید عددی باشد",
		},
		"Description": {
			"required": "توضیحات محصول الزامی است",
			"min":      "توضیحات محصول باید حداقل 3 کاراکتر باشد",
			"max":      "توضیحات محصول باید حداکثر 1000 کاراکتر باشد",
		},
		"Name": {
			"required": "نام محصول الزامی است",
			"min":      "نام محصول باید حداقل 3 کاراکتر باشد",
			"max":      "نام محصول باید حداکثر 100 کاراکتر باشد",
		},
		"ShortDescription": {
			"required": "توضیحات کوتاه محصول الزامی است",
			"min":      "توضیحات کوتاه محصول باید حداقل 3 کاراکتر باشد",
			"max":      "توضیحات کوتاه محصول باید حداکثر 300 کاراکتر باشد",
		},
		"Quantity": {
			"numeric": "تعداد محصول باید عددی باشد",
		},
		"Price": {
			"numeric": "قیمت محصول باید عددی باشد",
		},
		"FreeSend": {
			"numeric": "وضعیت ارسال رایگان باید عددی باشد",
		},
		"Weight": {
			"numeric": "وزن محصول باید عددی باشد",
		},
		"Height": {
			"numeric": "ارتفاع محصول باید عددی باشد",
		},
		"Width": {
			"numeric": "عرض محصول باید عددی باشد",
		},
		"StartedAt": {
			"datetime": "تاریخ شروع فروش محصول باید تاریخ باشد",
		},
		"EndedAt": {
			"datetime": "تاریخ پایان فروش محصول باید تاریخ باشد",
		},
		"DeliveryTime": {
			"numeric": "زمان تحویل محصول باید عددی باشد",
		},
		"OptionId": {
			"numeric": "شناسه گزینه محصول باید عددی باشد",
		},
		"OptionItemID": {
			"numeric": "شناسه آیتم گزینه محصول باید عددی باشد",
		},
	}
	err := c.BindWith(&dto, binding.FormMultipart)
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
