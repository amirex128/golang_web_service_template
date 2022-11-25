package order

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowOrder
// @Summary نمایش جزئیات سفارش
// @description مشتری نیاز دارد سفارش خود را از طریق پنل مشتری مشاهده نماید
// @Tags order
// @Router       /user/order/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه سفارش" SchemaExample(1)
func ShowOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showOrder", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	orderID := utils.StringToUint64(c.Param("id"))
	order, err := models.NewMysqlManager(c).FindOrderWithItemByID(orderID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}
