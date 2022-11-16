package models

import (
	"context"
	"github.com/amirex128/selloora_backend/internal/app/DTOs"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
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

func initOrderItem(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&OrderItem{})
}

func (m *MysqlManager) CreateOrderItem(c *gin.Context, ctx context.Context, dto []DTOs.OrderItem, orderID uint64) error {
	span, ctx := apm.StartSpan(ctx, "CreateOrderItem", "model")
	defer span.End()
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
			"type":    "model",
		})
		return err
	}
	return nil
}
