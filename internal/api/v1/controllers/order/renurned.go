package order

import (
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

// ReturnedOrder
// @Summary ثبت درخواست مرجوعی توسط مشتری
// @description مشتری میتواند بعد از دریافت سفارش ان را مرجوع کند
// @Tags order
// @Router       /user/order/returned [post]
func ReturnedOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:returnedOrder", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	//TODO

}
