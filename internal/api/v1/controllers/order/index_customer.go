package order

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// IndexCustomerOrders
// @Summary نمایش لیست سفارشات مشتری
// @description مشتری میتواند سفارشات خود را در یک پنل ساده مشاهده نمیاد
// @Tags order
// @Router       /customer/order/list [get]
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexCustomerOrders(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexCustomerOrders", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexOrderCustomer(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	customer, err := models.NewMysqlManager(c).FindCustomerByMobileAndVerifyCode(dto.Mobile, dto.VerifyCode)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	orders, err := models.NewMysqlManager(c).FindOrdersByCustomerID(customer.ID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}
