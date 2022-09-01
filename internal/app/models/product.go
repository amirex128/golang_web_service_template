package models

import (
	"backend/internal/app/DTOs"
	"encoding/gob"
	jwt "github.com/appleboy/gin-jwt/v2"
	"io"
	"time"
)

type Product struct {
	ID               int64   `json:"id"`
	UserId           string  `json:"user_id"`
	ManufacturerId   int     `json:"manufacturer_id"`
	Description      string  `json:"description"`
	Name             string  `json:"name"`
	ShortDescription string  `json:"short_description"`
	TotalSales       string  `json:"total_sales"`
	TotalVisitors    string  `json:"total_visitors"`
	BlockStatus      string  `json:"block_status"` // ok filter black_list disable_ok disable_filter disable_black_list
	Quantity         int     `json:"quantity"`
	Price            float32 `json:"price"`
	FreeSend         string  `json:"free_send"`
	Weight           uint    `json:"weight"`
	Height           uint    `json:"height"`
	Width            uint    `json:"width"`
	Active           byte    `json:"active"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	StartedAt        string  `json:"started_at"`
	EndedAt          string  `json:"ended_at"`
	DeliveryTime     string  `json:"delivery_time"` // مدت زمان ارسال
	OptionId         int     `json:"option_id"`
	OptionItemID     int     `json:"option_item_id"`
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

func (m *MysqlManager) CreateProduct(dto DTOs.CreateProduct, claims jwt.MapClaims) error {
	l, _ := time.LoadLocation("Asia/Tehran")
	startedAt, err := time.ParseInLocation("2006-01-02 15:04:05", dto.StartedAt, l)
	if err != nil {
		return err
	}
	endedAt, err := time.ParseInLocation("2006-01-02 15:04:05", dto.EndedAt, l)
	if err != nil {
		return err
	}
	var product = Product{
		UserId:           claims["id"].(string),
		ManufacturerId:   dto.ManufacturerId,
		Description:      dto.Description,
		Name:             dto.Name,
		ShortDescription: dto.ShortDescription,
		Quantity:         dto.Quantity,
		Price:            dto.Price,
		FreeSend:         dto.FreeSend,
		Weight:           dto.Weight,
		Height:           dto.Height,
		Width:            dto.Width,
		Active:           dto.Active,
		CreatedAt:        time.Now().In(l).Format("2006-01-02 15:04:05"),
		UpdatedAt:        time.Now().In(l).Format("2006-01-02 15:04:05"),
		StartedAt:        startedAt.String(),
		EndedAt:          endedAt.String(),
		DeliveryTime:     dto.DeliveryTime,
		OptionId:         dto.OptionId,
		OptionItemID:     dto.OptionItemID,
	}

	err = m.GetConn().Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}
