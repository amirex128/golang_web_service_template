package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

type Product struct {
	ID               uint64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID           uint64     `json:"user_id"`
	User             User       `gorm:"foreignKey:user_id" json:"user"`
	ShopID           uint64     `json:"shop_id"`
	Shop             Shop       `gorm:"foreignKey:shop_id" json:"shop"`
	Description      string     `json:"description"`
	Name             string     `json:"name"`
	ShortDescription string     `json:"short_description"`
	TotalSales       uint32     `json:"total_sales"`
	Status           string     `json:"block_status" sql:"type:ENUM('block','ok')"`
	Quantity         uint32     `json:"quantity"`
	Price            float32    `json:"price"`
	Active           byte       `json:"active"`
	CreatedAt        string     `json:"created_at"`
	UpdatedAt        string     `json:"updated_at"`
	StartedAt        string     `json:"started_at"`
	EndedAt          string     `json:"ended_at"`
	Categories       []Category `gorm:"many2many:category_product;" json:"categories"`
	Galleries        []Gallery  `gorm:"many2many:gallery_product;" json:"galleries"`
	DeliveryTime     uint32     `json:"delivery_time"` // مدت زمان ارسال
}

func (c Product) GetID() uint64 {
	return c.ID
}

func InitProduct(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Product{})
	for i := 0; i < 100; i++ {
		manager.CreateProduct(&gin.Context{}, context.Background(), DTOs.CreateProduct{
			ShopID:           1,
			ManufacturerId:   1,
			Description:      fmt.Sprintf("توضیحات محصول %d", i),
			Name:             fmt.Sprintf("محصول %d", i),
			ShortDescription: fmt.Sprintf("توضیحات کوتاه محصول %d", i),
			Quantity:         10,
			Price:            10000,
			Weight:           100,
			Height:           100,
			Width:            100,
			StartedAt:        "2020-01-01 00:00:00",
			EndedAt:          "2024-01-01 00:00:00",
			DeliveryTime:     1,
			OptionId:         1,
			OptionItemID:     1,
			CategoryID:       1,
			GalleryIDs:       []uint64{1},
		}, 1)
	}
}

func (m *MysqlManager) GetAllProductWithPagination(c *gin.Context, ctx context.Context, dto DTOs.IndexProduct) (*DTOs.Pagination, error) {
	span, ctx := apm.StartSpan(ctx, "GetAllProductWithPagination", "model")
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در دریافت محصولات",
			"error":   err.Error(),
			"type":    "model",
		})
		return pagination, err
	}
	pagination.Data = products
	return pagination, nil
}

func (m *MysqlManager) CreateProduct(c *gin.Context, ctx context.Context, dto DTOs.CreateProduct, userID uint64) error {
	span, ctx := apm.StartSpan(ctx, "CreateProduct", "model")
	defer span.End()

	var product = Product{
		UserID:           userID,
		ShopID:           dto.ShopID,
		Description:      dto.Description,
		Name:             dto.Name,
		ShortDescription: dto.ShortDescription,
		Quantity:         dto.Quantity,
		Price:            dto.Price,
		Active:           1,
		CreatedAt:        utils.NowTime(),
		UpdatedAt:        utils.NowTime(),
		StartedAt:        utils.DateTimeConvert(dto.StartedAt),
		EndedAt:          utils.DateTimeConvert(dto.EndedAt),
		DeliveryTime:     dto.DeliveryTime,
	}

	var galleryIDs []Gallery
	for i := range dto.GalleryIDs {
		galleryIDs = append(galleryIDs, Gallery{ID: dto.GalleryIDs[i]})

	}
	err := m.GetConn().Create(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"type":    "model",
			"message": "خطا در ایجاد محصول",
		})
		return err
	}
	err = m.GetConn().Model(&product).Association("Galleries").Append(galleryIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"type":    "model",
			"message": "خطا در ایجاد محصول",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) UpdateProduct(c *gin.Context, ctx context.Context, dto DTOs.UpdateProduct) error {
	span, ctx := apm.StartSpan(ctx, "UpdateProduct", "model")
	defer span.End()
	product, err := m.FindProductById(c, ctx, dto.ID)
	if err != nil {
		return err
	}

	userID := GetUser(c)
	if product.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "خطای دسترسی",
			"type":    "model",
			"message": "شما دسترسی کافی برای ویرایش این محصول را ندارید",
		})
		return err
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
	if dto.ShortDescription != "" {
		product.ShortDescription = dto.ShortDescription
	}
	if dto.Quantity > 0 {
		product.Quantity = dto.Quantity
	}
	if dto.Price > 0 {
		product.Price = dto.Price
	}
	if dto.Active != "" {
		product.Active = utils.ActiveConvert(dto.Active)
	}
	product.UpdatedAt = utils.NowTime()
	if dto.StartedAt != "" {
		product.StartedAt = utils.DateTimeConvert(dto.StartedAt)
	}
	if dto.EndedAt != "" {
		product.EndedAt = utils.DateTimeConvert(dto.EndedAt)
	}
	if dto.DeliveryTime > 0 {
		product.DeliveryTime = dto.DeliveryTime
	}

	err = m.GetConn().Save(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"type":    "model",
			"message": "خطا در بروزرسانی محصول",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) DeleteProduct(c *gin.Context, ctx context.Context, id uint64) error {
	span, ctx := apm.StartSpan(ctx, "DeleteProduct", "model")
	defer span.End()
	userID := GetUser(c)
	var product Product
	err := m.GetConn().Where("id = ?", id).First(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"type":    "model",
			"message": "خطا در حذف محصول",
		})
	}
	if product.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "خطای دسترسی",
			"type":    "model",
			"message": "شما دسترسی کافی برای ویرایش این محصول را ندارید",
		})
		return err
	}
	err = m.GetConn().Delete(&Product{}, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"type":    "model",
			"message": "خطا در حذف محصول",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) FindProductById(c *gin.Context, ctx context.Context, id uint64) (Product, error) {
	span, ctx := apm.StartSpan(ctx, "FindProductById", "model")
	defer span.End()
	var product Product
	err := m.GetConn().Where("id = ?", id).First(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"type":    "model",
			"message": "خطا در دریافت محصول",
		})
		return product, err
	}
	return product, nil
}

func (m *MysqlManager) FindProductByIds(c *gin.Context, ctx context.Context, ids []uint64) ([]Product, error) {
	span, ctx := apm.StartSpan(ctx, "FindProductByIds", "model")
	defer span.End()
	var products []Product
	err := m.GetConn().Where("id IN (?)", ids).Find(&products).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"type":    "model",
			"message": "خطا در دریافت محصولات",
		})
		return products, err
	}
	return products, nil
}

func (m *MysqlManager) MoveProducts(c *gin.Context, ctx context.Context, shopID, newShopID, userID uint64) error {
	span, ctx := apm.StartSpan(ctx, "MoveProducts", "model")
	defer span.End()
	err := m.GetConn().Model(&Product{}).Where("shop_id = ? AND user_id = ?", shopID, userID).Update("shop_id", newShopID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"type":    "model",
			"message": "خطا در انتقال محصولات",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) DeleteProducts(c *gin.Context, ctx context.Context, shopID, userID uint64) error {
	span, ctx := apm.StartSpan(ctx, "DeleteProducts", "model")
	defer span.End()
	err := m.GetConn().Where("shop_id = ? AND user_id = ?", shopID, userID).Delete(&Product{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"type":    "model",
			"message": "خطا در حذف محصولات",
		})
		return err
	}
	return nil

}
