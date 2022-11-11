package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func indexCategory(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexCategory", "request")
	defer span.End()
	dto, err := validations.IndexCategory(c)
	if err != nil {
		return
	}
	categories, err := models.NewMainManager().GetAllCategoryWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})

}

func createCategory(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createCategory", "request")
	defer span.End()
	dto, err := validations.CreateCategory(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().CreateCategory(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دسته بندی با موفقیت ایجاد شد",
	})
}

func updateCategory(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updateCategory", "request")
	defer span.End()
	dto, err := validations.UpdateCategory(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().UpdateCategory(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دسته بندی با موفقیت ویرایش شد",
	})
}

func deleteCategory(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteCategory", "request")
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))
	err := models.NewMainManager().DeleteCategory(c, ctx, id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دسته بندی با موفقیت حذف شد",
	})
}
