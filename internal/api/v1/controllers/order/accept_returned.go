package order

import (
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

// AcceptReturnedOrder
// @Summary تائید درخواست مرجوعی توسط مدیر
// @description بعد از درخواست مرجوعی با این درخواست توسط ادمین بررسی شود و در صورت تائید سفارش مرجوع شود و سرویس دهنده قبلی جهت جمع آوری ارسال شود
// @Tags order
// @Router       /user/order/returned/accept/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه" SchemaExample(1)
func AcceptReturnedOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:acceptReturnedOrder", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	//TODO

}
