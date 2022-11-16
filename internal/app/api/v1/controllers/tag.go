package controllers

import (
	"github.com/amirex128/selloora_backend/internal/app/models"
	"github.com/amirex128/selloora_backend/internal/app/utils"
	"github.com/amirex128/selloora_backend/internal/app/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
	"net/http"
)

func CreateTag(c *gin.Context) {
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

func IndexTag(c *gin.Context) {
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

func DeleteTag(c *gin.Context) {
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

func AddTag(c *gin.Context) {
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
