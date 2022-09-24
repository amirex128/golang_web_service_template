package models

import (
	"backend/internal/app/utils"
	"database/sql"
	"encoding/gob"
	"github.com/sirupsen/logrus"
	"io"
)

type City struct {
	ID          int            `json:"id"`
	ProvinceID  int            `json:"province_id"`
	PersianName string         `json:"persian_name"`
	EnglishName sql.NullString `json:"english_name"`
	Code        sql.NullString `json:"code"`
	Lat         sql.NullString `json:"lat"`
	Lng         sql.NullString `json:"lng"`
}

type CityProduct struct {
	ProductID int64 `json:"product_id"`
	CityID    int   `json:"city_id"`
}

type CityArr []City

func (s CityArr) Len() int {
	return len(s)
}
func (s CityArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s CityArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *City) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *City) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initCity(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&City{})
	manager.GetConn().AutoMigrate(&CityProduct{})
	cities := utils.ReadCsvFile("../../csv/cities.csv")
	manager.CreateAllCities(cities)
}
func (m *MysqlManager) CreateAllCities(files [][]string) {
	cities := make([]City, 0)
	for i := range files {
		value := files[i]
		cities = append(cities, City{
			ID:          utils.StringToInt(value[0]),
			ProvinceID:  utils.StringToInt(value[1]),
			PersianName: value[2],
			EnglishName: utils.StringConvert(value[5]),
			Code:        utils.StringConvert(value[6]),
			Lat:         utils.StringConvert(value[7]),
			Lng:         utils.StringConvert(value[8]),
		})
	}
	err := m.GetConn().CreateInBatches(cities, 100).Error
	if err != nil {
		logrus.Error("seed cities error: ", err)
	}
}
