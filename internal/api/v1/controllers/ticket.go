package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
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
	span, ctx := apm.StartSpan(c.Request.Context(), "createTicket", "request")
	defer span.End()
	dto, err := validations.CreateTicket(c)
	if err != nil {
		return
	}
	var userID uint64
	if dto.GuestMobile == "" {
		userID = *models.GetUser(c)
	}
	err = models.NewMysqlManager(ctx).CreateTicket(c, ctx, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تیکت شما با موفقیت ثبت شد و همکاران ما به زودی با شما تماس خواهد گرفت",
	})
}

// IndexTicket
// @Summary لیست تیکت ها
// @description کاربران میتوانند در صورت بروزمشکل از طریق پنل خود برای مدیریت و تیم پشتیبانی ما تیکت ارسال نماییند
// @Tags ticket
// @Router       /user/ticket [get]
// @Param	Authorization	 header string	true "Authentication"
func IndexTicket(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexTicket", "request")
	defer span.End()
	userID := models.GetUser(c)
	dto, err := validations.IndexTicket(c)
	shops, err := models.NewMysqlManager(ctx).GetAllTicketWithPagination(c, ctx, dto, *userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shops": shops,
	})
}

// ShowTicket
// @Summary نمایش تیکت های یک موضوع
// @description کاربران میتوانند در صورت بروزمشکل از طریق پنل خود برای مدیریت و تیم پشتیبانی ما تیکت ارسال نماییند
// @Tags ticket
// @Router       /user/ticket/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه تیکت های یک موضوع" SchemaExample(1)
func ShowTicket(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "showTicket", "request")
	defer span.End()
	ticketID := utils.StringToUint64(c.Param("ticketID"))
	ticket, err := models.NewMysqlManager(ctx).GetTicketWithChildren(c, ctx, ticketID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tickets": ticket,
	})
}
