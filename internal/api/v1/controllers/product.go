package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
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
	userID := models.GetUser(c)

	dto, err := validations.CreateProduct(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	err = models.NewMysqlManager(c).CreateProduct(dto, *userID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
