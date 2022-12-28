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

// IndexOrder
// @Summary لیست سفارشات
// @description لیست سفارشات بر اساس حالت های مختلف قابلیت فیلتر شدن داد
// @Tags order
// @Router       /user/order [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	order_status	 query   string	false "new,processing,returned,completed وضعیت سفارش" example(new)
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexOrder", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	orderStatus := c.Query("order_status")
	userID := utils.GetUserID(c)
	var orders []*models.Order
	var err error
	if orderStatus == "new" {
		orders, err = models.NewMysqlManager(c).GetOrders(*userID, []string{
			constants.PendingAcceptOrderStatus,
		})
	} else if orderStatus == "processing" {
		orders, err = models.NewMysqlManager(c).GetOrders(*userID, []string{
			constants.AcceptedOrderStatus,
			constants.PendingReceivePostOrderStatus,
			constants.ReceivedPostOrderStatus,
			constants.ReceivedCustomerOrderStatus,
			constants.PendingReceivePostReturnOrderStatus,
			constants.ReceivedPostReturnOrderStatus,
			constants.ReceivedOwnerOrderStatus,
		})
	} else if orderStatus == "returned" {
		orders, err = models.NewMysqlManager(c).GetOrders(*userID, []string{
			constants.PendingReturnOrderStatus,
			constants.AcceptedReturnOrderStatus,
			constants.RejectedReturnOrderStatus,
		})
	} else if orderStatus == "completed" {
		orders, err = models.NewMysqlManager(c).GetOrders(*userID, []string{
			constants.FinishedOrderStatus,
		})
	}
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})

}
