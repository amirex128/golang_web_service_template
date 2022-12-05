package models

import (
	"database/sql"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/sirupsen/logrus"
)

type FeatureItem struct {
	ID            int            `gorm:"primary_key;auto_increment" json:"id"`
	FeatureItemID int            `json:"feature_group_id"`
	Title         string         `json:"title"`
	Type          string         `json:"type" sql:"type:ENUM('label','color','boolean','other')"`              //      => تعیین نوع مجاز پاسخ برای این خصوصیت : پاسخ متنی ساده - پاسخ از نوع رنگ - پاسخ بله یا خیر - پاسخ چند جوابه
	Actions       string         `json:"actions" sql:"type:ENUM('search','top_feature','short_desc','point')"` //     => خاص تر کردن یک خصوصیت و تعیین اینکه این خصوصیت عملکرد و استفاده برای جستجو دارد و خصوصیت برجسته است یا خصوصیت برای استفاده در توضیحات مختصر است و یا خصوصیت برای نظردهی است
	Active        byte           `json:"active"`
	Icon          sql.NullString `json:"icon"`
	InputType     string         `json:"input_type" sql:"type:ENUM('number','text','float','all')"` //     => نوع ورودی های مجاز برای مقدار ویژگی ها
	Sort          uint           `json:"sort"`
}

func initFeatureItem(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&FeatureItem{}) {
		manager.GetConn().Migrator().CreateTable(&FeatureItem{})
		initFeatureItemProduct(manager)
		featureItems := utils.ReadCsvFile("./csv/feature_items.csv")
		manager.CreateAllFeatureItems(featureItems)

	}
}

func (m *MysqlManager) CreateAllFeatureItems(files [][]string) {
	featureItems := make([]FeatureItem, 0)
	for i := range files {
		value := files[i]
		featureItems = append(featureItems, FeatureItem{
			ID:            utils.StringToInt(value[0]),
			FeatureItemID: utils.StringToInt(value[1]),
			Title:         value[2],
			Type:          value[3],
			Actions:       value[4],
			Active:        utils.ActiveConvert(value[5]),
			Icon:          utils.StringConvert(value[6]),
			InputType:     value[7],
			Sort:          utils.StringToUint(value[8]),
		})
	}
	err := m.GetConn().CreateInBatches(featureItems, 100).Error
	if err != nil {
		logrus.Error("seed featureItems error: ", err)
	}
}
