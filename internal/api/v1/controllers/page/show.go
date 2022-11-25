package page

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowPage
// @Summary نمایش صفحه
// @description هر فروشگاه میتواند به تعداد دلخواه صفحه ایجاد کند صفحات از دو حالت معمولی و خالی تشکیل میشوند که در حالت معمولی کاربر میتواند کل یک اچ تی ام ال را ذخیره نماید تا بدون چارچوب های قالب نمایش داده شود و در حالت معمولی همراه با چهار چوب ها نمایش داده شود
// @Tags page
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
