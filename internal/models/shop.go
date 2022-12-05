package models

import (
	"fmt"
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/brianvoe/gofakeit/v6"
	"go.elastic.co/apm/v2"
	"os"
)

type Shop struct {
	ID            uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Name          string     `json:"name"`
	Type          string     `json:"type" sql:"type:ENUM('instagram','telegram','website')"`
	SocialAddress string     `json:"social_address"`
	SendPrice     float32    `json:"send_price"`
	Description   string     `json:"description"`
	Phone         string     `json:"phone"`
	Mobile        string     `json:"mobile"`
	TelegramID    string     `json:"telegram_id"`
	InstagramID   string     `json:"instagram_id"`
	WhatsappID    string     `json:"whatsapp_id"`
	Email         string     `json:"email"`
	Website       string     `json:"website"`
	ThemeID       *uint64    `gorm:"default:null" json:"theme_id"`
	Theme         *Theme     `gorm:"foreignKey:ThemeID" json:"theme"`
	Products      []Product  `gorm:"foreignKey:shop_id" json:"products"`
	UserID        *uint64    `json:"user_id"`
	User          User       `gorm:"foreignKey:user_id" json:"user"`
	Categories    []Category `gorm:"many2many:category_shops;" json:"categories"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
	GalleryID     *uint64    `gorm:"default:null" json:"gallery_id"`
	Gallery       *Gallery   `gorm:"foreignKey:gallery_id" json:"gallery"`
}

func initShop(manager *MysqlManager) {

	if !manager.GetConn().Migrator().HasTable(&Shop{}) {

		manager.GetConn().AutoMigrate(&Shop{})
		for i := 0; i < 100; i++ {
			model := new(DTOs.CreateShop)
			gofakeit.Struct(model)

			manager.CreateShop(*model)
		}
	}
}

func (m *MysqlManager) CreateShop(dto DTOs.CreateShop) (*Shop, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreateShop", "model")
	defer span.End()
	userID := GetUser(m.Ctx)
	shop := &Shop{
		Name:          dto.Name,
		Type:          dto.Type,
		SocialAddress: dto.SocialAddress,
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
		ThemeID:   nil,
		CreatedAt: utils.NowTime(),
		UpdatedAt: utils.NowTime(),
	}
	err := m.GetConn().Create(shop).Error
	if err != nil {
		return shop, errorx.New("خطایی در ایجاد فروشگاه رخ داده است", "model", err)
	}
	return shop, nil
}

func (m *MysqlManager) FindShopByID(shopID uint64) (*Shop, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindShopByID", "model")
	defer span.End()
	res := &Shop{}
	err := m.GetConn().Where("id = ?", shopID).Preload("Gallery").First(res).Error
	if err != nil {
		return nil, errorx.New("فروشگاه یافت نشد", "model", err)
	}
	return res, nil
}

func (m *MysqlManager) UpdateShop(dto DTOs.UpdateShop) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:UpdateShop", "model")
	defer span.End()
	shop := &Shop{}
	err := m.GetConn().Where("id = ?", dto.ID).First(shop).Error
	if err != nil {
		return errorx.New("فروشگاه یافت نشد", "model", err)
	}
	userID := GetUser(m.Ctx)
	if *shop.UserID != *userID {
		return errorx.New("شما اجازه دسترسی به این فروشگاه را ندارید", "model", err)
	}
	if dto.Name != "" {
		shop.Name = dto.Name
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
	err = m.GetConn().Save(shop).Error
	if err != nil {
		return errorx.New("خطایی در ویرایش فروشگاه رخ داده است", "model", err)
	}
	return nil
}

func (m *MysqlManager) DeleteShop(shopID uint64, userID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:DeleteShop", "model")
	defer span.End()
	shop := &Shop{}
	err := m.GetConn().Where("id = ?", shopID).First(shop).Error
	if err != nil {
		return errorx.New("فروشگاه یافت نشد", "model", err)
	}
	if *shop.UserID != userID {
		return errorx.New("شما اجازه دسترسی به این فروشگاه را ندارید", "model", err)
	}
	err = m.GetConn().Delete(shop).Error
	if err != nil {
		return errorx.New("خطایی در حذف فروشگاه رخ داده است", "model", err)
	}
	return nil
}

func (m *MysqlManager) GetAllShopWithPagination(dto DTOs.IndexShop) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetAllShopWithPagination", "model")
	defer span.End()
	conn := m.GetConn()
	var shops []Shop
	userID := GetUser(m.Ctx)
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("shops", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%").Order("id DESC")
	}
	err := conn.Where("user_id = ?", userID).Preload("Gallery").Find(&shops).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت فروشگاه ها", "model", err)
	}
	pagination.Data = shops
	return pagination, nil
}

func (m *MysqlManager) FindShopByDomain(name string) (*Shop, *Domain, *Theme, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindShopByDomain", "model")
	defer span.End()
	domain, err := m.FindDomainByName(name)
	if err != nil {
		return nil, nil, nil, err
	}
	shop, err := m.FindShopByID(*domain.ShopID)
	if err != nil {
		return nil, nil, nil, err
	}
	theme, err := m.FindThemeByID(*shop.ThemeID)
	if err != nil {
		return nil, nil, nil, err
	}
	return shop, domain, theme, nil
}

func (m *MysqlManager) SelectThemeByID(themeID uint64, shopID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:SelectThemeByID", "model")
	defer span.End()
	shop, err := m.FindShopByID(shopID)
	if err != nil {
		return err
	}
	userID := GetUser(m.Ctx)
	if *shop.UserID != *userID {
		return errorx.New("شما اجازه دسترسی به این فروشگاه را ندارید", "model", err)
	}
	shop.ThemeID = &themeID
	err = m.GetConn().Save(shop).Error
	if err != nil {
		return errorx.New("خطایی در انتخاب قالب رخ داده است", "model", err)
	}
	pages, err := m.FindPageByShopID(shopID)
	if err != nil {
		return err
	}
	m.SetPageForTheme(pages, shopID, themeID)
	return nil
}

func (m *MysqlManager) SetPageForTheme(pages []*Page, shopID uint64, themeID uint64) error {

	tID := fmt.Sprintf("%d", themeID)
	files, err := os.ReadDir("./csv/themes/" + tID)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			var body string
			readFile, err := os.ReadFile("./csv/themes/" + tID + "/" + file.Name())
			if err != nil {
				return errorx.New("خطایی در انتخاب قالب رخ داده است", "model", err)
			}
			body = string(readFile)

			model := new(DTOs.CreatePage)
			gofakeit.Struct(model)
			model.Title = file.Name()
			model.Body = body
			model.Type = "blank"
			model.ShopID = shopID
			model.Slug = file.Name()
			for _, page := range pages {
				if page.Slug == file.Name() {
					continue
				}
			}
			_, err = m.CreatePage(*model)
			if err != nil {
				return err
			}
		}

	}
	return nil
}
