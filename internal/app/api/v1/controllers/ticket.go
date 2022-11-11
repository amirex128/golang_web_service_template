package controllers

import (
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"backend/internal/app/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func CreateTicket(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createTicket", "request")
	defer span.End()
	dto, err := validations.CreateTicket(c)
	if err != nil {
		return
	}
	var userID uint64
	if dto.GuestMobile == "" {
		userID = models.GetUser(c)
	}
	err = models.NewMainManager().CreateTicket(c, ctx, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تیکت شما با موفقیت ثبت شد و همکاران ما به زودی با شما تماس خواهد گرفت",
	})
}
func IndexTicket(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexTicket", "request")
	defer span.End()
	userID := models.GetUser(c)
	dto, err := validations.IndexTicket(c)
	shops, err := models.NewMainManager().GetAllTicketWithPagination(c, ctx, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shops": shops,
	})
}

func ShowTicket(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "showTicket", "request")
	defer span.End()
	ticketID := utils.StringToUint64(c.Param("ticketID"))
	ticket, err := models.NewMainManager().GetTicketWithChildren(c, ctx, ticketID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tickets": ticket,
	})
}
