package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func createAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createAddress", "request")
	defer span.End()
	dto, err := validations.CreateAddress(c)
	if err != nil {
		return
	}
	userID := models.GetUser(c)
	err = models.NewMainManager().CreateAddress(c, ctx, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت ایجاد شد",
	})
}

func updateAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updateAddress", "request")
	defer span.End()
	dto, err := validations.UpdateAddress(c)
	if err != nil {
		return
	}
	userID := models.GetUser(c)
	addressID := utils.StringToUint64(c.Param("id"))
	err = models.NewMainManager().UpdateAddress(c, ctx, dto, addressID, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت ویرایش شد",
	})
}

func deleteAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteAddress", "request")
	defer span.End()
	userID := models.GetUser(c)
	addressID := utils.StringToUint64(c.Param("id"))
	err := models.NewMainManager().DeleteAddress(c, ctx, addressID, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت حذف شد",
	})
}

func indexAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexAddress", "request")
	defer span.End()
	userID := models.GetUser(c)
	addresses, err := models.NewMainManager().IndexAddress(c, ctx, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"addresses": addresses,
	})
}
