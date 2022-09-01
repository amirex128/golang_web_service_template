package models

import (
	"backend/internal/app/helpers"
	"encoding/gob"
	"github.com/sirupsen/logrus"
	"io"
)

type Province struct {
	ID          int    `json:"id"`
	PersianName string `json:"persian_name"`
	EnglishName string `json:"english_name"`
	Cod         string `json:"COD"` // کد سیستم پست

}
type ProductProvince struct {
	ProductID  string `json:"product_id"`
	ProvinceID string `json:"province_id"`
}
type ProvinceArr []Province

func (s ProvinceArr) Len() int {
	return len(s)
}
func (s ProvinceArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ProvinceArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Province) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Province) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initProvince(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Province{})
	manager.GetConn().AutoMigrate(&ProductProvince{})
	provinces := helpers.ReadCsvFile("../../csv/provinces.csv")
	manager.CreateAllProvinces(provinces)

}
func (m *MysqlManager) CreateAllProvinces(files [][]string) {
	province := make([]Province, 0)
	for i := range files {
		value := files[i]
		province = append(province, Province{
			ID:          helpers.Int32Convert(value[0]),
			PersianName: value[1],
			EnglishName: value[2],
			Cod:         value[3],
		})
	}
	err := m.GetConn().CreateInBatches(province, 100).Error
	if err != nil {
		logrus.Error("seed province error: ", err)
	}
}
