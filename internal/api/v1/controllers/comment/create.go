package comment

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreateComment
// @Summary ایجاد دیدگاه
// @description مدیریت نظرات و دیدگه هایی که کاربران در مورد محصولات و مقالات می ثبتند
// @Tags comment
// @Router       /comment/create [post]
// @Param comment body DTOs.CreateComment true "ورودی"
func CreateComment(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createComment", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateComment(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	comment, err := models.NewMysqlManager(c).CreateComment(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "نظر شما با موفقیت ثبت شد و پس از تایید مدیر نمایش داده خواهد شد",
		"data":    comment,
	})
}
