package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createComment(c *gin.Context) {
	dto, err := validations.CreateComment(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().CreateComment(c, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دیدگاه با موفقیت ایجاد شد",
	})
}

func acceptComment(c *gin.Context) {
	id := c.Param("id")
	err := models.NewMainManager().AcceptComment(c, utils.StringToUint64(id))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دیدگاه با موفقیت تایید شد",
	})
}

func deleteComment(c *gin.Context) {
	id := c.Param("id")
	err := models.NewMainManager().DeleteComment(c, utils.StringToUint64(id))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دیدگاه با موفقیت حذف شد",
	})
}

func indexComment(c *gin.Context) {
	dto, err := validations.IndexComment(c)
	if err != nil {
		return
	}
	pagination, err := models.NewMainManager().GetAllCommentWithPagination(c, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"comments": pagination,
	})
}
