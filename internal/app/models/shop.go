package models

import (
	"encoding/gob"
	"io"
)

type Shop struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	Name      string `json:"name"`
	Logo      string `json:"logo"`
	Type      string `json:"type" sql:"type:ENUM('instagram','telegram','website')"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type ShopArr []Shop

func (s ShopArr) Len() int {
	return len(s)
}
func (s ShopArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ShopArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Shop) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Shop) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initShop(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Shop{})
}
