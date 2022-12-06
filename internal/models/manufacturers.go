package models

import (
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/sirupsen/logrus"
)

type Manufacturer struct {
	ID               int    `gorm:"primary_key;auto_increment" json:"id"`
	CountryName      string `json:"country_name"`
	ShortDescription string `json:"short_description"`
	PersianName      string `json:"persian_name"`
	EnglishName      string `json:"english_name"`
}

func initManufacturer(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&Manufacturer{}) {
		manager.GetConn().Migrator().CreateTable(&Manufacturer{})
		var manufacturer [][]string
		if utils.IsTest() {
			manufacturer = utils.ReadCsvFile("../../../csv/manufacturers.csv")
		} else {
			manufacturer = utils.ReadCsvFile("./csv/manufacturers.csv")
		}
		manager.CreateAllManufacturer(manufacturer)

	}
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
