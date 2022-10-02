package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type Product struct {
	ID               uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID           uint64    `json:"user_id"`
	User             User      `gorm:"foreignKey:user_id" json:"user"`
	ShopID           uint64    `json:"shop_id"`
	Shop             Shop      `gorm:"foreignKey:shop_id" json:"shop"`
	Description      string    `json:"description"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"short_description"`
	TotalSales       uint32    `json:"total_sales"`
	Status           string    `json:"block_status" sql:"type:ENUM('block','ok')"`
	Quantity         uint32    `json:"quantity"`
	Price            float32   `json:"price"`
	Weight           uint32    `json:"weight"`
	Height           uint32    `json:"height"`
	Width            uint32    `json:"width"`
	Active           byte      `json:"active"`
	Images           string    `json:"images"`
	CreatedAt        string    `json:"created_at"`
	UpdatedAt        string    `json:"updated_at"`
	StartedAt        string    `json:"started_at"`
	EndedAt          string    `json:"ended_at"`
	CategoryID       uint64    `json:"category_id"`
	Category         Category  `gorm:"foreignKey:category_id" json:"category"`
	Galleries        []Gallery `gorm:"foreignKey:product_id" json:"galleries"`
	DeliveryTime     uint32    `json:"delivery_time"` // مدت زمان ارسال
}

func (c Product) GetID() uint64 {
	return c.ID
}

type ProductArr []Product

func (s ProductArr) Len() int {
	return len(s)
}
func (s ProductArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ProductArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Product) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Product) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func InitProduct(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Product{})
	for i := 0; i < 100; i++ {
		manager.CreateProduct(&gin.Context{}, DTOs.CreateProduct{
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
			Images:           nil,
			ImagePath:        nil,
			CategoryID:       1,
		}, 1)
	}
}

func (m *MysqlManager) GetAllProductWithPagination(c *gin.Context, dto DTOs.IndexProduct) (*DTOs.Pagination, error) {
	conn := m.GetConn()
	var products []Product
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate(ProductTable, pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%").Order("id DESC")
	}
	err := conn.Find(&products).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در دریافت محصولات",
			"error":   err.Error(),
		})
		return pagination, err
	}
	pagination.Data = products
	return pagination, nil
}

func (m *MysqlManager) CreateProduct(c *gin.Context, dto DTOs.CreateProduct, userID uint64) error {

	var product = Product{
		UserID:           userID,
		ShopID:           dto.ShopID,
		Description:      dto.Description,
		Name:             dto.Name,
		ShortDescription: dto.ShortDescription,
		Quantity:         dto.Quantity,
		Price:            dto.Price,
		Weight:           dto.Weight,
		Height:           dto.Height,
		Width:            dto.Width,
		Active:           1,
		CategoryID:       dto.CategoryID,
		Images:           strings.Join(dto.ImagePath, ","),
		CreatedAt:        utils.NowTime(),
		UpdatedAt:        utils.NowTime(),
		StartedAt:        utils.DateTimeConvert(dto.StartedAt),
		EndedAt:          utils.DateTimeConvert(dto.EndedAt),
		DeliveryTime:     dto.DeliveryTime,
	}

	err := m.GetConn().Create(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "خطا در ایجاد محصول",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) UpdateProduct(c *gin.Context, dto DTOs.UpdateProduct) error {
	product, err := m.FindProductById(c, dto.ID)
	if err != nil {
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
	if dto.Weight > 0 {
		product.Weight = dto.Weight
	}
	if dto.Height > 0 {
		product.Height = dto.Height
	}
	if dto.Width > 0 {
		product.Width = dto.Width
	}
	if dto.Active != "" {
		product.Active = utils.ActiveConvert(dto.Active)
	}
	if len(dto.ImagePath) > 0 {
		product.Images = strings.Join(dto.ImagePath, ",")
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
			"message": "خطا در بروزرسانی محصول",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) DeleteProduct(c *gin.Context, id uint64) error {
	product, err := m.FindProductById(c, id)
	if err != nil {
		return err
	}

	utils.RemoveImages(strings.Split(product.Images, ","))

	err = m.GetConn().Delete(&Product{}, id).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "خطا در حذف محصول",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) FindProductById(c *gin.Context, id uint64) (Product, error) {
	var product Product
	err := m.GetConn().Where("id = ?", id).First(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "خطا در دریافت محصول",
		})
		return product, err
	}
	return product, nil
}
func (m *MysqlManager) FindProductByIds(c *gin.Context, ids []uint64) ([]Product, error) {
	var products []Product
	err := m.GetConn().Where("id IN (?)", ids).Find(&products).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "خطا در دریافت محصولات",
		})
		return products, err
	}
	return products, nil
}

func (m *MysqlManager) CheckAccessProduct(c *gin.Context, id uint64, userID uint64) error {
	product, err := m.FindProductById(c, id)
	if err != nil {
		return err
	}
	if product.UserID != userID {
		err = errors.New("access denied")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   err.Error(),
			"message": "شما دسترسی کافی برای ویرایش این محصول را ندارید",
		})
		return err
	}

	return nil
}
