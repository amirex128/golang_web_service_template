package theme

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// IndexTheme
// @Summary لیست قالب ها
// @description کاربر در هنگام ایجاد فروشگاه باید قالب فروشگاه خود را انتخاب نماید یا این که در آینده بتواند قالب خود را تغییر دهد
// @Tags theme
// @Router       /user/theme [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexTheme(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexTheme", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexTheme(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	pages, err := models.NewMysqlManager(c).GetAllThemeWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pages": pages,
	})
}
