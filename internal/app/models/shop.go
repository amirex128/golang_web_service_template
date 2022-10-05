package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Shop struct {
	ID            uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Name          string     `json:"name"`
	Logo          string     `json:"logo"`
	Type          string     `json:"type" sql:"type:ENUM('instagram','telegram','website')"`
	SocialAddress string     `json:"social_address"`
	VerifySocial  bool       `json:"verify"`
	SendPrice     float32    `json:"send_price"`
	Description   string     `json:"description"`
	Phone         string     `json:"phone"`
	Mobile        string     `json:"mobile"`
	TelegramID    string     `json:"telegram_id"`
	InstagramID   string     `json:"instagram_id"`
	WhatsappID    string     `json:"whatsapp_id"`
	Email         string     `json:"email"`
	Website       string     `json:"website"`
	Products      []Product  `gorm:"foreignKey:shop_id" json:"products"`
	UserID        uint64     `json:"user_id"`
	User          User       ` json:"user"`
	Categories    []Category `gorm:"many2many:category_shops;" json:"categories"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
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
		Name:          "فروشگاه امیر",
		Type:          "instagram",
		SocialAddress: "amirex_dev",
	}, 1)
}
func (m *MysqlManager) CreateShop(c *gin.Context, dto DTOs.CreateShop, userID uint64) error {
	shop := &Shop{
		Name:          dto.Name,
		Logo:          dto.LogoPath,
		Type:          dto.Type,
		SocialAddress: dto.SocialAddress,
		VerifySocial:  false,
		SendPrice:     0,
		Description:   dto.Description,
		Phone:         dto.Phone,
		Mobile:        dto.Mobile,
		TelegramID:    dto.TelegramID,
		InstagramID:   dto.InstagramID,
		WhatsappID:    dto.WhatsappID,
		Email:         dto.Email,
		Website:       dto.Website,
		UserID:        userID,
		CreatedAt:     utils.NowTime(),
		UpdatedAt:     utils.NowTime(),
	}
	err := m.GetConn().Create(shop).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطایی در ایجاد فروشگاه رخ داده است",
			"error":   err.Error(),
		})
		return err
	}
	return nil
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

func (m *MysqlManager) UpdateShop(c *gin.Context, dto DTOs.UpdateShop, shopID, userID uint64) error {
	shop := &Shop{}
	err := m.GetConn().Where("shop_id = ?", shopID).First(shop).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "فروشگاه یافت نشد",
			"error":   err.Error(),
		})
		return err
	}
	if shop.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "شما اجازه دسترسی به این فروشگاه را ندارید",
		})
		return err
	}
	if dto.Name != "" {
		shop.Name = dto.Name
	}
	if dto.LogoPath != "" {
		shop.Logo = dto.LogoPath
	}
	if dto.Type != "" {
		shop.Type = dto.Type
	}
	if dto.SocialAddress != "" {
		shop.SocialAddress = dto.SocialAddress
	}
	if dto.Description != "" {
		shop.Description = dto.Description
	}
	if dto.Phone != "" {
		shop.Phone = dto.Phone
	}
	if dto.Mobile != "" {
		shop.Mobile = dto.Mobile
	}
	if dto.TelegramID != "" {
		shop.TelegramID = dto.TelegramID
	}
	if dto.InstagramID != "" {
		shop.InstagramID = dto.InstagramID
	}
	if dto.WhatsappID != "" {
		shop.WhatsappID = dto.WhatsappID
	}
	if dto.Email != "" {
		shop.Email = dto.Email
	}
	if dto.Website != "" {
		shop.Website = dto.Website
	}
	if dto.SendPrice != 0 {
		shop.SendPrice = dto.SendPrice
	}
	shop.VerifySocial = dto.VerifySocial
	err = m.GetConn().Save(shop).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطایی در ویرایش فروشگاه رخ داده است",
			"error":   err.Error(),
		})
		return err
	}
	return nil
}

func (m *MysqlManager) DeleteShop(c *gin.Context, shopID uint64, userID uint64) error {
	shop := &Shop{}
	err := m.GetConn().Where("id = ?", shopID).First(shop).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "فروشگاه یافت نشد",
			"error":   err.Error(),
		})
		return err
	}
	if shop.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "شما اجازه دسترسی به این فروشگاه را ندارید",
		})
		return err
	}
	err = m.GetConn().Delete(shop).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطایی در حذف فروشگاه رخ داده است",
			"error":   err.Error(),
		})
		return err
	}
	return nil
}

func (m *MysqlManager) IndexShop(c *gin.Context, userID uint64) ([]*Shop, error) {
	var res []*Shop
	err := m.GetConn().Where("user_id = ?", userID).Find(&res).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطایی در دریافت فروشگاه ها رخ داده است",
			"error":   err.Error(),
		})
		return nil, err
	}
	return res, nil
}
