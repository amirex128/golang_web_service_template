package models

import (
	"backend/internal/app/helpers"
	"encoding/gob"
	"github.com/sirupsen/logrus"
	"io"
)

type Option struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type" sql:"type:ENUM('text','color')"`
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

func (m *MysqlManager) CreateAllOptions(files [][]string) {
	options := make([]Option, 0)
	for i := range files {
		value := files[i]
		options = append(options, Option{
			ID:          int32Convert(value[0]),
			Name:        value[1],
			Description: value[2],
			Type:        value[3],
		})
	}
	err := m.GetConn().CreateInBatches(options, 100).Error
	if err != nil {
		logrus.Error("seed options error: ", err)
	}
}
func initOption(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Option{})
	initOptionItem(manager)
	options := helpers.ReadCsvFile("../../csv/options.csv")
	manager.CreateAllOptions(options)

}
