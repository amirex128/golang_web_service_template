package address

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// IndexAddress
// @Summary لیست آدرس
// @description کاربران میتوانند برای خود لیستی از ادرس های مختلف ایجاد کنند تا هر بار به راحتی مشخص کنند محصول خود را میخواند از کدام ادرس ارسال نمایید
// @Tags address
// @Router       /user/address/list [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexAddress", "request")
	c.Request.WithContext(ctx)
	defer span.End()

	dto, err := validations.IndexAddress(c)
	addresses, err := models.NewMysqlManager(c).GetAllAddressWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": addresses,
	})
}
