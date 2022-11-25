package page

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

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
