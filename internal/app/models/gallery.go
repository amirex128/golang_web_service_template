package models

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.elastic.co/apm/v2"
	"net/http"
)

type Gallery struct {
	ID       uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Path     string  `json:"path"`
	FullPath string  `json:"full_path"`
	UserID   uint64  `json:"user_id"`
	MimeType string  `json:"mime_type"`
	Size     float64 `json:"size"`
	Width    uint32  `json:"width"`
	Height   uint32  `json:"height"`
}

func initGallery(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Gallery{})
	manager.GetConn().Create(&Gallery{
		ID:       1,
		Path:     "https://montiego.ir/wp-content/uploads/2020/03/75af01aff54b42a.jpg",
		FullPath: "https://montiego.ir/wp-content/uploads/2020/03/75af01aff54b42a.jpg",
		UserID:   1,
		MimeType: "image/jpeg",
		Size:     0,
		Width:    0,
		Height:   0,
	})
}

func (m *MysqlManager) UploadImage(c *gin.Context, ctx context.Context, gallery *Gallery) (uint64, error) {
	span, ctx := apm.StartSpan(ctx, "UploadImage", "model")
	defer span.End()
	gallery.FullPath = viper.GetString("server_url") + gallery.Path
	err := m.GetConn().Create(gallery).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در آپلود تصویر",
			"error":   err.Error(),
			"type":    "model",
		})
		return 0, err
	}
	return gallery.ID, nil
}

func (m *MysqlManager) DeleteGallery(c *gin.Context, ctx context.Context, galleryID uint64) error {
	span, ctx := apm.StartSpan(ctx, "DeleteGallery", "model")
	defer span.End()
	var err error
	err = m.GetConn().Table("shops").Where("gallery_id=?", galleryID).Update("gallery_id", nil).Error
	err = m.GetConn().Table("users").Where("gallery_id=?", galleryID).Update("gallery_id", nil).Error
	err = m.GetConn().Table("posts").Where("gallery_id=?", galleryID).Update("gallery_id", nil).Error
	err = m.GetConn().Table("tickets").Where("gallery_id=?", galleryID).Update("gallery_id", nil).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در حذف تصویر",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	err = m.GetConn().Delete(&Gallery{}, galleryID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در حذف تصویر",
			"error":   err.Error(),
		})
		return err
	}

	return nil
}

func (m *MysqlManager) FindGalleryByID(c *gin.Context, ctx context.Context, galleryID uint64, userID uint64) (*Gallery, error) {
	span, ctx := apm.StartSpan(ctx, "FindGalleryByID", "model")
	defer span.End()
	var gallery Gallery
	err := m.GetConn().Where("id = ?", galleryID).First(&gallery).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در حذف تصویر",
			"error":   err.Error(),
			"type":    "model",
		})
		return nil, err
	}
	if gallery.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "شما اجازه حذف این تصویر را ندارید",
			"error":   errors.New("شما اجازه حذف این تصویر را ندارید"),
			"type":    "model",
		})
	}
	return &gallery, nil
}
