package admin

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func createTicketAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createTicketAdmin", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateTicket(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	var userID uint64
	if dto.GuestMobile == "" {
		userID = models.GetUser(c)
	}
	_, err = models.NewMysqlManager(c).CreateTicket(c, ctx, dto, userID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تیکت با موفقیت ثبت شد",
	})
}
func indexTicketAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexTicketAdmin", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	userID := models.GetUser(c)
	dto, err := validations.IndexTicket(c)
	shops, err := models.NewMysqlManager(c).GetAllTicketWithPagination(c, ctx, dto, userID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shops": shops,
	})
}

func showTicketAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showTicketAdmin", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	ticketID := utils.StringToUint64(c.Param("ticketID"))
	ticket, err := models.NewMysqlManager(c).GetTicketWithChildren(c, ctx, ticketID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tickets": ticket,
	})
}
