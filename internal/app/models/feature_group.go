package models

type FeatureGroup struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	Title       string `json:"title"`
	Icon        string `json:"icon"`
	Sort        uint   `json:"sort"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
