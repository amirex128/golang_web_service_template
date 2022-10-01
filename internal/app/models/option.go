package models

import (
	"encoding/gob"
	"io"
)

type Option struct {
	ID         uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Variant    string     `json:"variant"`
	Name       string     `json:"name"`
	Price      float32    `json:"price"`
	Quantity   uint32     `json:"quantity"`
	Categories []Category `gorm:"many2many:category_options;" json:"categories"`
}
type OptionArr []Option

func (s OptionArr) Len() int {
	return len(s)
}
func (s OptionArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s OptionArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Option) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Option) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initOption(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Option{})
}
