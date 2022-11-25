package slider

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// DeleteSlider
// @Summary حذف اسلایدر
// @description هر فروشگاه برای خود میتواند به تعداد دلخواه اسلایدر در موقعیت های مختلف مثل بالای صفحه و پایین صفحه ایجاد نماید
// @Tags slider
// @Router       /user/slider/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه اسلایدر" SchemaExample(1)
func DeleteSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteSlider", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	sliderID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(c).DeleteSlider(sliderID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اسلایدر با موفقیت حذف شد",
	})
}
