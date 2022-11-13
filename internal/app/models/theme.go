package models

import (
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
