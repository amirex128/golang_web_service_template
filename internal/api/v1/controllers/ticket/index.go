package ticket

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// IndexTicket
// @Summary لیست تیکت ها
// @description کاربران میتوانند در صورت بروزمشکل از طریق پنل خود برای مدیریت و تیم پشتیبانی ما تیکت ارسال نماییند
// @Tags ticket
// @Router       /user/ticket [get]
// @Param	Authorization	 header string	true "Authentication"
func IndexTicket(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexTicket", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	userID := models.GetUserID(c)
	dto, err := validations.IndexTicket(c)
	shops, err := models.NewMysqlManager(c).GetAllTicketWithPagination(dto, *userID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shops": shops,
	})
}
