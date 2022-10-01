package models

import (
	"encoding/gob"
	"io"
)

type OrderItem struct {
	ID        int64 `gorm:"primary_key;auto_increment" json:"id"`
	ProductID int64 `json:"product_id"`
	Count     uint  `json:"count"`
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
