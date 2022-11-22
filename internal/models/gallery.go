package models

import (
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/spf13/viper"
	"go.elastic.co/apm/v2"
)

type Gallery struct {
	ID       uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Path     string  `json:"path"`
	FullPath string  `json:"full_path"`
	UserID   uint64  `gorm:"default:null" json:"user_id"`
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

func (m *MysqlManager) UploadImage(gallery *Gallery) (uint64, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:UploadImage", "model")
	defer span.End()
	gallery.FullPath = viper.GetString("server_url") + gallery.Path
	err := m.GetConn().Create(gallery).Error
	if err != nil {
		return 0, errorx.New("خطا در آپلود تصویر", "model", err)
	}
	return gallery.ID, nil
}

func (m *MysqlManager) DeleteGallery(galleryID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:DeleteGallery", "model")
	defer span.End()
	var err error
	err = m.GetConn().Table("shops").Where("gallery_id=?", galleryID).Update("gallery_id", nil).Error
	err = m.GetConn().Table("users").Where("gallery_id=?", galleryID).Update("gallery_id", nil).Error
	err = m.GetConn().Table("posts").Where("gallery_id=?", galleryID).Update("gallery_id", nil).Error
	err = m.GetConn().Table("tickets").Where("gallery_id=?", galleryID).Update("gallery_id", nil).Error
	if err != nil {
		return errorx.New("خطا در حذف تصویر", "model", err)
	}
	err = m.GetConn().Delete(&Gallery{}, galleryID).Error
	if err != nil {
		return errorx.New("خطا در حذف تصویر", "model", err)
	}

	return nil
}

func (m *MysqlManager) FindGalleryByID(galleryID uint64) (*Gallery, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindGalleryByID", "model")
	defer span.End()
	var gallery Gallery
	err := m.GetConn().Where("id = ?", galleryID).First(&gallery).Error
	if err != nil {
		return nil, errorx.New("خطا در حذف تصویر", "model", err)
	}
	userID := GetUser(m.Ctx)
	if gallery.UserID != *userID {
		return nil, errorx.New("شما اجازه حذف این تصویر را ندارید", "model", err)
	}
	return &gallery, nil
}
