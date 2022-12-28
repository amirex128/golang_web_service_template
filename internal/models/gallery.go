package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/spf13/viper"
	"go.elastic.co/apm/v2"
)

type Gallery struct {
	ID       uint64  `gorm:"primary_key;auto_increment" json:"id" fake:"{custom_uint64:0}"`
	Path     string  `json:"path" fake:"{imageurl}"`
	FullPath string  `json:"full_path" fake:"{imageurl}"`
	UserID   *uint64 `gorm:"default:null" json:"user_id" fake:"{number:1,100}"`
	MimeType string  `json:"mime_type" fake:"{custom_string:image/webp}"`
	Size     float64 `json:"size" fake:"{number:1000,2000}"`
	Width    uint32  `json:"width" fake:"{number:100,200}"`
	Height   uint32  `json:"height" fake:"{number:100,200}"`
}

func initGallery(manager *MysqlManager) {

	if !manager.GetConn().Migrator().HasTable(&Gallery{}) {
		manager.GetConn().Migrator().CreateTable(&Gallery{})

		for i := 0; i < 100; i++ {
			model := new(Gallery)
			gofakeit.Struct(model)

			manager.GetConn().Create(model)
		}
	}

}

func (m *MysqlManager) UploadImage(gallery *Gallery) (*Gallery, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:UploadImage", "model")
	defer span.End()
	gallery.FullPath = viper.GetString("server_url") + gallery.Path
	gallery.UserID = utils.GetUserID(m.Ctx)
	err := m.GetConn().Create(gallery).Error
	if err != nil {
		return nil, errorx.New("خطا در آپلود تصویر", "model", err)
	}
	return gallery, nil
}

func (m *MysqlManager) DeleteGallery(galleryID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:DeleteGallery", "model")
	defer span.End()
	var err error

	conn := m.GetConn()
	tx := conn.Begin()
	err = tx.Table("shops").Where("gallery_id=?", galleryID).Update("gallery_id", nil).Error
	err = tx.Table("users").Where("gallery_id=?", galleryID).Update("gallery_id", nil).Error
	err = tx.Table("posts").Where("gallery_id=?", galleryID).Update("gallery_id", nil).Error
	err = tx.Table("tickets").Where("gallery_id=?", galleryID).Update("gallery_id", nil).Error
	if err != nil {
		tx.Rollback()
		return errorx.New("خطا در حذف تصویر", "model", err)
	}
	gallery := &Gallery{}
	err = tx.Where("id = ?", galleryID).First(gallery).Error
	if err != nil {
		tx.Rollback()
		return errorx.New("خطا در حذف تصویر", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, gallery.UserID); err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Delete(gallery).Error
	if err != nil {
		tx.Rollback()
		return errorx.New("خطا در حذف تصویر", "model", err)
	}

	tx.Commit()
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
	return &gallery, nil
}
func (m *MysqlManager) GetAllGalleryWithPagination(dto DTOs.IndexGallery) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showGallery", "model")
	defer span.End()
	conn := m.GetConn()
	var galleries []Gallery
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	userID := utils.GetUserID(m.Ctx)
	conn = conn.Scopes(DTOs.Paginate("galleries", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("user_id = ? ", userID)
	}
	err := conn.Order("id DESC").Find(&galleries).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت تصاویر", "model", err)
	}
	pagination.Data = galleries
	return pagination, nil
}
