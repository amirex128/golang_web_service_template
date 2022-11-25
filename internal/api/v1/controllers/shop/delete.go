package shop

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// DeleteShop
// @Summary حذف فروشگاه
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه فروشگاه" SchemaExample(1)
func DeleteShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteShop", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	shopID := utils.StringToUint64(c.Param("id"))
	userID := models.GetUser(c)
	dto, err := validations.DeleteShop(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	if dto.ProductBehave == "move" {
		err = models.NewMysqlManager(c).MoveProducts(shopID, dto.NewShopID, *userID)
		if err != nil {
			return
		}
	} else if dto.ProductBehave == "delete_product" {
		err = models.NewMysqlManager(c).DeleteProducts(shopID, *userID)
		if err != nil {
			return
		}
	}

	err = models.NewMysqlManager(c).DeleteShop(shopID, *userID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "فروشگاه با موفقیت حذف شد",
	})
}
