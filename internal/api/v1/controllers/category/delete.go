package category

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// DeleteCategory
// @Summary حذف دسته بندی
// @description کاربران برای دسته بندی کردن محصولات خود و مقالات خود از این دسته بندی ها استفاده میکنند که دو نوع میباشد نوع اول برای محصولات و نوع دوم ان برای مقالات این دو نوع از هم جدا هستن ولی از یک ای پی ای ساخته می شوند و نمایش داده میشوند
// @Tags category
// @Router       /user/category/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه دسته بندی" SchemaExample(1)
func DeleteCategory(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteCategory", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(c).DeleteCategory(id)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دسته بندی با موفقیت حذف شد",
	})
}
