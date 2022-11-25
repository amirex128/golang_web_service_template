package ticket

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowTicket
// @Summary نمایش تیکت های یک موضوع
// @description کاربران میتوانند در صورت بروزمشکل از طریق پنل خود برای مدیریت و تیم پشتیبانی ما تیکت ارسال نماییند
// @Tags ticket
// @Router       /user/ticket/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه تیکت های یک موضوع" SchemaExample(1)
func ShowTicket(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showTicket", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	ticketID := utils.StringToUint64(c.Param("ticketID"))
	ticket, err := models.NewMysqlManager(c).GetTicketWithChildren(ticketID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tickets": ticket,
	})
}
