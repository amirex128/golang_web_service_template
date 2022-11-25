package post

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// DeletePost
// @Summary حدف مقاله
// @description فروشگاه ها میتوانند برای خود مقاله بسازند تا در سئو بهتر باشند
// @Tags post
// @Router       /user/post/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه " SchemaExample(1)
func DeletePost(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deletePost", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	postID := c.Param("id")
	err := models.NewMysqlManager(c).DeletePost(utils.StringToUint64(postID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله با موفقیت حذف شد",
	})
}
