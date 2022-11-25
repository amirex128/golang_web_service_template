package category

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// UpdateCategory
// @Summary ویرایش دسته بندی ها
// @description کاربران برای دسته بندی کردن محصولات خود و مقالات خود از این دسته بندی ها استفاده میکنند که دو نوع میباشد نوع اول برای محصولات و نوع دوم ان برای مقالات این دو نوع از هم جدا هستن ولی از یک ای پی ای ساخته می شوند و نمایش داده میشوند
// @Tags category
// @Router       /user/category/update/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	false "شناسه " SchemaExample(1)
// @Param message body DTOs.UpdateCategory true "ورودی"
func UpdateCategory(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:updateCategory", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.UpdateCategory(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).UpdateCategory(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دسته بندی با موفقیت ویرایش شد",
	})
}
