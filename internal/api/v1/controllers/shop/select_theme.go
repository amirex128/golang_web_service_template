package shop

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SelectThemeShop
// @Summary انتخاب یک قالب جدید برای فروشگاه
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop/select/theme [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.SelectThemeShop  	true "ورودی"
func SelectThemeShop(c *gin.Context) {
	dto, err := validations.SelectThemeShop(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(c).SelectThemeByID(dto.ThemeID, dto.ShopID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "قالب با موفقیت انتخاب گردید",
	})
}
