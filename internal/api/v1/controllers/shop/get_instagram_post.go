package shop

import (
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

// GetInstagramPost
// @Summary دریافت پست های اینتستاگرام فروشگاه و تبدیل آن ها به محصول
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop/instagram [get]
// @Param	Authorization	 header string	true "Authentication"
func GetInstagramPost(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:getInstagramPost", "request")
	c.Request.WithContext(ctx)
	defer span.End()

}
