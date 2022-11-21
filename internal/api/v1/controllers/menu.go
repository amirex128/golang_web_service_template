package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func CreateMenu(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createMenu", "request")
	defer span.End()
	dto, err := validations.CreateMenu(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).CreateMenu(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت ایجاد شد",
	})
}
func UpdateMenu(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updateMenu", "request")
	defer span.End()
	dto, err := validations.UpdateMenu(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).UpdateMenu(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت ویرایش شد",
	})
}
func DeleteMenu(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteMenu", "request")
	defer span.End()
	menuID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(ctx).DeleteMenu(c, ctx, menuID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت حذف شد",
	})
}
func IndexMenu(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexMenu", "request")
	defer span.End()
	dto, err := validations.IndexMenu(c)
	if err != nil {
		return
	}
	menus, err := models.NewMysqlManager(ctx).GetAllMenuWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"menus": menus,
	})
}
