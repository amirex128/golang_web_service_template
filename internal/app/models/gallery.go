package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gallery struct {
	ID       uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Path     string  `json:"path"`
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
		Path:     "https://cdn.shopify.com/s/files/1/0070/7032/9496/products/1_1024x1024.jpg?v=1588888888",
		UserID:   1,
		MimeType: "image/jpeg",
		Size:     0,
		Width:    0,
		Height:   0,
	})
}

func (m *MysqlManager) UploadImage(c *gin.Context, gallery *Gallery) (uint64, error) {
	err := m.GetConn().Create(gallery).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در آپلود تصویر",
			"error":   err.Error(),
		})
		return 0, err
	}
	return gallery.ID, nil
}
