package shop

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// SendPriceShop
// @Summary ویرایش هزینه ارسال سفارشات به صورت جدا
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop/send-price/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.SendPriceShop  	true "ورودی"
func SendPriceShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:sendPrice", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.SendPriceShop(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).UpdateShop(DTOs.UpdateShop{
		SendPrice: dto.SendPrice,
	})
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "هزینه ارسال با موفقیت بروزرسانی شد",
	})
}
