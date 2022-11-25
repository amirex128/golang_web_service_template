package order

import (
	"github.com/amirex128/selloora_backend/internal/constants"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// SendOrder
// @Summary دریافت اطلاعات ارسال و انتخاب ارسال کننده
// @description بعد از تائید سفارش باید اطلاعات سفارش از قبلی وزن وارد شود و هزینه ارسال هر سرویس دهنده محاسبه شود و توسط ادمین انتخاب شود سرویس دهنده جهت ارسال
// @Tags order
// @Router       /user/order/send [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.SendOrder  	true "ورودی"
func SendOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:sendOrder", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.SendOrder(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	order, err := models.NewMysqlManager(c).FindOrderByID(dto.OrderID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	if order.UserID != *models.GetUser(c) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "شما اجازه دسترسی به این سفارش را ندارید",
		})
		return
	}
	order.Courier = dto.Courier
	order.AddressID = dto.AddressID
	order.Status = constants.ChooseCourierOrderStatus
	order.Weight = dto.Weight
	order.PackageSize = dto.PackageSize
	order.LastUpdateStatusAt = utils.NowTime()

	err = models.NewMysqlManager(c).UpdateOrder(*order)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	if dto.Courier == "tipax" {
		err = utils.TipaxSendOrderRequest(c)
		if err != nil {
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اطلاعات سفارش ثبت شد در اتنظار پیک برای دریافت محصول باشید",
	})
}
