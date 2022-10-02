package models

import (
	"backend/internal/app/DTOs"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Shop struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `json:"name"`
	Logo        string    `json:"logo"`
	Type        string    `json:"type" sql:"type:ENUM('instagram','telegram','website')"`
	Social      string    `json:"social"`
	Verify      bool      `json:"verify"`
	SendPrice   float32   `json:"send_price"`
	GuildID     uint32    `json:"guild_id"`
	Description string    `json:"description"`
	Phone       string    `json:"phone"`
	Mobile      string    `json:"mobile"`
	TelegramID  string    `json:"telegram_id"`
	InstagramID string    `json:"instagram_id"`
	WhatsappID  string    `json:"whatsapp_id"`
	Email       string    `json:"email"`
	Website     string    `json:"website"`
	Products    []Product `gorm:"foreignKey:shop_id" json:"products"`
	UserID      uint64    `json:"user_id"`
	User        User      `gorm:"foreignKey:user_id" json:"user"`
	CategoryID  uint64    `json:"category_id"`
	Category    Category  `gorm:"foreignKey:category_id" json:"category"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}
type ShopArr []Shop

func (s ShopArr) Len() int {
	return len(s)
}
func (s ShopArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ShopArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Shop) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Shop) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initShop(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Shop{})
	manager.CreateShop(&gin.Context{}, DTOs.CreateShop{
		Name:       "فروشگاه امیر",
		Type:       "instagram",
		Social:     "amirex_dev",
		CategoryID: 1,
	}, 1)
}
func (m *MysqlManager) CreateShop(c *gin.Context, dto DTOs.CreateShop, userID uint64) (*Shop, error) {
	shop := &Shop{
		Name:       dto.Name,
		Type:       dto.Type,
		Social:     dto.Social,
		CategoryID: dto.CategoryID,
		UserID:     userID,
	}
	err := m.GetConn().Create(shop).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطایی در ایجاد فروشگاه رخ داده است",
			"error":   err.Error(),
		})
		return nil, err
	}
	return shop, nil
}

func (m *MysqlManager) FindShopByID(c *gin.Context, shopID uint64) (*Shop, error) {
	res := &Shop{}
	err := m.GetConn().Where("id = ?", shopID).First(res).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "فروشگاه یافت نشد",
			"error":   err.Error(),
		})
		return nil, err
	}
	return res, nil
}
