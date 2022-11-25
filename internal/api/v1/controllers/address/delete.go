package address

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// DeleteAddress
// @Summary حذف آدرس
// @description کاربران میتوانند برای خود لیستی از ادرس های مختلف ایجاد کنند تا هر بار به راحتی مشخص کنند محصول خود را میخواند از کدام ادرس ارسال نمایید
// @Tags address
// @Router       /user/address/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه آدرس" SchemaExample(1)
func DeleteAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteAddress", "request")
	c.Request.WithContext(ctx)
	defer span.End()

	addressID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(c).DeleteAddress(addressID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت حذف شد",
	})
}
