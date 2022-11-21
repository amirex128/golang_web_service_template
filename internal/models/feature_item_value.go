package models

import (
	utils2 "github.com/amirex128/selloora_backend/internal/utils"
	"github.com/sirupsen/logrus"
)

type FeatureItemValue struct {
	ID            int    `gorm:"primary_key;auto_increment" json:"id"`
	FeatureItemID int    `json:"feature_item_id"`
	Value         string `json:"value"`
}

func initFeatureItemValue(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&FeatureItemValue{})
	featureItemValues := utils2.ReadCsvFile("./csv/feature_item_values.csv")
	manager.CreateAllFeatureItemValues(featureItemValues)
}
func (m *MysqlManager) CreateAllFeatureItemValues(files [][]string) {
	featureItemValue := make([]FeatureItemValue, 0)
	for i := range files {
		value := files[i]
		featureItemValue = append(featureItemValue, FeatureItemValue{
			ID:            utils2.StringToInt(value[0]),
			FeatureItemID: utils2.StringToInt(value[1]),
			Value:         value[2],
		})
	}
	err := m.GetConn().CreateInBatches(featureItemValue, 100).Error
	if err != nil {
		logrus.Error("seed featureItemValue error: ", err)
	}
}
