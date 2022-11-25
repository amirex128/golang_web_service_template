package gallery

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowGallery
// @Summary نمایش منو
// @description با آپلود یک تصویر میتوانید شناسه آن را در بخش های مختلف استفاده نمایید و در آینده بر اساس همین شناسه تصویر را حذف نمایید همچنین تمامی تصاویر به فرمت وب پی تبدیل میشوند
// @Tags gallery
// @Router       /user/gallery/show/{id} [get]
// @Param	Authorization	header string	true "Authentication"
// @Param	id			path string	true "شناسه" SchemaExample(1)
func ShowGallery(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:ShowGallery", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	galleryID := c.Param("id")
	gallery, err := models.NewMysqlManager(c).FindGalleryByID(utils.StringToUint64(galleryID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": gallery,
	})
}
