package menu

import (
	"github.com/amirex128/selloora_backend/internal/models"
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
