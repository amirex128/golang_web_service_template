package models

import (
	"database/sql"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/sirupsen/logrus"
)

type FeatureGroup struct {
	ID          int            `gorm:"primary_key;auto_increment" json:"id"`
	CategoryID  int            `json:"category_id"`
	Title       string         `json:"title"`
	Icon        sql.NullString `json:"icon"`
	Sort        uint           `json:"sort"`
	Description sql.NullString `json:"description"`
}

func initFeatureGroup(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&FeatureGroup{}) {
		manager.GetConn().AutoMigrate(&FeatureGroup{})
		featureGroups := utils.ReadCsvFile("./csv/feature_groups.csv")
		manager.CreateAllFeatureGroups(featureGroups)

	}
}
func (m *MysqlManager) CreateAllFeatureGroups(files [][]string) {
	featureGroups := make([]FeatureGroup, 0)
	for i := range files {
		value := files[i]
		featureGroups = append(featureGroups, FeatureGroup{
			ID:          utils.StringToInt(value[0]),
			CategoryID:  utils.StringToInt(value[1]),
			Title:       value[2],
			Icon:        utils.StringConvert(value[3]),
			Sort:        utils.StringToUint(value[4]),
			Description: utils.StringConvert(value[5]),
		})
	}
	err := m.GetConn().CreateInBatches(featureGroups, 100).Error
	if err != nil {
		logrus.Error("seed featureGroups error: ", err)
	}
}
