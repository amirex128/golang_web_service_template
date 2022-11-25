package customer

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// IndexCustomer
// @Summary لیست مشتری
// @description
// @Tags customer
// @Router       /user/customer [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	customer			 query   string	false "شماره مشتری"
// @Param	customer_size		 query   string	false "تعداد مشتری"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexCustomer(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexCustomer", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexCustomer(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	customers, err := models.NewMysqlManager(c).GetAllCustomerWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": customers,
	})
}
