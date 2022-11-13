package models

import (
	"backend/internal/app/DTOs"
	"context"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

type Slider struct {
	ID          uint64  `gorm:"primary_key;auto_increment" json:"id"`
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
	manager.GetConn().AutoMigrate(&Slider{})
	for i := 0; i < 20; i++ {
		manager.CreateSlider(&gin.Context{}, context.Background(), DTOs.CreateSlider{
			Title:       "ِسشیبش",
			GalleryID:   1,
			Description: "fdfsdf",
			Link:        "DSADSA",
			ShopID:      1,
			Position:    "top",
		})
	}
}

func (m *MysqlManager) CreateSlider(c *gin.Context, ctx context.Context, dto DTOs.CreateSlider) error {
	span, ctx := apm.StartSpan(ctx, "showSlider", "model")
	defer span.End()

	// find last sort number
	var lastSort *Slider
	err := m.GetConn().Order("sort desc").First(lastSort).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در ایجاد اسلایدر",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}

	slider := &Slider{
		Title:       dto.Title,
		GalleryID:   &dto.GalleryID,
		Description: dto.Description,
		Link:        dto.Link,
		ShopID:      &dto.ShopID,
		Position:    dto.Position,
		Sort: func() uint32 {
			if lastSort != nil {
				return lastSort.Sort + 1
			}
			return 1
		}(),
	}
	err = m.GetConn().Create(&slider).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در ایجاد کد اسلایدر",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}
func (m *MysqlManager) UpdateSlider(c *gin.Context, ctx context.Context, dto DTOs.UpdateSlider) error {
	span, ctx := apm.StartSpan(ctx, "showSlider", "model")
	defer span.End()
	slider := &Slider{}
	err := m.GetConn().Where("id = ?", dto.ID).First(slider).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "اسلایدر یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در ویرایش اسلایدر"})
		return err
	}
	return nil
}
func (m *MysqlManager) DeleteSlider(c *gin.Context, ctx context.Context, sliderID uint64) error {
	span, ctx := apm.StartSpan(ctx, "showSlider", "model")
	defer span.End()
	slider := Slider{}
	err := m.GetConn().Where("id = ?", sliderID).First(&slider).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "اسلایدر یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}

	err = m.GetConn().Delete(&slider).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در حذف اسلایدر",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}
func (m *MysqlManager) GetAllSliderWithPagination(c *gin.Context, ctx context.Context, dto DTOs.IndexSlider) (*DTOs.Pagination, error) {
	span, ctx := apm.StartSpan(ctx, "showSlider", "model")
	defer span.End()
	conn := m.GetConn()
	var sliders []Slider
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("sliders", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("title LIKE ?", "%"+dto.Search+"%").Where("shop_id = ? ", dto.ShopID).Order("id DESC")
	}
	err := conn.Find(&sliders).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در دریافت اسلایدر ها",
			"error":   err.Error(),
			"type":    "model",
		})
		return pagination, err
	}
	pagination.Data = sliders
	return pagination, nil
}
