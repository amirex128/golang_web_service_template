package models

type Province struct {
	ID          int    `json:"id"`
	PersianName string `json:"persian_name"`
	EnglishName string `json:"english_name"`
	Cod         string `json:"COD"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
type ProductProvince struct {
	ProductID  string `json:"product_id"`
	ProvinceID string `json:"province_id"`
}
