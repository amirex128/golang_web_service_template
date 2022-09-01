package models

import (
	"backend/internal/app/DTOs"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"io"
)

type Order struct {
	ID                 int64  `json:"id"`
	UserID             int64  `json:"user_id"`
	CustomerID         int64  `json:"customer_id"`
	DiscountID         int64  `json:"discount_id"`
	IP                 string `json:"ip"`
	Status             string `json:"status"` // suspend ready wrong_ready seller_not_to_attend personal_sent service_sent virtual_sent posted unacceptable waited not_distribution pre_distribution distributed confirmed accept return return_final cancel khesarat gheramati amadesazi merge inprocessing ready_schedule logistic_sent pre_return_logistic ready_logistic
	LastUpdateStatusAt string `json:"last_update_status_at"`
	CreatedAt          string `json:"created_at"`
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

func (m *MysqlManager) CreateOrder(c *gin.Context, dto DTOs.CreateOrder) error {

	return nil
}
