package product

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreateProduct
// @Summary ایجاد محصول
// @description محصولات میتوانند خصوصیات مختلفی داشته باشند که این خصوصیت ها شامل رنگ و اندازه میباشد
// @Tags product
// @Router       /user/product/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateProduct  	true "ورودی"
func CreateProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createProduct", "request")
	c.Request.WithContext(ctx)
	defer span.End()

	dto, err := validations.CreateProduct(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	product, err := models.NewMysqlManager(c).CreateProduct(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ایجاد شد",
		"data":    product,
	})
	return
}
