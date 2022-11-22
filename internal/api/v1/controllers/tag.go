package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
	"net/http"
)

// CreateTag
// @Summary ایجاد تگ
// @description مقالات یا محصولات برای سئو بیشتر میتوانند بر روی خود تگ قرار دهند تا برای کلمات کلیدی مورد نظر بهترین نتیجه را به دست آورند
// @Tags tag
// @Router       /user/tag/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateTag  	true "ورودی"
func CreateTag(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createTag", "request")
	defer span.End()
	dto, err := validations.CreateTag(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).CreateTag(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تگ با موفقیت ایجاد شد",
	})

}

// IndexTag
// @Summary لیست تگ ها
// @description مقالات یا محصولات برای سئو بیشتر میتوانند بر روی خود تگ قرار دهند تا برای کلمات کلیدی مورد نظر بهترین نتیجه را به دست آورند
// @Tags tag
// @Router       /user/tag [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexTag(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexTag", "request")
	defer span.End()
	dto, err := validations.IndexTag(c)
	if err != nil {
		return
	}
	pagination, err := models.NewMysqlManager(ctx).GetAllTagsWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tags": pagination,
	})
}

// DeleteTag
// @Summary حذف تگ
// @description مقالات یا محصولات برای سئو بیشتر میتوانند بر روی خود تگ قرار دهند تا برای کلمات کلیدی مورد نظر بهترین نتیجه را به دست آورند
// @Tags tag
// @Router       /user/tag/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه تگ" SchemaExample(1)
func DeleteTag(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteTag", "request")
	defer span.End()
	id := c.Param("id")
	err := models.NewMysqlManager(ctx).DeleteTag(c, ctx, utils.StringToUint64(id))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تگ با موفقیت حذف شد",
	})
}

// AddTag
// @Summary افزودن یک تگ به محصول یا مقاله
// @description مقالات یا محصولات برای سئو بیشتر میتوانند بر روی خود تگ قرار دهند تا برای کلمات کلیدی مورد نظر بهترین نتیجه را به دست آورند
// @Tags tag
// @Router       /user/tag/add [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateTag  	true "ورودی"
func AddTag(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "addTag", "request")
	defer span.End()
	dto, err := validations.AddTag(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).AddTag(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تگ با موفقیت به پست اضافه شد",
	})
}