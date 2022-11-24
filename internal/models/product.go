package models

import (
	"fmt"
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"go.elastic.co/apm/v2"
)

type Product struct {
	ID           uint64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID       uint64     `json:"user_id"`
	User         User       `gorm:"foreignKey:user_id" json:"user"`
	ShopID       uint64     `json:"shop_id"`
	Shop         Shop       `gorm:"foreignKey:shop_id" json:"shop"`
	Description  string     `json:"description"`
	Name         string     `json:"name"`
	TotalSales   uint32     `json:"total_sales"`
	Status       string     `json:"block_status" sql:"type:ENUM('block','ok')"`
	Quantity     uint32     `json:"quantity"`
	Price        float32    `json:"price"`
	Active       byte       `json:"active"`
	CreatedAt    string     `json:"created_at"`
	UpdatedAt    string     `json:"updated_at"`
	StartedAt    string     `json:"started_at"`
	EndedAt      string     `json:"ended_at"`
	Tags         []Tag      `gorm:"many2many:product_tag;" json:"tags"`
	Categories   []Category `gorm:"many2many:category_product;" json:"categories"`
	Galleries    []Gallery  `gorm:"many2many:gallery_product;" json:"galleries"`
	DeliveryTime uint32     `json:"delivery_time"` // مدت زمان ارسال
}

func (c Product) GetID() uint64 {
	return c.ID
}

func InitProduct(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Product{})
	for i := 0; i < 100; i++ {
		manager.CreateProduct(DTOs.CreateProduct{
			ShopID:       1,
			Manufacturer: "سامسونگ",
			Description:  fmt.Sprintf("توضیحات محصول %d", i),
			Name:         fmt.Sprintf("محصول %d", i),
			Quantity:     10,
			Price:        10000,
			StartedAt:    "2020-01-01 00:00:00",
			EndedAt:      "2024-01-01 00:00:00",
			OptionId:     1,
			OptionItemID: 1,
			CategoryID:   1,
			GalleryIDs:   []uint64{1},
		}, 1)
	}
}

func (m *MysqlManager) GetAllProductWithPagination(dto DTOs.IndexProduct) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetAllProductWithPagination", "model")
	defer span.End()
	conn := m.GetConn()
	var products []Product
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("products", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%").Order("id DESC")
	}
	err := conn.Where("shop_id=?", dto.ShopID).Find(&products).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت محصولات", "model", err)
	}
	pagination.Data = products
	return pagination, nil
}

func (m *MysqlManager) CreateProduct(dto DTOs.CreateProduct, userID uint64) (*Product, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreateProduct", "model")
	defer span.End()

	var product = &Product{
		UserID:      userID,
		ShopID:      dto.ShopID,
		Description: dto.Description,
		Name:        dto.Name,
		Quantity:    dto.Quantity,
		Price:       dto.Price,
		Active:      1,
		CreatedAt:   utils.NowTime(),
		UpdatedAt:   utils.NowTime(),
		StartedAt:   utils.DateTimeConvert(dto.StartedAt),
		EndedAt:     utils.DateTimeConvert(dto.EndedAt),
	}

	var galleryIDs []Gallery
	for i := range dto.GalleryIDs {
		galleryIDs = append(galleryIDs, Gallery{ID: dto.GalleryIDs[i]})

	}
	err := m.GetConn().Create(product).Error
	if err != nil {
		return nil, errorx.New("خطا در ایجاد محصول", "model", err)
	}
	err = m.GetConn().Model(product).Association("Galleries").Append(galleryIDs)
	if err != nil {
		return product, errorx.New("خطا در ایجاد محصول", "model", err)
	}
	return product, nil
}

func (m *MysqlManager) UpdateProduct(dto DTOs.UpdateProduct) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:UpdateProduct", "model")
	defer span.End()
	product, err := m.FindProductById(dto.ID)
	if err != nil {
		return err
	}

	userID := GetUser(m.Ctx)
	if product.UserID != *userID {
		return errorx.New("شما دسترسی کافی برای ویرایش این محصول را ندارید", "model", err)
	}

	if dto.ShopID > 0 {
		product.ShopID = dto.ShopID
	}
	if dto.Description != "" {
		product.Description = dto.Description
	}
	if dto.Name != "" {
		product.Name = dto.Name
	}
	if dto.Quantity > 0 {
		product.Quantity = dto.Quantity
	}
	if dto.Price > 0 {
		product.Price = dto.Price
	}
	product.Active = utils.ActiveConvert(dto.Active)

	product.UpdatedAt = utils.NowTime()
	if dto.StartedAt != "" {
		product.StartedAt = utils.DateTimeConvert(dto.StartedAt)
	}
	if dto.EndedAt != "" {
		product.EndedAt = utils.DateTimeConvert(dto.EndedAt)
	}

	err = m.GetConn().Save(&product).Error
	if err != nil {
		return errorx.New("خطا در بروزرسانی محصول", "model", err)
	}
	return nil
}

func (m *MysqlManager) DeleteProduct(id uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:DeleteProduct", "model")
	defer span.End()
	userID := GetUser(m.Ctx)
	var product Product
	err := m.GetConn().Where("id = ?", id).First(&product).Error
	if err != nil {
		return errorx.New("خطا در حذف محصول", "model", err)
	}
	if product.UserID != *userID {
		return errorx.New("شما دسترسی کافی برای ویرایش این محصول را ندارید", "model", err)
	}
	err = m.GetConn().Delete(&Product{}, id).Error
	if err != nil {
		return errorx.New("خطا در حذف محصول", "model", err)
	}
	return nil
}

func (m *MysqlManager) FindProductById(id uint64) (*Product, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindProductById", "model")
	defer span.End()
	var product *Product
	err := m.GetConn().Where("id = ?", id).First(product).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت محصول", "model", err)
	}
	return product, nil
}

func (m *MysqlManager) FindProductByIds(ids []uint64) ([]*Product, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindProductByIds", "model")
	defer span.End()
	products := make([]*Product, 0)
	err := m.GetConn().Where("id IN (?)", ids).Find(&products).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت محصولات", "model", err)
	}
	return products, nil
}

func (m *MysqlManager) MoveProducts(shopID, newShopID, userID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:MoveProducts", "model")
	defer span.End()
	err := m.GetConn().Model(&Product{}).Where("shop_id = ? AND user_id = ?", shopID, userID).Update("shop_id", newShopID).Error
	if err != nil {
		return errorx.New("خطا در انتقال محصولات", "model", err)
	}
	return nil
}

func (m *MysqlManager) DeleteProducts(shopID, userID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:DeleteProducts", "model")
	defer span.End()
	err := m.GetConn().Where("shop_id = ? AND user_id = ?", shopID, userID).Delete(&Product{}).Error
	if err != nil {
		return errorx.New("خطا در حذف محصولات", "model", err)
	}
	return nil

}
