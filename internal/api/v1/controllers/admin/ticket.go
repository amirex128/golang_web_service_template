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
	span, ctx := apm.StartSpan(c.Request.Context(), "createTicketAdmin", "request")
	defer span.End()
	dto, err := validations.CreateTicket(c)
	if err != nil {
		return
	}
	var userID uint64
	if dto.GuestMobile == "" {
		userID = models.GetUser(c)
	}
	err = models.NewMysqlManager(ctx).CreateTicket(c, ctx, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تیکت با موفقیت ثبت شد",
	})
}
func indexTicketAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexTicketAdmin", "request")
	defer span.End()
	userID := models.GetUser(c)
	dto, err := validations.IndexTicket(c)
	shops, err := models.NewMysqlManager(ctx).GetAllTicketWithPagination(c, ctx, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shops": shops,
	})
}

func showTicketAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "showTicketAdmin", "request")
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
