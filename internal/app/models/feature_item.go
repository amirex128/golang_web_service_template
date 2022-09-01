package models

import (
	"backend/internal/app/helpers"
	"database/sql"
	"encoding/gob"
	"github.com/sirupsen/logrus"
	"io"
)

type FeatureItem struct {
	ID            int            `json:"id"`
	FeatureItemID int            `json:"feature_group_id"`
	Title         string         `json:"title"`
	Type          string         `json:"type" sql:"type:ENUM('label','color','boolean','other')"`              //      => تعیین نوع مجاز پاسخ برای این خصوصیت : پاسخ متنی ساده - پاسخ از نوع رنگ - پاسخ بله یا خیر - پاسخ چند جوابه
	Actions       string         `json:"actions" sql:"type:ENUM('search','top_feature','short_desc','point')"` //     => خاص تر کردن یک خصوصیت و تعیین اینکه این خصوصیت عملکرد و استفاده برای جستجو دارد و خصوصیت برجسته است یا خصوصیت برای استفاده در توضیحات مختصر است و یا خصوصیت برای نظردهی است
	Active        byte           `json:"active"`
	Icon          sql.NullString `json:"icon"`
	InputType     string         `json:"input_type" sql:"type:ENUM('number','text','float','all')"` //     => نوع ورودی های مجاز برای مقدار ویژگی ها
	Sort          uint           `json:"sort"`
}
type FeatureItemArr []FeatureItem

func (s FeatureItemArr) Len() int {
	return len(s)
}
func (s FeatureItemArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s FeatureItemArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *FeatureItem) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *FeatureItem) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}

func initFeatureItem(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&FeatureItem{})
	initFeatureItemProduct(manager)
	featureItems := helpers.ReadCsvFile("../../csv/feature_items.csv")
	manager.CreateAllFeatureItems(featureItems)
}

func (m *MysqlManager) CreateAllFeatureItems(files [][]string) {
	featureItems := make([]FeatureItem, 0)
	for i := range files {
		value := files[i]
		featureItems = append(featureItems, FeatureItem{
			ID:            int32Convert(value[0]),
			FeatureItemID: int32Convert(value[1]),
			Title:         value[2],
			Type:          value[3],
			Actions:       value[4],
			Active:        activeConvert(value[5]),
			Icon:          stringConvert(value[6]),
			InputType:     value[7],
			Sort:          uintConvert(value[8]),
		})
	}
	err := m.GetConn().CreateInBatches(featureItems, 100).Error
	if err != nil {
		logrus.Error("seed featureItems error: ", err)
	}
}
