package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createAddress(c *gin.Context) {
	dto, err := validations.CreateAddress(c)
	if err != nil {
		return
	}
	userID := models.GetUser(c)
	err = models.NewMainManager().CreateAddress(c, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت ایجاد شد",
	})
}

func updateAddress(c *gin.Context) {
	dto, err := validations.UpdateAddress(c)
	if err != nil {
		return
	}
	userID := models.GetUser(c)
	addressID := utils.StringToUint64(c.Param("id"))
	err = models.NewMainManager().UpdateAddress(c, dto, addressID, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت ویرایش شد",
	})
}

func deleteAddress(c *gin.Context) {
	userID := models.GetUser(c)
	addressID := utils.StringToUint64(c.Param("id"))
	err := models.NewMainManager().DeleteAddress(c, addressID, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت حذف شد",
	})
}

func indexAddress(c *gin.Context) {
	userID := models.GetUser(c)
	addresses, err := models.NewMainManager().IndexAddress(c, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"addresses": addresses,
	})
}
