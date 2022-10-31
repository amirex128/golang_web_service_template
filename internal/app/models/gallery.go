package models

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Gallery struct {
	ID        uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Path      string  `json:"path"`
	UserID    uint64  `json:"user_id"`
	OwnerID   uint64  `json:"owner_id"`
	OwnerType string  `json:"owner_type "`
	MimeType  string  `json:"mime_type"`
	Size      float64 `json:"size"`
	Width     uint32  `json:"width"`
	Height    uint32  `json:"height"`
	Sort      uint    `json:"sort"`
}
type GalleryArr []Gallery

func (s GalleryArr) Len() int {
	return len(s)
}
func (s GalleryArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s GalleryArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Gallery) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Gallery) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initGallery(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Gallery{})
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
