package tag

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// DeleteTag
// @Summary حذف تگ
// @description مقالات یا محصولات برای سئو بیشتر میتوانند بر روی خود تگ قرار دهند تا برای کلمات کلیدی مورد نظر بهترین نتیجه را به دست آورند
// @Tags tag
// @Router       /user/tag/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه تگ" SchemaExample(1)
func DeleteTag(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteTag", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	id := c.Param("id")
	err := models.NewMysqlManager(c).DeleteTag(utils.StringToUint64(id))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تگ با موفقیت حذف شد",
	})
}
