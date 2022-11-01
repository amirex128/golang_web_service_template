package models

import (
	"backend/internal/app/utils"
	"github.com/sirupsen/logrus"
)

type Province struct {
	ID          int    `gorm:"primary_key;auto_increment" json:"id"`
	PersianName string `json:"persian_name"`
	EnglishName string `json:"english_name"`
	Cod         string `json:"COD"` // کد سیستم پست

}
type ProductProvince struct {
	ProductID  string `json:"product_id"`
	ProvinceID string `json:"province_id"`
}

func initProvince(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Province{})
	manager.GetConn().AutoMigrate(&ProductProvince{})
	provinces := utils.ReadCsvFile("../../csv/provinces.csv")
	manager.CreateAllProvinces(provinces)

}
func (m *MysqlManager) CreateAllProvinces(files [][]string) {
	province := make([]Province, 0)
	for i := range files {
		value := files[i]
		province = append(province, Province{
			ID:          utils.StringToInt(value[0]),
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
