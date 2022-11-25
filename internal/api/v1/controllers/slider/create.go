package slider

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreateSlider
// @Summary ایجاد اسلایدر
// @description هر فروشگاه برای خود میتواند به تعداد دلخواه اسلایدر در موقعیت های مختلف مثل بالای صفحه و پایین صفحه ایجاد نماید
// @Tags slider
// @Router       /user/slider/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateSlider  	true "ورودی"
func CreateSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createSlider", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateSlider(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	slider, err := models.NewMysqlManager(c).CreateSlider(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اسلایدر با موفقیت ایجاد شد",
		"data":    slider,
	})
}
