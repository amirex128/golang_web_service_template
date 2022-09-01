package models

import (
	"backend/internal/app/helpers"
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
	featureGroups := helpers.ReadCsvFile("../../csv/feature_groups.csv")
	manager.CreateAllFeatureGroups(featureGroups)
}
func (m *MysqlManager) CreateAllFeatureGroups(files [][]string) {
	featureGroups := make([]FeatureGroup, 0)
	for i := range files {
		value := files[i]
		featureGroups = append(featureGroups, FeatureGroup{
			ID:          helpers.Int32Convert(value[0]),
			CategoryID:  helpers.Int32Convert(value[1]),
			Title:       value[2],
			Icon:        helpers.StringConvert(value[3]),
			Sort:        helpers.UintConvert(value[4]),
			Description: helpers.StringConvert(value[5]),
		})
	}
	err := m.GetConn().CreateInBatches(featureGroups, 100).Error
	if err != nil {
		logrus.Error("seed featureGroups error: ", err)
	}
}
