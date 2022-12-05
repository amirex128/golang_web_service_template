package models

type FeatureItemProduct struct {
	ID                 int    `gorm:"primary_key;auto_increment" json:"id"`
	ProductID          int    `json:"product_id"`
	FeatureItemID      int    `json:"feature_item_id"`
	FeatureItemValueID int    `json:"feature_item_value_id"`
	FeatureGroupID     int    `json:"feature_group_id"`
	Value              string `json:"value"`
}

func initFeatureItemProduct(manager *MysqlManager) {
	manager.GetConn().Migrator().CreateTable(&FeatureItemProduct{})
}
