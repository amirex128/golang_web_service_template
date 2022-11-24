package models

import (
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"go.elastic.co/apm/v2"
)

type Order struct {
	ID                        uint64      `gorm:"primary_key;auto_increment" json:"id"`
	UserID                    uint64      `json:"user_id"`
	User                      User        `gorm:"foreignKey:user_id" json:"user"`
	ShopID                    uint64      `json:"shop_id"`
	Shop                      Shop        `gorm:"foreignKey:shop_id" json:"shop"`
	CustomerID                uint64      `json:"customer_id"`
	Customer                  Customer    `gorm:"foreignKey:customer_id" json:"customer"`
	DiscountID                uint64      `json:"discount_id"`
	Discount                  Discount    `gorm:"foreignKey:discount_id" json:"discount"`
	AddressID                 uint64      `json:"address_id"`
	OrderItems                []OrderItem `gorm:"foreignKey:order_id" json:"order_items"`
	IP                        string      `json:"ip"`
	TotalProductPrice         float32     `json:"total_product_price"`
	TotalDiscountPrice        float32     `json:"total_discount_price"`
	TotalTaxPrice             float32     `json:"total_tax_price"`
	TotalProductDiscountPrice float32     `json:"total_product_discount_price"`
	TotalFinalPrice           float32     `json:"total_final_price"`
	SendPrice                 float32     `json:"send_price"`
	Status                    string      `json:"status"`
	Description               string      `json:"description"`
	Weight                    uint32      `json:"weight"`
	PackageSize               string      `json:"package_size"`
	TrackingCode              string      `json:"tracking_code"`
	Courier                   string      `json:"courier"` // tipax post post-poshtaz
	LastUpdateStatusAt        string      `json:"last_update_status_at"`
	CreatedAt                 string      `json:"created_at"`
}

func initOrder(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Order{})
}

func (m *MysqlManager) CreateOrder(order *Order) (*Order, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreateOrder", "model")
	defer span.End()
	err := m.GetConn().Create(order).Error
	if err != nil {
		return nil, errorx.New("خطا در ثبت سفارش", "model", err)
	}
	return order, nil
}

func (m *MysqlManager) GetOrders(userID uint64, orderStatus []string) ([]*Order, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetOrders", "model")
	defer span.End()
	var orders []*Order
	err := m.GetConn().Where("user_id = ? AND status IN (?)", userID, orderStatus).Find(&orders).Error
	if err != nil {
		return orders, errorx.New("خطا در دریافت سفارشات", "model", err)
	}
	return orders, nil
}
func (m *MysqlManager) FindOrdersByCustomerID(customerID uint64) ([]*Order, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindOrdersByCustomerID", "model")
	defer span.End()
	var orders []*Order
	err := m.GetConn().Where("customer_id = ?", customerID).Find(&orders).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت سفارشات", "model", err)
	}
	return orders, nil
}

func (m *MysqlManager) FindOrderByID(orderID uint64) (*Order, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindOrderByID", "model")
	defer span.End()
	var order *Order
	err := m.GetConn().Where("id = ?", orderID).First(&order).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت سفارش", "model", err)
	}
	return order, nil
}

func (m *MysqlManager) UpdateOrder(order Order) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:UpdateOrder", "model")
	defer span.End()
	err := m.GetConn().Save(&order).Error
	if err != nil {
		return errorx.New("خطا در ثبت سفارش", "model", err)
	}
	return nil
}

func (m *MysqlManager) FindOrderWithItemByID(orderID uint64) (*Order, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindOrderWithItemByID", "model")
	defer span.End()
	var order *Order
	err := m.GetConn().Where("id = ?", orderID).Preload("OrderItems").Preload("Shop").Preload("Customer").First(order).Error
	if err != nil {

		return nil, errorx.New("خطا در دریافت سفارش", "model", err)
	}
	return order, nil
}
