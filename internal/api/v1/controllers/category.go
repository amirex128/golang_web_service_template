package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// IndexCategory
// @Summary لیست دسته بندی ها
// @description کاربران برای دسته بندی کردن محصولات خود و مقالات خود از این دسته بندی ها استفاده میکنند که دو نوع میباشد نوع اول برای محصولات و نوع دوم ان برای مقالات این دو نوع از هم جدا هستن ولی از یک ای پی ای ساخته می شوند و نمایش داده میشوند
// @Tags category
// @Router       /user/category [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexCategory(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexCategory", "request")
	defer span.End()
	dto, err := validations.IndexCategory(c)
	if err != nil {
		return
	}
	categories, err := models.NewMysqlManager(ctx).GetAllCategoryWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})

}

// CreateCategory
// @Summary ایجاد دسته بندی
// @description کاربران برای دسته بندی کردن محصولات خود و مقالات خود از این دسته بندی ها استفاده میکنند که دو نوع میباشد نوع اول برای محصولات و نوع دوم ان برای مقالات این دو نوع از هم جدا هستن ولی از یک ای پی ای ساخته می شوند و نمایش داده میشوند
// @Tags category
// @Router       /user/category/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param message body DTOs.CreateCategory true "ورودی"
func CreateCategory(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createCategory", "request")
	defer span.End()
	dto, err := validations.CreateCategory(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).CreateCategory(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دسته بندی با موفقیت ایجاد شد",
	})
}

// UpdateCategory
// @Summary ویرایش دسته بندی ها
// @description کاربران برای دسته بندی کردن محصولات خود و مقالات خود از این دسته بندی ها استفاده میکنند که دو نوع میباشد نوع اول برای محصولات و نوع دوم ان برای مقالات این دو نوع از هم جدا هستن ولی از یک ای پی ای ساخته می شوند و نمایش داده میشوند
// @Tags category
// @Router       /user/category/update [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param message body DTOs.UpdateCategory true "ورودی"
func UpdateCategory(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updateCategory", "request")
	defer span.End()
	dto, err := validations.UpdateCategory(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).UpdateCategory(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دسته بندی با موفقیت ویرایش شد",
	})
}

// DeleteCategory
// @Summary حذف دسته بندی
// @description کاربران برای دسته بندی کردن محصولات خود و مقالات خود از این دسته بندی ها استفاده میکنند که دو نوع میباشد نوع اول برای محصولات و نوع دوم ان برای مقالات این دو نوع از هم جدا هستن ولی از یک ای پی ای ساخته می شوند و نمایش داده میشوند
// @Tags category
// @Router       /user/category/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه آدرس" SchemaExample(1)
func DeleteCategory(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteCategory", "request")
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(ctx).DeleteCategory(c, ctx, id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دسته بندی با موفقیت حذف شد",
	})
}
