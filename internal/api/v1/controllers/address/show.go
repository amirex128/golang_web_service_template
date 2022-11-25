package address

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowAddress
// @Summary نمایش آدرس
// @description کاربران میتوانند برای خود لیستی از ادرس های مختلف ایجاد کنند تا هر بار به راحتی مشخص کنند محصول خود را میخواند از کدام ادرس ارسال نمایید
// @Tags address
// @Router       /user/address/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه " SchemaExample(1)
func ShowAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showAddress", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	addressID := c.Param("id")
	address, err := models.NewMysqlManager(c).FindAddressByID(utils.StringToUint64(addressID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": address,
	})
}
