package order

import (
	"github.com/amirex128/selloora_backend/internal/constants"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ApproveOrder
// @Summary تائید سفارش
// @description سفارشات بعد از ثبت شدن باید توسط ادمین تائید شوند و سپس به مرحله انتخاب سرویس ارسال کنند بروند
// @Tags order
// @Router       /user/order/approve/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه سفارش" SchemaExample(1)
func ApproveOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:approveOrder", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	orderID := utils.StringToUint64(c.Param("id"))
	userID := utils.GetUserID(c)
	order, err := models.NewMysqlManager(c).FindOrderByID(orderID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	if order.UserID != *userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "شما اجازه دسترسی به این سفارش را ندارید",
		})
		return
	}
	order.Status = constants.AcceptedOrderStatus
	err = models.NewMysqlManager(c).UpdateOrder(*order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در تایید سفارش",
		})
		return
	}
}
