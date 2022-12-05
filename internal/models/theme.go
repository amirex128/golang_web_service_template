package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"go.elastic.co/apm/v2"
	"strconv"
)

type Theme struct {
	ID          uint64   `gorm:"primary_key;auto_increment" json:"id"`
	Name        string   `json:"name"`
	GalleryID   *uint64  `gorm:"default:null" json:"gallery_id"`
	Gallery     *Gallery `json:"gallery"`
	Description string   `json:"description"`
	Type        string   `gorm:"default:website" json:"type" sql:"type:ENUM('instagram','website')"`
}

func initTheme(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&Theme{}) {
		manager.GetConn().AutoMigrate(&Theme{})
		for i := 0; i < 4; i++ {
			var id uint64 = 1
			manager.CreateTheme(Theme{
				Name:        "قالب شماره" + strconv.Itoa(i),
				Description: "قالب شماره" + strconv.Itoa(i),
				GalleryID:   &id,
			})
		}
	}

}

func (m *MysqlManager) CreateTheme(theme Theme) error {
	err := m.GetConn().Create(&theme).Error
	if err != nil {
		return errorx.New("خطایی در سرور رخ داده است", "model", err)
	}
	return nil
}

func (m *MysqlManager) FindThemeByID(themeID uint64) (*Theme, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindThemeByID", "model")
	defer span.End()
	theme := &Theme{}
	err := m.GetConn().Where("id = ?", themeID).First(theme).Error
	if err != nil {
		return nil, errorx.New("خطایی در سرور رخ داده است", "model", err)
	}
	return theme, nil
}
func (m *MysqlManager) DeleteTheme(themeID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showTheme", "model")
	defer span.End()
	theme := Theme{}
	err := m.GetConn().Where("id = ?", themeID).First(&theme).Error
	if err != nil {
		return errorx.New("قالب یافت نشد", "model", err)
	}

	err = m.GetConn().Delete(&theme).Error
	if err != nil {
		return errorx.New("خطا در حذف قالب", "model", err)
	}
	return nil
}
func (m *MysqlManager) GetAllThemeWithPagination(dto DTOs.IndexTheme) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showTheme", "model")
	defer span.End()
	conn := m.GetConn()
	var themes []Theme
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("themes", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%").Order("id DESC")
	}
	err := conn.Find(&themes).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت قالب ها", "model", err)
	}
	pagination.Data = themes
	return pagination, nil
}
