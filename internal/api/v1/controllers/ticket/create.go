package ticket

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreateTicket
// @Summary ایجاد تیکت
// @description کاربران میتوانند در صورت بروزمشکل از طریق پنل خود برای مدیریت و تیم پشتیبانی ما تیکت ارسال نماییند
// @Tags ticket
// @Router       /user/ticket/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateTicket  	true "ورودی"
func CreateTicket(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createTicket", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateTicket(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	ticket, err := models.NewMysqlManager(c).CreateTicket(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تیکت شما با موفقیت ثبت شد و همکاران ما به زودی با شما تماس خواهد گرفت",
		"data":    ticket,
	})
}
