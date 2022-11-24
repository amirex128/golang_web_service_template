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

// CreateMenu
// @Summary ایجاد منو
// @description با ایجاد منو کاربر میتواند منو های بالای صفحه و پاین صفحه مربوط به قالب خود را کم و زیاد نماید
// @Tags menu
// @Router       /user/menu/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateMenu  	true "ورودی"
func CreateMenu(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createMenu", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateMenu(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	menu, err := models.NewMysqlManager(c).CreateMenu(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت ایجاد شد",
		"data":    menu,
	})
}

// UpdateMenu
// @Summary ویرایش منو
// @description با ایجاد منو کاربر میتواند منو های بالای صفحه و پاین صفحه مربوط به قالب خود را کم و زیاد نماید
// @Tags menu
// @Router       /user/menu/update/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	false "شناسه " SchemaExample(1)
// @Param	message	 body   DTOs.UpdateMenu  	true "ورودی"
func UpdateMenu(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:updateMenu", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.UpdateMenu(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).UpdateMenu(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت ویرایش شد",
	})
}

// DeleteMenu
// @Summary حذف منو
// @description با ایجاد منو کاربر میتواند منو های بالای صفحه و پاین صفحه مربوط به قالب خود را کم و زیاد نماید
// @Tags menu
// @Router       /user/menu/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه منو" SchemaExample(1)
func DeleteMenu(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteMenu", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	menuID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(c).DeleteMenu(menuID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت حذف شد",
	})
}

// IndexMenu
// @Summary لیست منو ها
// @description با ایجاد منو کاربر میتواند منو های بالای صفحه و پاین صفحه مربوط به قالب خود را کم و زیاد نماید
// @Tags menu
// @Router       /user/menu [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexMenu(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexMenu", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexMenu(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	menus, err := models.NewMysqlManager(c).GetAllMenuWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"menus": menus,
	})
}

// ShowMenu
// @Summary نمایش منو
// @description با ایجاد منو کاربر میتواند منو های بالای صفحه و پاین صفحه مربوط به قالب خود را کم و زیاد نماید
// @Tags post
// @Router       /user/menu/show/{id} [get]
// @Param	Authorization	header string	true "Authentication"
// @Param	id			path string	true "شناسه" SchemaExample(1)
func ShowMenu(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:ShowMenu", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	menuID := c.Param("id")
	menu, err := models.NewMysqlManager(c).FindMenuByID(utils.StringToUint64(menuID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"menu": menu,
	})
}
