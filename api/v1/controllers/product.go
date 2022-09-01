package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/helpers"
	"backend/internal/app/models"
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
	userID := uint64(jwt.ExtractClaims(c)["id"].(float64))

	dto, err := validations.CreateProduct(c)
	if err != nil {
		return
	}

	images, err := helpers.UploadImages(c, dto.Images, userID)
	if err != nil {
		return
	}

	dto.ImagePath = images
	err = models.NewMainManager().CreateProduct(c, dto, userID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ایجاد شد",
	})
	return
}

func updateProduct(c *gin.Context) {
	userID := uint64(jwt.ExtractClaims(c)["id"].(float64))

	dto, err := validations.UpdateProduct(c)
	if err != nil {
		return
	}
	manager := models.NewMainManager()

	err = manager.CheckAccessProduct(c, dto.ID, userID)
	if err != nil {
		return
	}

	helpers.RemoveImages(dto.ImageRemove)

	images, err := helpers.UploadImages(c, dto.Images, userID)
	if err != nil {
		return
	}

	dto.ImagePath = append(dto.ImagePath, images...)
	err = manager.UpdateProduct(c, dto, userID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ویرایش شد",
	})
	return
}

func deleteProduct(c *gin.Context) {
	userID := uint64(jwt.ExtractClaims(c)["id"].(float64))
	id := helpers.Uint64Convert(c.Param("id"))

	manager := models.NewMainManager()
	err := manager.CheckAccessProduct(c, id, userID)
	if err != nil {
		return
	}
	err = manager.DeleteProduct(c, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت حذف شد",
	})
	return
}

func showProduct(c *gin.Context) {
	id := helpers.Uint64Convert(c.Param("id"))

	manager := models.NewMainManager()
	product, err := manager.FindProductById(c, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, product)
	return
}
