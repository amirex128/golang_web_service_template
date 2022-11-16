package controllers

import (
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"backend/internal/app/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func CreateAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createAddress", "request")
	defer span.End()
	dto, err := validations.CreateAddress(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().CreateAddress(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت ایجاد شد",
	})
}

func UpdateAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updateAddress", "request")
	defer span.End()
	dto, err := validations.UpdateAddress(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().UpdateAddress(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت ویرایش شد",
	})
}

func DeleteAddress(c *gin.Context) {
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

func IndexAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexAddress", "request")
	defer span.End()
	dto, err := validations.IndexAddress(c)
	addresses, err := models.NewMainManager().GetAllAddressWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"addresses": addresses,
	})
}
