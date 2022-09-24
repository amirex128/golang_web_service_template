package models

import (
	"backend/internal/app/utils"
	"database/sql"
	"encoding/gob"
	"github.com/sirupsen/logrus"
	"io"
)

type FeatureGroup struct {
	ID          int            `json:"id"`
	CategoryID  int            `json:"category_id"`
	Title       string         `json:"title"`
	Icon        sql.NullString `json:"icon"`
	Sort        uint           `json:"sort"`
	Description sql.NullString `json:"description"`
}

type FeatureGroupArr []FeatureGroup

func (s FeatureGroupArr) Len() int {
	return len(s)
}
func (s FeatureGroupArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s FeatureGroupArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *FeatureGroup) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *FeatureGroup) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}

func initFeatureGroup(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&FeatureGroup{})
	featureGroups := utils.ReadCsvFile("../../csv/feature_groups.csv")
	manager.CreateAllFeatureGroups(featureGroups)
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
