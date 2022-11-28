package ticket

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// DeleteTicket
// @Summary حذف تیکت
// @description کاربران میتوانند در صورت بروزمشکل از طریق پنل خود برای مدیریت و تیم پشتیبانی ما تیکت ارسال نماییند
// @Tags ticket
// @Router       /user/ticket/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه" SchemaExample(1)
func DeleteTicket(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteTicket", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	ticketID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(c).DeleteTicket(ticketID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تیکت با موفقیت حذف شد",
	})
}
