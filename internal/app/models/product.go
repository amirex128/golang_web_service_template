package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"database/sql"
	"encoding/gob"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type Product struct {
	ID               uint64         `json:"id"`
	UserId           uint64         `json:"user_id"`
	Description      string         `json:"description"`
	Name             string         `json:"name"`
	ShortDescription string         `json:"short_description"`
	TotalSales       uint32         `json:"total_sales"`
	BlockStatus      string         `json:"block_status"` // ok filter black_list disable_ok disable_filter disable_black_list
	Quantity         int            `json:"quantity"`
	Price            float32        `json:"price"`
	FreeSend         byte           `json:"free_send"`
	Weight           sql.NullInt32  `json:"weight"`
	Height           sql.NullInt32  `json:"height"`
	Width            sql.NullInt32  `json:"width"`
	Active           byte           `json:"active"`
	Images           string         `json:"images"`
	CreatedAt        string         `json:"created_at"`
	UpdatedAt        string         `json:"updated_at"`
	StartedAt        sql.NullString `json:"started_at"`
	EndedAt          sql.NullString `json:"ended_at"`
	DeliveryTime     uint           `json:"delivery_time"` // مدت زمان ارسال
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
}

func (m *MysqlManager) IndexProduct(dto DTOs.IndexProduct) (*DTOs.Pagination, error) {
	conn := m.GetConn()
	var products []Product
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate(ProductTable, pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%")
	}
	err := conn.Find(&products).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = products
	return pagination, nil
}

func (m *MysqlManager) CreateProduct(c *gin.Context, dto DTOs.CreateProduct, userID uint64) error {

	var product = Product{
		UserId:           userID,
		Description:      dto.Description,
		Name:             dto.Name,
		ShortDescription: dto.ShortDescription,
		Quantity: func() int {
			if dto.Quantity == 0 {
				return -1
			}
			return dto.Quantity
		}(),
		Price:    dto.Price,
		FreeSend: utils.ActiveConvert(dto.FreeSend),
		Weight: func() sql.NullInt32 {
			if dto.Weight > 0 {
				return sql.NullInt32{
					Int32: int32(dto.Weight),
					Valid: true,
				}
			}
			return sql.NullInt32{
				Int32: 0,
				Valid: false,
			}
		}(),
		Height: func() sql.NullInt32 {
			if dto.Height > 0 {
				return sql.NullInt32{
					Int32: int32(dto.Height),
					Valid: true,
				}
			}
			return sql.NullInt32{
				Int32: 0,
				Valid: false,
			}
		}(),
		Width: func() sql.NullInt32 {
			if dto.Width > 0 {
				return sql.NullInt32{
					Int32: int32(dto.Width),
					Valid: true,
				}
			}
			return sql.NullInt32{
				Int32: 0,
				Valid: false,
			}
		}(),
		Active:    1,
		Images:    strings.Join(dto.ImagePath, ","),
		CreatedAt: utils.NowTime(),
		UpdatedAt: utils.NowTime(),
		StartedAt: func() sql.NullString {
			if dto.StartedAt == "" {
				return sql.NullString{
					String: "",
					Valid:  false,
				}
			}
			return sql.NullString{
				String: utils.DateTimeConvert(dto.StartedAt),
				Valid:  true,
			}
		}(),
		EndedAt: func() sql.NullString {
			if dto.EndedAt == "" {
				return sql.NullString{
					String: "",
					Valid:  false,
				}
			}
			return sql.NullString{
				String: utils.DateTimeConvert(dto.EndedAt),
				Valid:  true,
			}
		}(),
		DeliveryTime: dto.DeliveryTime,
	}

	err := m.GetConn().Create(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در ایجاد محصول"})
		return err
	}
	return nil
}

func (m *MysqlManager) UpdateProduct(c *gin.Context, dto DTOs.UpdateProduct) error {
	product, err := m.FindProductById(c, dto.ID)
	if err != nil {
		return err
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
	if dto.Quantity > 0 || dto.Quantity == -1 {
		product.Quantity = dto.Quantity
	}
	if dto.Price > 0 {
		product.Price = dto.Price
	}
	if dto.FreeSend != "" {
		product.FreeSend = utils.ActiveConvert(dto.FreeSend)
	}
	if dto.Weight > 0 {
		product.Weight = sql.NullInt32{
			Int32: int32(dto.Weight),
			Valid: true,
		}
	}
	if dto.Height > 0 {
		product.Height = sql.NullInt32{
			Int32: int32(dto.Height),
			Valid: true,
		}
	}
	if dto.Width > 0 {
		product.Width = sql.NullInt32{
			Int32: int32(dto.Width),
			Valid: true,
		}
	}
	if dto.Active != "" {
		product.Active = utils.ActiveConvert(dto.Active)
	}
	if len(dto.ImagePath) > 0 {
		product.Images = strings.Join(dto.ImagePath, ",")
	}
	product.UpdatedAt = utils.NowTime()
	if dto.StartedAt != "" {
		product.StartedAt = sql.NullString{
			String: utils.DateTimeConvert(dto.StartedAt),
			Valid:  true,
		}
	}
	if dto.EndedAt != "" {
		product.EndedAt = sql.NullString{
			String: utils.DateTimeConvert(dto.EndedAt),
			Valid:  true,
		}
	}
	if dto.DeliveryTime > 0 {
		product.DeliveryTime = dto.DeliveryTime
	}

	err = m.GetConn().Save(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در بروزرسانی محصول"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در حذف محصول"})
		return err
	}
	return nil
}

func (m *MysqlManager) FindProductById(c *gin.Context, id uint64) (Product, error) {
	var product Product
	err := m.GetConn().Where("id = ?", id).First(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در دریافت محصول"})
		return product, err
	}
	return product, nil
}
func (m *MysqlManager) FindProductByIds(c *gin.Context, ids []uint64) ([]Product, error) {
	var products []Product
	err := m.GetConn().Where("id IN (?)", ids).First(&products).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در دریافت محصولات"})
		return products, err
	}
	return products, nil
}

func (m *MysqlManager) CheckAccessProduct(c *gin.Context, id uint64, userID uint64) error {
	product, err := m.FindProductById(c, id)
	if err != nil {
		return err
	}
	if product.UserId != userID {
		err = errors.New("access denied")
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "شما دسترسی کافی برای ویرایش این محصول را ندارید"})
		return err
	}

	return nil
}
