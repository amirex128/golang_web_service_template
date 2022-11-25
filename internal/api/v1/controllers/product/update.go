package product

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// UpdateProduct
// @Summary ویرایش محصول
// @description محصولات میتوانند خصوصیات مختلفی داشته باشند که این خصوصیت ها شامل رنگ و اندازه میباشد
// @Tags product
// @Router       /user/product/update/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	false "شناسه " SchemaExample(1)
// @Param	message	 body   DTOs.UpdateProduct  	true "ورودی"
func UpdateProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:updateProduct", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.UpdateProduct(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	manager := models.NewMysqlManager(c)

	err = manager.UpdateProduct(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ویرایش شد",
	})
	return
}
