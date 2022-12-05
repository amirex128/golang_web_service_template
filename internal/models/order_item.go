package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"go.elastic.co/apm/v2"
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
	if !manager.GetConn().Migrator().HasTable(&OrderItem{}) {
		manager.GetConn().AutoMigrate(&OrderItem{})

	}

}

func (m *MysqlManager) CreateOrderItem(dto []DTOs.OrderItem, orderID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreateOrderItem", "model")
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
		return errorx.New("خطا در ثبت سفارش", "model", err)
	}
	return nil
}
