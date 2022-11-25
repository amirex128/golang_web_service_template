package product

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowProduct
// @Summary نمایش محصول
// @description محصولات میتوانند خصوصیات مختلفی داشته باشند که این خصوصیت ها شامل رنگ و اندازه میباشد
// @Tags product
// @Router       /user/product/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه محصول" SchemaExample(1)
func ShowProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showProduct", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))

	manager := models.NewMysqlManager(c)
	product, err := manager.FindProductById(id)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
	return
}
