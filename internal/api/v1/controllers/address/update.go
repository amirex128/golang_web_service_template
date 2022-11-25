package address

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// UpdateAddress
// @Summary ویرایش آدرس
// @description کاربران میتوانند برای خود لیستی از ادرس های مختلف ایجاد کنند تا هر بار به راحتی مشخص کنند محصول خود را میخواند از کدام ادرس ارسال نمایید
// @Tags address
// @Router       /user/address/update/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	false "شناسه آدرس" SchemaExample(1)
// @Param message body DTOs.UpdateAddress true "ورودی"
func UpdateAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:updateAddress", "request")
	c.Request.WithContext(ctx)
	defer span.End()

	dto, err := validations.UpdateAddress(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).UpdateAddress(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت ویرایش شد",
	})
}
