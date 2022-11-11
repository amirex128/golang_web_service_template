package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func createTag(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createTag", "request")
	defer span.End()
	dto, err := validations.CreateTag(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().CreateTag(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تگ با موفقیت ایجاد شد",
	})

}

func indexTag(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexTag", "request")
	defer span.End()
	dto, err := validations.IndexTag(c)
	if err != nil {
		return
	}
	pagination, err := models.NewMainManager().GetAllTagsWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tags": pagination,
	})
}

func deleteTag(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteTag", "request")
	defer span.End()
	id := c.Param("id")
	err := models.NewMainManager().DeleteTag(c, ctx, utils.StringToUint64(id))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تگ با موفقیت حذف شد",
	})
}

func addTag(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "addTag", "request")
	defer span.End()
	dto, err := validations.AddTag(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().AddTag(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تگ با موفقیت به پست اضافه شد",
	})
}
