package models

import (
	"encoding/gob"
	"io"
)

type GroupDiscount struct {
	ID        int64   `gorm:"primary_key;auto_increment" json:"id"`
	ProductID int64   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Value     float32 `json:"value"`
	Type      string  `json:"type"` // percent or fixed
	Percent   string  `json:"percent"`
	StartedAt string  `json:"started_at"`
	EndedAt   string  `json:"ended_at"`
	CreatedAt string  `json:"created_at"`
}
type GroupDiscountArr []GroupDiscount

func (s GroupDiscountArr) Len() int {
	return len(s)
}
func (s GroupDiscountArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s GroupDiscountArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *GroupDiscount) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *GroupDiscount) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initGroupDiscount(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&GroupDiscount{})
}
