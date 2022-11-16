package controllers

import (
	"github.com/amirex128/selloora_backend/internal/app/models"
	"github.com/amirex128/selloora_backend/internal/app/utils"
	"github.com/amirex128/selloora_backend/internal/app/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func IndexCategory(c *gin.Context) {
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

func CreateCategory(c *gin.Context) {
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

func UpdateCategory(c *gin.Context) {
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

func DeleteCategory(c *gin.Context) {
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
