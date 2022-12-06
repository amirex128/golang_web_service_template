package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/brianvoe/gofakeit/v6"
	"go.elastic.co/apm/v2"
)

type Slider struct {
	ID          uint64  `gorm:"primary_key;auto_increment" json:"id"`
	UserID      *uint64 `json:"user_id"`
	ShopID      *uint64 `json:"shop_id"`
	Link        string  `json:"link"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	GalleryID   *uint64 `json:"gallery_id"`
	Gallery     Gallery `json:"gallery"`
	Sort        uint32  `json:"sort"`
	Position    string  `gorm:"default:top" json:"position" sql:"type:ENUM('top','bottom','left','right')"`
}

func initSlider(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&Slider{}) {
		manager.GetConn().Migrator().CreateTable(&Slider{})
		for i := 0; i < 100; i++ {
			model := new(DTOs.CreateSlider)
			gofakeit.Struct(model)

			manager.CreateSlider(*model)
		}

	}

}

func (m *MysqlManager) CreateSlider(dto DTOs.CreateSlider) (*Slider, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showSlider", "model")
	defer span.End()

	// find last sort number
	var lastSort Slider
	err := m.GetConn().Order("sort desc").First(&lastSort).Error
	if err != nil {
		lastSort.Sort = 0
	}

	slider := &Slider{
		Title:       dto.Title,
		GalleryID:   &dto.GalleryID,
		Description: dto.Description,
		Link:        dto.Link,
		UserID:      GetUserID(m.Ctx),
		ShopID:      &dto.ShopID,
		Position:    dto.Position,
		Sort:        lastSort.Sort + 1,
	}
	err = m.GetConn().Create(slider).Error
	if err != nil {
		return slider, errorx.New("خطا در ایجاد کد اسلایدر", "model", err)
	}
	return slider, nil
}
func (m *MysqlManager) UpdateSlider(dto DTOs.UpdateSlider) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showSlider", "model")
	defer span.End()
	slider := &Slider{}
	err := m.GetConn().Where("id = ?", dto.ID).First(slider).Error
	if err != nil {
		return errorx.New("اسلایدر یافت نشد", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, slider.UserID); err != nil {
		return err
	}
	if dto.Title != "" {
		slider.Title = dto.Title
	}
	if dto.Description != "" {
		slider.Description = dto.Description
	}
	if dto.GalleryID != 0 {
		slider.GalleryID = &dto.GalleryID
	}
	if dto.Link != "" {
		slider.Link = dto.Link
	}
	if dto.Position != "" {
		slider.Position = dto.Position
	}
	if dto.Sort != 0 {
		slider.Sort = dto.Sort
	}
	err = m.GetConn().Save(slider).Error
	if err != nil {
		return errorx.New("خطا در ویرایش اسلایدر", "model", err)
	}
	return nil
}
func (m *MysqlManager) DeleteSlider(sliderID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showSlider", "model")
	defer span.End()
	slider := Slider{}
	err := m.GetConn().Where("id = ?", sliderID).First(&slider).Error
	if err != nil {
		return errorx.New("اسلایدر یافت نشد", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, slider.UserID); err != nil {
		return err
	}
	err = m.GetConn().Delete(&slider).Error
	if err != nil {
		return errorx.New("خطا در حذف اسلایدر", "model", err)
	}
	return nil
}
func (m *MysqlManager) GetAllSliderWithPagination(dto DTOs.IndexSlider) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showSlider", "model")
	defer span.End()
	conn := m.GetConn()
	var sliders []Slider
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("sliders", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("title LIKE ?", "%"+dto.Search+"%")
	}
	err := conn.Where("shop_id = ? ", dto.ShopID).Order("id DESC").Find(&sliders).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت اسلایدر ها", "model", err)
	}
	pagination.Data = sliders
	return pagination, nil
}

func (m *MysqlManager) FindSliderByID(sliderID uint64) (*Slider, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindSliderByID", "model")
	defer span.End()
	slider := &Slider{}
	err := m.GetConn().Where("id = ?", sliderID).First(slider).Error
	if err != nil {
		return nil, errorx.New("مشکلی در یافتن اسلایدر پیش آمده است", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, slider.UserID); err != nil {
		return nil, err
	}
	return slider, nil
}
