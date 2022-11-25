package address

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreateAddress
// @Summary ایجاد آدرس جدید
// @description کاربران میتوانند برای خود لیستی از ادرس های مختلف ایجاد کنند تا هر بار به راحتی مشخص کنند محصول خود را میخواند از کدام ادرس ارسال نمایید
// @Tags address
// @Router       /user/address/create [post]
// @Param	Authorization	 header string	false "Authentication"
// @Param message body DTOs.CreateAddress true "ورودی"
func CreateAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createAddress", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateAddress(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	address, err := models.NewMysqlManager(c).CreateAddress(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت ایجاد شد",
		"data":    address,
	})
}
