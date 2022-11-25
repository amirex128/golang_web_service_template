package shop

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// UpdateShop
// @Summary ویرایش فروشگاه
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop/update/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	false "شناسه " SchemaExample(1)
// @Param	message	 body   DTOs.UpdateShop  	true "ورودی"
func UpdateShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:updateShop", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.UpdateShop(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	err = models.NewMysqlManager(c).UpdateShop(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "فروشگاه با موفقیت ویرایش شد",
	})
}
