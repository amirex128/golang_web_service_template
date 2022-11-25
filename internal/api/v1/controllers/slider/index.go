package slider

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// IndexSlider
// @Summary لیست اسلایدر ها
// @description هر فروشگاه برای خود میتواند به تعداد دلخواه اسلایدر در موقعیت های مختلف مثل بالای صفحه و پایین صفحه ایجاد نماید
// @Tags slider
// @Router       /user/slider [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexSlider", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexSlider(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	sliders, err := models.NewMysqlManager(c).GetAllSliderWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sliders": sliders,
	})
}
