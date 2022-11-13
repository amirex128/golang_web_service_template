package controllers

import (
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"backend/internal/app/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func CreatePage(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createPage", "request")
	defer span.End()
	dto, err := validations.CreatePage(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().CreatePage(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت ایجاد شد",
	})
}
func UpdatePage(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updatePage", "request")
	defer span.End()
	dto, err := validations.UpdatePage(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().UpdatePage(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت ویرایش شد",
	})
}
func DeletePage(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deletePage", "request")
	defer span.End()
	pageID := utils.StringToUint64(c.Param("id"))
	err := models.NewMainManager().DeletePage(c, ctx, pageID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت حذف شد",
	})
}
func IndexPage(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexPage", "request")
	defer span.End()
	dto, err := validations.IndexPage(c)
	if err != nil {
		return
	}
	pages, err := models.NewMainManager().GetAllPageWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pages": pages,
	})
}
