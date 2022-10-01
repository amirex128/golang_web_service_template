package controllers

import (
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sadadPaymentRequest(c *gin.Context) {
	err := utils.SadadPayRequest(c, 10000000, 10000.0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
}

func sadadPaymentVerify(c *gin.Context) {
	err := utils.SadadVerify(c, 1, 1000.0, 100000, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
}
