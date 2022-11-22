package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreatePage @Summary ایجاد صفحه
// @description هر فروشگاه میتواند به تعداد دلخواه صفحه ایجاد کند صفحات از دو حالت معمولی و خالی تشکیل میشوند که در حالت معمولی کاربر میتواند کل یک اچ تی ام ال را ذخیره نماید تا بدون چارچوب های قالب نمایش داده شود و در حالت معمولی همراه با چهار چوب ها نمایش داده شود
// @Tags page
// @Router       /user/page/create [post]
// @Param	Authorization	header string	true "Authentication"
// @Param	message	body DTOs.CreatePage 	true "ورودی"
func CreatePage(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createPage", "request")
	defer span.End()
	dto, err := validations.CreatePage(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).CreatePage(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت ایجاد شد",
	})
}

// UpdatePage
// @Summary ویرایش صفحه
// @description هر فروشگاه میتواند به تعداد دلخواه صفحه ایجاد کند صفحات از دو حالت معمولی و خالی تشکیل میشوند که در حالت معمولی کاربر میتواند کل یک اچ تی ام ال را ذخیره نماید تا بدون چارچوب های قالب نمایش داده شود و در حالت معمولی همراه با چهار چوب ها نمایش داده شود
// @Tags page
// @Router       /user/page/update [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.UpdatePage  	true "ورودی"
func UpdatePage(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updatePage", "request")
	defer span.End()
	dto, err := validations.UpdatePage(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).UpdatePage(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت ویرایش شد",
	})
}

// DeletePage
// @Summary حذف صفحه
// @description هر فروشگاه میتواند به تعداد دلخواه صفحه ایجاد کند صفحات از دو حالت معمولی و خالی تشکیل میشوند که در حالت معمولی کاربر میتواند کل یک اچ تی ام ال را ذخیره نماید تا بدون چارچوب های قالب نمایش داده شود و در حالت معمولی همراه با چهار چوب ها نمایش داده شود
// @Tags page
// @Router       /user/page/delete/{id} [post]
// @Param	Authorization	header string	true "Authentication"
// @Param	id			path string	true "شناسه صفحه" SchemaExample(1)
func DeletePage(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deletePage", "request")
	defer span.End()
	pageID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(ctx).DeletePage(c, ctx, pageID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت حذف شد",
	})
}

// IndexPage
// @Summary لیست صفحه
// @description هر فروشگاه میتواند به تعداد دلخواه صفحه ایجاد کند صفحات از دو حالت معمولی و خالی تشکیل میشوند که در حالت معمولی کاربر میتواند کل یک اچ تی ام ال را ذخیره نماید تا بدون چارچوب های قالب نمایش داده شود و در حالت معمولی همراه با چهار چوب ها نمایش داده شود
// @Tags page
// @Router       /user/page [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexPage(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexPage", "request")
	defer span.End()
	dto, err := validations.IndexPage(c)
	if err != nil {
		return
	}
	pages, err := models.NewMysqlManager(ctx).GetAllPageWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pages": pages,
	})
}
