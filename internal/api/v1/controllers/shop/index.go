package shop

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// IndexShop
// @Summary لیست فروشگاه
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexShop", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexShop(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	shops, err := models.NewMysqlManager(c).GetAllShopWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shops": shops,
	})
}
