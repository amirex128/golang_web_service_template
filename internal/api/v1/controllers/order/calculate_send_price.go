package order

import (
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CalculateSendPrice
// @Summary دریافت اطلاعات ارسال و انتخاب ارسال کننده
// @description بعد از تائید سفارش باید اطلاعات سفارش از قبلی وزن وارد شود و هزینه ارسال هر سرویس دهنده محاسبه شود و توسط ادمین انتخاب شود سرویس دهنده جهت ارسال
// @Tags order
// @Router       /user/order/calculate [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CalculateOrder  	true "ورودی"
func CalculateSendPrice(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:calculateSendPrice", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CalculateOrder(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	err = utils.CalculateSendPriceTipax(dto)

	c.JSON(http.StatusOK, gin.H{
		"tipax": "",
	})
}
