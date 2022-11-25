package category

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreateCategory
// @Summary ایجاد دسته بندی
// @description کاربران برای دسته بندی کردن محصولات خود و مقالات خود از این دسته بندی ها استفاده میکنند که دو نوع میباشد نوع اول برای محصولات و نوع دوم ان برای مقالات این دو نوع از هم جدا هستن ولی از یک ای پی ای ساخته می شوند و نمایش داده میشوند
// @Tags category
// @Router       /user/category/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param message body DTOs.CreateCategory true "ورودی"
func CreateCategory(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createCategory", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateCategory(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	category, err := models.NewMysqlManager(c).CreateCategory(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دسته بندی با موفقیت ایجاد شد",
		"data":    category,
	})
}
