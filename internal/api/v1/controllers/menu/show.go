package menu

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowMenu
// @Summary نمایش منو
// @description با ایجاد منو کاربر میتواند منو های بالای صفحه و پاین صفحه مربوط به قالب خود را کم و زیاد نماید
// @Tags menu
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
