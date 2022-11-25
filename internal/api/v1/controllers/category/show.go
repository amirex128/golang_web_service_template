package category

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowCategory
// @Summary نمایش دسته بندی
// @description هر فروشگاه برای خود میتواند به تعداد دلخواه اسلایدر در موقعیت های مختلف مثل بالای صفحه و پایین صفحه ایجاد نماید
// @Tags category
// @Router       /user/category/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه " SchemaExample(1)
func ShowCategory(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showCategory", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	categoryID := c.Param("id")
	category, err := models.NewMysqlManager(c).FindCategoryByID(utils.StringToUint64(categoryID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}
