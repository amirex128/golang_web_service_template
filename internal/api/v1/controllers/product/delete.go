package product

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// DeleteProduct
// @Summary حذف محصول
// @description محصولات میتوانند خصوصیات مختلفی داشته باشند که این خصوصیت ها شامل رنگ و اندازه میباشد
// @Tags product
// @Router       /user/product/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه محصول" SchemaExample(1)
func DeleteProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteProduct", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))

	manager := models.NewMysqlManager(c)
	err := manager.DeleteProduct(id)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت حذف شد",
	})
	return
}
