package models

type Gallery struct {
	ID        int    `json:"id"`
	ProductID int64  `json:"product_id"`
	Path      string `json:"path"`
	Caption   string `json:"caption"`
	Type      string `json:"type"` // image video cover
	Sort      uint   `json:"sort"`
	IsMain    byte   `json:"is_main"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
