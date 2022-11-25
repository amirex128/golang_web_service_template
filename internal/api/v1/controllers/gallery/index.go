package gallery

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// IndexGallery
// @Summary لیست منو ها
// @description با آپلود یک تصویر میتوانید شناسه آن را در بخش های مختلف استفاده نمایید و در آینده بر اساس همین شناسه تصویر را حذف نمایید همچنین تمامی تصاویر به فرمت وب پی تبدیل میشوند
// @Tags gallery
// @Router       /user/gallery [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexGallery(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexGallery", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexGallery(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	galleries, err := models.NewMysqlManager(c).GetAllGalleryWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": galleries,
	})
}
