package models

type Manufacturer struct {
	ID               int    `json:"id"`
	CountryName      string `json:"country_name"`
	ShortDescription string `json:"short_description"`
	Logo             string `json:"logo"`
	PersianName      string `json:"persian_name"`
	EnglishName      string `json:"english_name"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
