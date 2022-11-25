package comment

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowComment
// @Summary نمایش نظر
// @description مدیریت نظرات و دیدگه هایی که کاربران در مورد محصولات و مقالات می ثبتند
// @Tags comment
// @Router       /user/comment/show/{id} [get]
// @Param	Authorization	header string	true "Authentication"
// @Param	id			path string	true "شناسه" SchemaExample(1)
func ShowComment(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:ShowComment", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	commentID := c.Param("id")
	comment, err := models.NewMysqlManager(c).FindCommentByID(utils.StringToUint64(commentID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}
