package page

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreatePage
// @Summary ایجاد صفحه
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
	page, err := models.NewMysqlManager(c).CreatePage(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "صفحه با موفقیت ایجاد شد",
		"data":    page,
	})
}
