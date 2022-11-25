package shop

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreateShop
// @Summary ایجاد فروشگاه
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateShop  	true "ورودی"
func CreateShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createShop", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateShop(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	userID := models.GetUser(c)

	shop, err := models.NewMysqlManager(c).CreateShop(dto, *userID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "فروشگاه با موفقیت ایجاد شد",
		"data":    shop,
	})
}
