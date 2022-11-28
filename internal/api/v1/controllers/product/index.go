package product

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// IndexProduct
// @Summary لیست محصولات
// @description محصولات میتوانند خصوصیات مختلفی داشته باشند که این خصوصیت ها شامل رنگ و اندازه میباشد
// @Tags product
// @Router       /user/product [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexProduct", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexProduct(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	products, err := models.NewMysqlManager(c).GetAllProductWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
	return
}
