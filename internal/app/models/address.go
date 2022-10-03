package models

type Address struct {
	ID         uint64   `gorm:"primary_key;auto_increment" json:"id"`
	UserID     uint64   `gorm:"not null" json:"user_id"`
	User       User     `gorm:"foreignKey:user_id" json:"user"`
	Title      string   `gorm:"not null" json:"title"`
	ProvinceID uint64   `gorm:"not null" json:"province_id"`
	Province   Province `gorm:"foreignKey:province_id" json:"province"`
	CityID     uint64   `gorm:"not null" json:"city_id"`
	City       City     `gorm:"foreignKey:city_id" json:"city"`
	Address    string   `gorm:"not null" json:"address"`
	PostalCode string   `gorm:"not null" json:"postal_code"`
	Phone      string   `gorm:"not null" json:"phone"`
	FullName   string   `gorm:"not null" json:"full_name"`
	Lat        string   `gorm:"not null" json:"lat"`
	Long       string   `gorm:"not null" json:"long"`
}
