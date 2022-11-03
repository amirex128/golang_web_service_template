package admin

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createTicket(c *gin.Context) {
	dto, err := validations.CreateTicket(c)
	if err != nil {
		return
	}
	var userID uint64
	if dto.GuestMobile == "" {
		userID = models.GetUser(c)
	}
	err = models.NewMainManager().CreateTicket(c, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تیکت با موفقیت ثبت شد",
	})
}
func indexTicket(c *gin.Context) {
	userID := models.GetUser(c)
	dto, err := validations.IndexTicket(c)
	shops, err := models.NewMainManager().GetAllTicketWithPagination(c, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shops": shops,
	})
}

func showTicket(c *gin.Context) {
	ticketID := utils.StringToUint64(c.Param("ticketID"))
	ticket, err := models.NewMainManager().GetTicketWithChildren(c, ticketID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tickets": ticket,
	})
}
