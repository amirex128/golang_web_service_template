package comment

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// DeleteComment
// @Summary حذف دیدگاه
// @description مدیریت نظرات و دیدگه هایی که کاربران در مورد محصولات و مقالات می ثبتند
// @Tags comment
// @Router       /user/comment/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه دیدگاه" SchemaExample(1)
func DeleteComment(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteCommentAdmin", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	id := c.Param("id")
	err := models.NewMysqlManager(c).DeleteComment(utils.StringToUint64(id))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دیدگاه با موفقیت حذف شد",
	})
}
