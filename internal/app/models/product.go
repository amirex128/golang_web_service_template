package models

import (
	"backend/internal/app/DTOs"
	"database/sql"
	"encoding/gob"
	"io"
	"strings"
)

type Product struct {
	ID               uint64        `json:"id"`
	UserId           uint64        `json:"user_id"`
	ManufacturerId   sql.NullInt32 `json:"manufacturer_id"`
	Description      string        `json:"description"`
	Name             string        `json:"name"`
	ShortDescription string        `json:"short_description"`
	TotalSales       string        `json:"total_sales"`
	TotalVisitors    string        `json:"total_visitors"`
	BlockStatus      string        `json:"block_status"` // ok filter black_list disable_ok disable_filter disable_black_list
	Quantity         int           `json:"quantity"`
	Price            float32       `json:"price"`
	FreeSend         byte          `json:"free_send"`
	Weight           uint          `json:"weight"`
	Height           uint          `json:"height"`
	Width            uint          `json:"width"`
	Active           byte          `json:"active"`
	Images           string        `json:"images"`
	CreatedAt        string        `json:"created_at"`
	UpdatedAt        string        `json:"updated_at"`
	StartedAt        string        `json:"started_at"`
	EndedAt          string        `json:"ended_at"`
	DeliveryTime     uint          `json:"delivery_time"` // مدت زمان ارسال
	OptionId         uint          `json:"option_id"`
	OptionItemID     uint          `json:"option_item_id"`
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

func (m *MysqlManager) CreateProduct(dto DTOs.CreateProduct, userID uint64) error {

	var product = Product{
		UserId: userID,
		ManufacturerId: func() sql.NullInt32 {
			if dto.ManufacturerId > 0 {
				return sql.NullInt32{
					Int32: int32(dto.ManufacturerId),
					Valid: true,
				}
			}
			return sql.NullInt32{
				Int32: 0,
				Valid: false,
			}
		}(),
		Description:      dto.Description,
		Name:             dto.Name,
		ShortDescription: dto.ShortDescription,
		Quantity: func() int {
			if dto.Quantity == 0 {
				return -1
			}
			return dto.Quantity
		}(),
		Price:        dto.Price,
		FreeSend:     activeConvert(dto.FreeSend),
		Weight:       dto.Weight,
		Height:       dto.Height,
		Width:        dto.Width,
		Images:       strings.Join(dto.ImagePath, ","),
		CreatedAt:    nowTime(),
		UpdatedAt:    nowTime(),
		StartedAt:    dateTimeConvert(dto.StartedAt),
		EndedAt:      dateTimeConvert(dto.EndedAt),
		DeliveryTime: dto.DeliveryTime,
		OptionId:     dto.OptionId,
		OptionItemID: dto.OptionItemID,
	}

	err := m.GetConn().Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}
