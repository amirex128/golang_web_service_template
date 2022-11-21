package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
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
	span, ctx := apm.StartSpan(c.Request.Context(), "createMenu", "request")
	defer span.End()
	dto, err := validations.CreateMenu(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).CreateMenu(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت ایجاد شد",
	})
}

// UpdateMenu
// @Summary ویرایش منو
// @description با ایجاد منو کاربر میتواند منو های بالای صفحه و پاین صفحه مربوط به قالب خود را کم و زیاد نماید
// @Tags menu
// @Router       /user/menu/update [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.UpdateMenu  	true "ورودی"
func UpdateMenu(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updateMenu", "request")
	defer span.End()
	dto, err := validations.UpdateMenu(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).UpdateMenu(c, ctx, dto)
	if err != nil {
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
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteMenu", "request")
	defer span.End()
	menuID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(ctx).DeleteMenu(c, ctx, menuID)
	if err != nil {
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
	span, ctx := apm.StartSpan(c.Request.Context(), "indexMenu", "request")
	defer span.End()
	dto, err := validations.IndexMenu(c)
	if err != nil {
		return
	}
	menus, err := models.NewMysqlManager(ctx).GetAllMenuWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"menus": menus,
	})
}
