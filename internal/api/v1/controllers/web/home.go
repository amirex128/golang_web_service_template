package web

import "net/http"

// HomeWeb
// @Summary نمایش آدرس
// @description کاربران میتوانند برای خود لیستی از ادرس های مختلف ایجاد کنند تا هر بار به راحتی مشخص کنند محصول خود را میخواند از کدام ادرس ارسال نمایید
// @Tags web
// @Router       /user/address/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه فروشگاه " SchemaExample(1)
func HomeWeb(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:HomeWeb", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	shopID := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"data": address,
	})
}
