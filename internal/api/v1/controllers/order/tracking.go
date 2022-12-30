package order

import (
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

// TrackingOrder
// @Summary پیگیری وضعیت ارسال سفارش
// @description مشتری میتواند سفارش خود را پیگیری نماید و مشاهده نماید که این سفارش در چه مرحله ای به سر میبرد
// @Tags order
// @Router       /customer/order/tracking/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه سفارش" SchemaExample(1)
func TrackingOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:trackingOrder", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	trackingCode := c.Param("id")
	utils.TrackingOrder(trackingCode)
}
