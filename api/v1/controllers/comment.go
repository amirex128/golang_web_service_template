package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func createComment(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createComment", "request")
	defer span.End()
	dto, err := validations.CreateComment(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().CreateComment(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "نظر شما با موفقیت ثبت شد و پس از تایید مدیر نمایش داده خواهد شد",
	})
}

func approveCommentAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "approveCommentAdmin", "request")
	defer span.End()
	id := c.Param("id")
	err := models.NewMainManager().ApproveComment(c, ctx, utils.StringToUint64(id))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دیدگاه با موفقیت تایید شد",
	})
}

func deleteCommentAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteCommentAdmin", "request")
	defer span.End()
	id := c.Param("id")
	err := models.NewMainManager().DeleteComment(c, ctx, utils.StringToUint64(id))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دیدگاه با موفقیت حذف شد",
	})
}

func indexCommentAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexCommentAdmin", "request")
	defer span.End()
	dto, err := validations.IndexComment(c)
	if err != nil {
		return
	}
	pagination, err := models.NewMainManager().GetAllCommentWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"comments": pagination,
	})
}
