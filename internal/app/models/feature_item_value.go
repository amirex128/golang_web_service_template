package models

type FeatureItemValue struct {
	ID            int    `json:"id"`
	FeatureItemID int    `json:"feature_item_id"`
	Value         string `json:"value"`
}
