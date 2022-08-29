package models

type City struct {
	ID          int    `json:"id"`
	ProvinceID  string `json:"province_id"`
	PersianName string `json:"persian_name"`
	EnglishName string `json:"english_name"`
	Code        string `json:"code"`
	Lat         string `json:"lat"`
	Lng         string `json:"lng"`
}

type CityProduct struct {
	ProductID int64 `json:"product_id"`
	CityID    int   `json:"city_id"`
}
