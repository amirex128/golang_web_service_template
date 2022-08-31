package models

import (
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

func (m *MysqlManager) CreateAllFeatureGroups(files [][]string) {
	featureGroups := make([]FeatureGroup, 0)
	for i := range files {
		value := files[i]
		featureGroups = append(featureGroups, FeatureGroup{
			ID:          intConvert(value[0]),
			CategoryID:  intConvert(value[1]),
			Title:       value[2],
			Icon:        stringConvert(value[3]),
			Sort:        uintConvert(value[4]),
			Description: stringConvert(value[5]),
		})
	}
	err := m.GetConn().CreateInBatches(featureGroups, 100).Error
	if err != nil {
		logrus.Error("seed featureGroups error: ", err)
	}
}
