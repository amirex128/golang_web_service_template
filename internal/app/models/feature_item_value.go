package models

import (
	"encoding/gob"
	"github.com/sirupsen/logrus"
	"io"
)

type FeatureItemValue struct {
	ID            int    `json:"id"`
	FeatureItemID int    `json:"feature_item_id"`
	Value         string `json:"value"`
}
type FeatureItemValueArr []FeatureItemValue

func (s FeatureItemValueArr) Len() int {
	return len(s)
}
func (s FeatureItemValueArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s FeatureItemValueArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *FeatureItemValue) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *FeatureItemValue) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}

func (m *MysqlManager) CreateAllFeatureItemValues(files [][]string) {
	featureItemValue := make([]FeatureItemValue, 0)
	for i := range files {
		value := files[i]
		featureItemValue = append(featureItemValue, FeatureItemValue{
			ID:            intConvert(value[0]),
			FeatureItemID: intConvert(value[1]),
			Value:         value[2],
		})
	}
	err := m.GetConn().CreateInBatches(featureItemValue, 100).Error
	if err != nil {
		logrus.Error("seed featureItemValue error: ", err)
	}
}
