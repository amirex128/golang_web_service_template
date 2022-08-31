package models

import (
	"encoding/gob"
	"io"
)

type FeatureItemProduct struct {
	ID                 int    `json:"id"`
	ProductID          int    `json:"product_id"`
	FeatureItemID      int    `json:"feature_item_id"`
	FeatureItemValueID int    `json:"feature_item_value_id"`
	FeatureGroupID     int    `json:"feature_group_id"`
	Value              string `json:"value"`
}
type FeatureItemProductArr []FeatureItemProduct

func (s FeatureItemProductArr) Len() int {
	return len(s)
}
func (s FeatureItemProductArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s FeatureItemProductArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *FeatureItemProduct) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *FeatureItemProduct) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
