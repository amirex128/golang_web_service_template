package models

import (
	"encoding/gob"
	"github.com/sirupsen/logrus"
	"io"
)

type OptionItem struct {
	ID       int    `json:"id"`
	OptionID int    `json:"option_id"`
	Name     string `json:"name"`
	Value    string `json:"value"`
	Sort     uint   `json:"sort"`
}
type OptionValueArr []OptionItem

func (s OptionValueArr) Len() int {
	return len(s)
}
func (s OptionValueArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s OptionValueArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *OptionItem) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *OptionItem) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}

func (m *MysqlManager) CreateAllOptionItems(files [][]string) {
	optionItems := make([]OptionItem, 0)
	for i := range files {
		value := files[i]
		optionItems = append(optionItems, OptionItem{
			ID:       intConvert(value[0]),
			OptionID: intConvert(value[1]),
			Name:     value[2],
			Value:    value[3],
			Sort:     uintConvert(value[4]),
		})
	}
	err := m.GetConn().CreateInBatches(optionItems, 100).Error
	if err != nil {
		logrus.Error("seed optionItems error: ", err)
	}
}
