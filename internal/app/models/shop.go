package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"net/http"
)

type Shop struct {
	ID            uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Name          string     `json:"name"`
	EnglishName   string     `json:"english_name"`
	Type          string     `json:"type" sql:"type:ENUM('instagram','telegram','website')"`
	SocialAddress string     `json:"social_address"`
	VerifySocial  bool       `json:"verify_social"`
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
	User          User       `gorm:"foreignKey:user_id" json:"user"`
	Categories    []Category `gorm:"many2many:category_shops;" json:"categories"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
	GalleryID     *uint64    `gorm:"default:null" json:"gallery_id"`
	Gallery       *Gallery   `gorm:"foreignKey:gallery_id" json:"gallery"`
}

func initShop(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Shop{})
	for i := 0; i < 20; i++ {
		manager.CreateShop(&gin.Context{}, DTOs.CreateShop{
			Name:          "فروشگاه امیر",
			Type:          "instagram",
			EnglishName:   "instagram",
			SocialAddress: "amirex_dev",
			GalleryID:     1,
		}, 1)
	}
}
func (m *MysqlManager) CreateShop(c *gin.Context, dto DTOs.CreateShop, userID uint64) error {
	shop := &Shop{
		Name:          dto.Name,
		EnglishName:   slug.MakeLang(dto.EnglishName, "en"),
		Type:          dto.Type,
		SocialAddress: dto.SocialAddress,
		VerifySocial:  false,
		SendPrice:     dto.SendPrice,
		Description:   dto.Description,
		Phone:         dto.Phone,
		Mobile:        dto.Mobile,
		TelegramID:    dto.TelegramID,
		InstagramID:   dto.InstagramID,
		WhatsappID:    dto.WhatsappID,
		Email:         dto.Email,
		Website:       dto.Website,
		UserID:        userID,
		GalleryID: func() *uint64 {
			if dto.GalleryID == 0 {
				return nil
			}
			return &dto.GalleryID
		}(),
		CreatedAt: utils.NowTime(),
		UpdatedAt: utils.NowTime(),
	}
	err := m.GetConn().Create(shop).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطایی در ایجاد فروشگاه رخ داده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) FindShopByID(c *gin.Context, shopID uint64, userID uint64) (*Shop, error) {
	res := &Shop{}
	err := m.GetConn().Where("id = ?", shopID).Preload("Gallery").First(res).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "فروشگاه یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
		return nil, err
	}
	if res.UserID != userID && userID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "شما اجازه دسترسی به این فروشگاه را ندارید",
			"error":   errors.New("شما اجازه دسترسی به این فروشگاه را ندارید"),
			"type":    "permission",
		})
		return nil, err
	}
	return res, nil
}

func (m *MysqlManager) UpdateShop(c *gin.Context, dto DTOs.UpdateShop, shopID, userID uint64) error {
	shop := &Shop{}
	err := m.GetConn().Where("id = ?", shopID).First(shop).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "فروشگاه یافت نشد",
			"error":   err.Error(),
			"type":    "model",
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
	if dto.EnglishName != "" {
		shop.EnglishName = slug.MakeLang(dto.EnglishName, "en")
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
	if dto.GalleryID != 0 {
		shop.GalleryID = &dto.GalleryID
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
			"type":    "model",
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
			"type":    "model",
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
			"type":    "model",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) GetAllShopWithPagination(c *gin.Context, dto DTOs.IndexShop) (*DTOs.Pagination, error) {
	conn := m.GetConn()
	var shops []Shop
	userID := GetUser(c)
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("shops", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%").Order("id DESC")
	}
	err := conn.Where("user_id = ?", userID).Preload("Gallery").Find(&shops).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در دریافت فروشگاه ها",
			"error":   err.Error(),
			"type":    "model",
		})
		return pagination, err
	}
	pagination.Data = shops
	return pagination, nil
}

func (m *MysqlManager) GetAllShop(c *gin.Context) ([]Shop, error) {
	userID := GetUser(c)
	var shops []Shop
	err := m.GetConn().Where("user_id = ?", userID).Preload("Gallery").Find(&shops).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در دریافت فروشگاه ها",
			"error":   err.Error(),
			"type":    "model",
		})
		return shops, err
	}
	return shops, nil
}
