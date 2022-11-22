package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
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
	span, ctx := apm.StartSpan(c.Request.Context(), "indexProduct", "request")
	defer span.End()
	dto, err := validations.IndexProduct(c)
	if err != nil {
		return
	}
	products, err := models.NewMysqlManager(ctx).GetAllProductWithPagination(c, ctx, dto)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
	return
}

// CreateProduct
// @Summary ایجاد محصول
// @description محصولات میتوانند خصوصیات مختلفی داشته باشند که این خصوصیت ها شامل رنگ و اندازه میباشد
// @Tags product
// @Router       /user/product/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateProduct  	true "ورودی"
func CreateProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createProduct", "request")
	defer span.End()
	userID := models.GetUser(c)

	dto, err := validations.CreateProduct(c)
	if err != nil {
		return
	}

	err = models.NewMysqlManager(ctx).CreateProduct(c, ctx, dto, *userID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ایجاد شد",
	})
	return
}

// UpdateProduct
// @Summary ویرایش محصول
// @description محصولات میتوانند خصوصیات مختلفی داشته باشند که این خصوصیت ها شامل رنگ و اندازه میباشد
// @Tags product
// @Router       /user/product/update [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.UpdateProduct  	true "ورودی"
func UpdateProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updateProduct", "request")
	defer span.End()
	dto, err := validations.UpdateProduct(c)
	if err != nil {
		return
	}
	manager := models.NewMysqlManager(ctx)

	err = manager.UpdateProduct(c, ctx, dto)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ویرایش شد",
	})
	return
}

// DeleteProduct
// @Summary حذف محصول
// @description محصولات میتوانند خصوصیات مختلفی داشته باشند که این خصوصیت ها شامل رنگ و اندازه میباشد
// @Tags product
// @Router       /user/product/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه محصول" SchemaExample(1)
func DeleteProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteProduct", "request")
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))

	manager := models.NewMysqlManager(ctx)
	err := manager.DeleteProduct(c, ctx, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت حذف شد",
	})
	return
}

// ShowProduct
// @Summary نمایش محصول
// @description محصولات میتوانند خصوصیات مختلفی داشته باشند که این خصوصیت ها شامل رنگ و اندازه میباشد
// @Tags product
// @Router       /user/product/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه محصول" SchemaExample(1)
func ShowProduct(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "showProduct", "request")
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))

	manager := models.NewMysqlManager(ctx)
	product, err := manager.FindProductById(c, ctx, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
	return
}
