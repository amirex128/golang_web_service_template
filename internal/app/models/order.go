package models

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Order struct {
	ID                        uint64  `json:"id"`
	UserID                    uint64  `json:"user_id"`
	CustomerID                uint64  `json:"customer_id"`
	DiscountID                uint64  `json:"discount_id"`
	IP                        string  `json:"ip"`
	TotalProductPrice         float32 `json:"total_product_price"`
	TotalDiscountPrice        float32 `json:"total_discount_price"`
	TotalTaxPrice             float32 `json:"total_tax_price"`
	TotalProductDiscountPrice float32 `json:"total_product_discount_price"`
	TotalFinalPrice           float32 `json:"total_final_price"`
	SendPrice                 float32 `json:"send_price"`
	Status                    string  `json:"status"`
	PaymentStatus             string  `json:"payment_status"`
	SendType                  string  `json:"send_type"` // tipax post post-poshtaz
	LastUpdateStatusAt        string  `json:"last_update_status_at"`
	CreatedAt                 string  `json:"created_at"`
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

func (m *MysqlManager) CreateOrder(c *gin.Context, order Order) error {
	err := m.GetConn().Create(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در ثبت سفارش",
		})
		return err
	}
	return nil
}
