package menu

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

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
