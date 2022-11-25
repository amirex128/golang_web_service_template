package post

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowPost
// @Summary نمایش مقاله
// @description فروشگاه ها میتوانند برای خود مقاله بسازند تا در سئو بهتر باشند
// @Tags post
// @Router       /user/post/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه " SchemaExample(1)
func ShowPost(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showPost", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	postID := c.Param("id")
	post, err := models.NewMysqlManager(c).FindPostByID(utils.StringToUint64(postID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}
