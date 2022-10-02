package models

import (
	"backend/internal/app/DTOs"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type OrderItem struct {
	ID        uint64  `gorm:"primary_key;auto_increment" json:"id"`
	OrderID   uint64  `json:"order_id"`
	Order     Order   `gorm:"foreignKey:order_id" json:"order"`
	ProductID uint64  `json:"product_id"`
	Product   Product `gorm:"foreignKey:product_id" json:"product"`
	Count     uint32  `json:"count"`
	OptionID  uint64  `json:"option_id"`
}
type OrderItemArr []OrderItem

func (s OrderItemArr) Len() int {
	return len(s)
}
func (s OrderItemArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s OrderItemArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *OrderItem) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *OrderItem) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initOrderItem(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&OrderItem{})
}

func (m *MysqlManager) CreateOrderItem(c *gin.Context, dto []DTOs.OrderItem, orderID uint64) error {
	var orderItems []OrderItem
	for i := range dto {
		orderItem := OrderItem{
			OrderID:   orderID,
			ProductID: dto[i].ProductID,
			Count:     dto[i].Count,
			OptionID:  dto[i].OptionID,
		}
		orderItems = append(orderItems, orderItem)
	}
	err := m.GetConn().Create(&orderItems).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "خطا در ثبت سفارش",
			"error":   err.Error(),
		})
		return err
	}
	return nil
}
