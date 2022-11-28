package order

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// DeleteOrder
// @Summary حذف سفارش
// @description با ایجاد سفارش کاربر میتواند سفارش های بالای صفحه و پاین صفحه مربوط به قالب خود را کم و زیاد نماید
// @Tags order
// @Router       /user/order/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه" SchemaExample(1)
func DeleteOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteOrder", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	orderID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(c).DeleteOrder(orderID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت حذف شد",
	})
}
