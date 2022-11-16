package controllers

import (
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"backend/internal/app/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func IndexProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexProduct", "request")
	defer span.End()
	dto, err := validations.IndexProduct(c)
	if err != nil {
		return
	}
	products, err := models.NewMainManager().GetAllProductWithPagination(c, ctx, dto)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
	return
}

func CreateProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createProduct", "request")
	defer span.End()
	userID := models.GetUser(c)

	dto, err := validations.CreateProduct(c)
	if err != nil {
		return
	}

	err = models.NewMainManager().CreateProduct(c, ctx, dto, userID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ایجاد شد",
	})
	return
}

func UpdateProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updateProduct", "request")
	defer span.End()
	dto, err := validations.UpdateProduct(c)
	if err != nil {
		return
	}
	manager := models.NewMainManager()

	err = manager.UpdateProduct(c, ctx, dto)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ویرایش شد",
	})
	return
}

func DeleteProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteProduct", "request")
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))

	manager := models.NewMainManager()
	err := manager.DeleteProduct(c, ctx, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت حذف شد",
	})
	return
}

func ShowProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "showProduct", "request")
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))

	manager := models.NewMainManager()
	product, err := manager.FindProductById(c, ctx, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
	return
}
