package customer

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowCustomer
// @Summary نمایش مشتری
// @description
// @Tags customer
// @Router       /user/customer/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه " SchemaExample(1)
func ShowCustomer(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showCustomer", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	customerID := c.Param("id")
	customer, err := models.NewMysqlManager(c).FindCustomerByID(utils.StringToUint64(customerID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": customer,
	})
}
