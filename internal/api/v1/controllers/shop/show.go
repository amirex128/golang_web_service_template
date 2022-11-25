package shop

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowShop
// @Summary نمایش فروشگاه
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه فروشگاه" SchemaExample(1)
func ShowShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showShop", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	shopID := utils.StringToUint64(c.Param("id"))
	shop, err := models.NewMysqlManager(c).FindShopByID(shopID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shop": shop,
	})
}
