package slider

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowSlider
// @Summary نمایش اسلایدر
// @description هر فروشگاه برای خود میتواند به تعداد دلخواه اسلایدر در موقعیت های مختلف مثل بالای صفحه و پایین صفحه ایجاد نماید
// @Tags slider
// @Router       /user/slider/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه " SchemaExample(1)
func ShowSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showSlider", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	sliderID := c.Param("id")
	slider, err := models.NewMysqlManager(c).FindSliderByID(utils.StringToUint64(sliderID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"slider": slider,
	})
}
