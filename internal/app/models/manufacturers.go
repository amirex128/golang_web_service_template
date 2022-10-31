package models

import (
	"backend/internal/app/utils"
	"encoding/gob"
	"github.com/sirupsen/logrus"
	"io"
)

type Manufacturer struct {
	ID               int    `gorm:"primary_key;auto_increment" json:"id"`
	CountryName      string `json:"country_name"`
	ShortDescription string `json:"short_description"`
	PersianName      string `json:"persian_name"`
	EnglishName      string `json:"english_name"`
}
type ManufacturerArr []Manufacturer

func (s ManufacturerArr) Len() int {
	return len(s)
}
func (s ManufacturerArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ManufacturerArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Manufacturer) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Manufacturer) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initManufacturer(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Manufacturer{})
	manufacturer := utils.ReadCsvFile("../../csv/manufacturers.csv")
	manager.CreateAllManufacturer(manufacturer)
}

func (m *MysqlManager) CreateAllManufacturer(files [][]string) {
	manufacturer := make([]Manufacturer, 0)
	for i := range files {
		value := files[i]
		manufacturer = append(manufacturer, Manufacturer{
			ID:               utils.StringToInt(value[0]),
			CountryName:      value[2],
			ShortDescription: value[3],
			PersianName:      value[5],
			EnglishName:      value[6],
		})
	}
	err := m.GetConn().CreateInBatches(manufacturer, 100).Error
	if err != nil {
		logrus.Error("seed manufacturer error: ", err)
	}
}
