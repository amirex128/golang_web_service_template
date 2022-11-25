package customer

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// DeleteCustomer
// @Summary حذف مشتری
// @description
// @Tags customer
// @Router       /user/customer/delete/{id} [post]
// @Param	Authorization	header string	true "Authentication"
// @Param	id			path string	true "شناسه مشتری" SchemaExample(1)
func DeleteCustomer(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteCustomer", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	customerID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(c).DeleteCustomer(customerID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "مشتری با موفقیت حذف شد",
	})
}
