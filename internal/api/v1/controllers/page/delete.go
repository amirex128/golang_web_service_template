package page

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

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
