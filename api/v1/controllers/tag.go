package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createTag(c *gin.Context) {
	dto, err := validations.CreateTag(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().CreateTag(c, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تگ با موفقیت ایجاد شد",
	})

}

func indexTag(c *gin.Context) {
	dto, err := validations.IndexTag(c)
	if err != nil {
		return
	}
	pagination, err := models.NewMainManager().GetAllTagsWithPagination(c, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tags": pagination,
	})
}

func deleteTag(c *gin.Context) {
	id := c.Param("id")
	err := models.NewMainManager().DeleteTag(c, utils.StringToUint64(id))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تگ با موفقیت حذف شد",
	})
}

func addTag(c *gin.Context) {
	dto, err := validations.AddTag(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().AddTag(c, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تگ با موفقیت به پست اضافه شد",
	})
}
