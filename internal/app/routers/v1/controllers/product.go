package controllers

import (
	"backend/internal/app/models"
	"backend/internal/app/routers/v1/validations"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexProduct(c *gin.Context) {

	dto, err := validations.IndexProduct(c)
	if err != nil {
		return
	}
	products, err := models.NewMainManager().IndexProduct(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در دریافت اطلاعات"})
	}
	c.JSON(http.StatusOK, products)
	return
}

func createProduct(c *gin.Context) {
	dto, err := validations.CreateProduct(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().CreateProduct(dto, jwt.ExtractClaims(c))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در ایجاد محصول"})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ایجاد شد",
	})
	return
}

func updateProduct(c *gin.Context) {

}

func deleteProduct(c *gin.Context) {

}

func showProduct(c *gin.Context) {

}
