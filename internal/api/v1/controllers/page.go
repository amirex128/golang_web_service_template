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

// CreatePage @Summary ایجاد صفحه
// @description هر فروشگاه میتواند به تعداد دلخواه صفحه ایجاد کند صفحات از دو حالت معمولی و خالی تشکیل میشوند که در حالت معمولی کاربر میتواند کل یک اچ تی ام ال را ذخیره نماید تا بدون چارچوب های قالب نمایش داده شود و در حالت معمولی همراه با چهار چوب ها نمایش داده شود
// @Tags page
// @Router       /user/page/create [post]
// @Param	Authorization	header string	true "Authentication"
// @Param	message	body DTOs.CreatePage 	true "ورودی"
func CreatePage(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createPage", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreatePage(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).CreatePage(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:updatePage", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.UpdatePage(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).UpdatePage(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deletePage", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	pageID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(c).DeletePage(pageID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexPage", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexPage(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	pages, err := models.NewMysqlManager(c).GetAllPageWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pages": pages,
	})
}

// ShowPage
// @Summary نمایش صفحه
// @description هر فروشگاه میتواند به تعداد دلخواه صفحه ایجاد کند صفحات از دو حالت معمولی و خالی تشکیل میشوند که در حالت معمولی کاربر میتواند کل یک اچ تی ام ال را ذخیره نماید تا بدون چارچوب های قالب نمایش داده شود و در حالت معمولی همراه با چهار چوب ها نمایش داده شود
// @Tags post
// @Router       /user/page/show/{id} [get]
// @Param	Authorization	header string	true "Authentication"
// @Param	id			path string	true "شناسه" SchemaExample(1)
func ShowPage(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:ShowPage", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	pageID := c.Param("id")
	page, err := models.NewMysqlManager(c).FindPageByID(utils.StringToUint64(pageID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"page": page,
	})
}
