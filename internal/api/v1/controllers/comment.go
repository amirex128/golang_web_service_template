package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func CreateComment(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createComment", "request")
	defer span.End()
	dto, err := validations.CreateComment(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).CreateComment(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "نظر شما با موفقیت ثبت شد و پس از تایید مدیر نمایش داده خواهد شد",
	})
}

func ApproveCommentAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "approveCommentAdmin", "request")
	defer span.End()
	id := c.Param("id")
	err := models.NewMysqlManager(ctx).ApproveComment(c, ctx, utils.StringToUint64(id))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دیدگاه با موفقیت تایید شد",
	})
}

func DeleteCommentAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteCommentAdmin", "request")
	defer span.End()
	id := c.Param("id")
	err := models.NewMysqlManager(ctx).DeleteComment(c, ctx, utils.StringToUint64(id))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دیدگاه با موفقیت حذف شد",
	})
}

func IndexCommentAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexCommentAdmin", "request")
	defer span.End()
	dto, err := validations.IndexComment(c)
	if err != nil {
		return
	}
	pagination, err := models.NewMysqlManager(ctx).GetAllCommentWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"comments": pagination,
	})
}
