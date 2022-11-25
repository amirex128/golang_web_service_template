package slider

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// UpdateSlider
// @Summary ویرایش اسلایدر
// @description هر فروشگاه برای خود میتواند به تعداد دلخواه اسلایدر در موقعیت های مختلف مثل بالای صفحه و پایین صفحه ایجاد نماید
// @Tags slider
// @Router       /user/slider/update/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	false "شناسه " SchemaExample(1)
// @Param	message	 body   DTOs.UpdateSlider  	true "ورودی"
func UpdateSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:updateSlider", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.UpdateSlider(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).UpdateSlider(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اسلایدر با موفقیت ویرایش شد",
	})
}
