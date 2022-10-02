package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexProduct(c *gin.Context) {

	dto, err := validations.IndexProduct(c)
	if err != nil {
		return
	}
	products, err := models.NewMainManager().GetAllProductWithPagination(c, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در دریافت اطلاعات"})
	}
	c.JSON(http.StatusOK, products)
	return
}

func createProduct(c *gin.Context) {
	userID := utils.GetUser(c)

	dto, err := validations.CreateProduct(c)
	if err != nil {
		return
	}

	images, err := utils.UploadMultiImage(c, dto.Images, "product/user_"+utils.Uint64ToString(userID))
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
	userID := utils.GetUser(c)

	dto, err := validations.UpdateProduct(c)
	if err != nil {
		return
	}
	manager := models.NewMainManager()

	err = manager.CheckAccessProduct(c, dto.ID, userID)
	if err != nil {
		return
	}

	utils.RemoveImages(dto.ImageRemove)

	if dto.Images != nil {
		images, err := utils.UploadMultiImage(c, dto.Images, "product/user_"+utils.Uint64ToString(userID))
		if err != nil {
			return
		}

		dto.ImagePath = append(dto.ImagePath, images...)
	}

	err = manager.UpdateProduct(c, dto)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ویرایش شد",
	})
	return
}

func deleteProduct(c *gin.Context) {
	userID := utils.GetUser(c)
	id := utils.StringToUint64(c.Param("id"))

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
	id := utils.StringToUint64(c.Param("id"))

	manager := models.NewMainManager()
	product, err := manager.FindProductById(c, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, product)
	return
}
