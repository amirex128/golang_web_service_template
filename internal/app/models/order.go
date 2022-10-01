package models

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Order struct {
	ID                        uint64      `gorm:"primary_key;auto_increment" json:"id"`
	UserID                    uint64      `json:"user_id"`
	User                      User        `gorm:"foreignKey:user_id" json:"user"`
	CustomerID                uint64      `json:"customer_id"`
	Customer                  Customer    `gorm:"foreignKey:customer_id" json:"customer"`
	DiscountID                uint64      `json:"discount_id"`
	Discount                  Discount    `gorm:"foreignKey:discount_id" json:"discount"`
	OrderItems                []OrderItem `gorm:"foreignKey:order_id" json:"order_items"`
	IP                        string      `json:"ip"`
	TotalProductPrice         float32     `json:"total_product_price"`
	TotalDiscountPrice        float32     `json:"total_discount_price"`
	TotalTaxPrice             float32     `json:"total_tax_price"`
	TotalProductDiscountPrice float32     `json:"total_product_discount_price"`
	TotalFinalPrice           float32     `json:"total_final_price"`
	SendPrice                 float32     `json:"send_price"`
	Status                    string      `json:"status"`
	SendType                  string      `json:"send_type"` // tipax post post-poshtaz
	LastUpdateStatusAt        string      `json:"last_update_status_at"`
	CreatedAt                 string      `json:"created_at"`
}

type OrderArr []Order

func (s OrderArr) Len() int {
	return len(s)
}
func (s OrderArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s OrderArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Order) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Order) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initOrder(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Order{})
}

func (m *MysqlManager) CreateOrder(c *gin.Context, order Order) (err error) {
	err = m.GetConn().Create(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در ثبت سفارش",
		})
		return
	}
	return
}

func (m *MysqlManager) GetOrders(c *gin.Context, userID uint64, orderStatus []string) (orders []Order, err error) {
	err = m.GetConn().Where("user_id = ? AND status IN (?)", userID, orderStatus).Find(&orders).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در دریافت سفارشات",
		})
		return
	}
	return
}
func (m *MysqlManager) FindOrderByID(c *gin.Context, orderID uint64, userID uint64) (order Order, err error) {
	err = m.GetConn().Where("id = ?", orderID).First(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در دریافت سفارش",
		})
		return
	}
	if order.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "شما اجازه دسترسی به این سفارش را ندارید",
		})
		return
	}
	return
}
func (m *MysqlManager) UpdateOrder(c *gin.Context, order Order) (err error) {
	err = m.GetConn().Save(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در ثبت سفارش",
		})
		return
	}
	return
}
