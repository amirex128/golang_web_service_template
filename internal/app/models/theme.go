package models

import (
	"github.com/amirex128/selloora_backend/internal/app/DTOs"
	"context"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
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
	manager.GetConn().AutoMigrate(&Theme{})
	for i := 0; i < 3; i++ {
		var id uint64 = 1
		manager.CreateTheme(&gin.Context{}, context.Background(), Theme{
			Name:        "قالب شماره" + strconv.Itoa(i),
			Description: "قالب شماره" + strconv.Itoa(i),
			GalleryID:   &id,
		})
	}
}

func (m *MysqlManager) CreateTheme(c *gin.Context, ctx context.Context, theme Theme) {
	err := m.GetConn().Create(&theme).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطایی در سرور رخ داده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return
	}
}

func (m *MysqlManager) FindThemeByID(c *gin.Context, ctx context.Context, themeID uint64) (*Theme, error) {
	span, ctx := apm.StartSpan(ctx, "FindThemeByID", "model")
	defer span.End()
	theme := &Theme{}
	err := m.GetConn().Where("id = ?", themeID).First(theme).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در پیدا کردن قالب رخ داده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return nil, err
	}
	return theme, nil
}
func (m *MysqlManager) DeleteTheme(c *gin.Context, ctx context.Context, themeID uint64) error {
	span, ctx := apm.StartSpan(ctx, "showTheme", "model")
	defer span.End()
	theme := Theme{}
	err := m.GetConn().Where("id = ?", themeID).First(&theme).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "قالب یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}

	err = m.GetConn().Delete(&theme).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در حذف قالب",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}
func (m *MysqlManager) GetAllThemeWithPagination(c *gin.Context, ctx context.Context, dto DTOs.IndexTheme) (*DTOs.Pagination, error) {
	span, ctx := apm.StartSpan(ctx, "showTheme", "model")
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در دریافت قالب ها",
			"error":   err.Error(),
			"type":    "model",
		})
		return pagination, err
	}
	pagination.Data = themes
	return pagination, nil
}
