package models

import (
	"database/sql"
	utils2 "github.com/amirex128/selloora_backend/internal/utils"
	"github.com/sirupsen/logrus"
)

type City struct {
	ID          int            `gorm:"primary_key;auto_increment" json:"id"`
	ProvinceID  int            `json:"province_id"`
	PersianName string         `json:"persian_name"`
	EnglishName sql.NullString `json:"english_name"`
	Code        sql.NullString `json:"code"`
	Lat         sql.NullString `json:"lat"`
	Lng         sql.NullString `json:"lng"`
}

func initCity(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&City{})
	cities := utils2.ReadCsvFile("./csv/cities.csv")
	manager.CreateAllCities(cities)
}
func (m *MysqlManager) CreateAllCities(files [][]string) {
	cities := make([]City, 0)
	for i := range files {
		value := files[i]
		cities = append(cities, City{
			ID:          utils2.StringToInt(value[0]),
			ProvinceID:  utils2.StringToInt(value[1]),
			PersianName: value[2],
			EnglishName: utils2.StringConvert(value[5]),
			Code:        utils2.StringConvert(value[6]),
			Lat:         utils2.StringConvert(value[7]),
			Lng:         utils2.StringConvert(value[8]),
		})
	}
	err := m.GetConn().CreateInBatches(cities, 100).Error
	if err != nil {
		logrus.Error("seed cities error: ", err)
	}
}
