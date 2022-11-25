package gallery

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
	"os"
	"path/filepath"
)

// DeleteGallery
// @Summary حذف گالری
// @description حذف یک تصویر از گالری
// @Tags gallery
// @Router       /user/gallery/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه گالری" SchemaExample(1)
func DeleteGallery(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteGallery", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	galleryID := c.Param("id")
	gallery, err := models.NewMysqlManager(c).FindGalleryByID(utils.StringToUint64(galleryID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	abs, _ := filepath.Abs("./")
	path := filepath.Join(abs, gallery.Path)
	if err := os.Remove(path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در حذف تصویر",
			"error":   err.Error(),
		})
		return
	}
	err = models.NewMysqlManager(c).DeleteGallery(utils.StringToUint64(galleryID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تصویر با موفقیت حذف شد",
	})
}
