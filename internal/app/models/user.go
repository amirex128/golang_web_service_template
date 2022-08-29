package models

type User struct {
	ID          int64  `json:"id"`
	Gender      string `json:"gender"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	ShopName    string `json:"shop_name"`
	GuildID     string `json:"guild_id"`
	ProvinceID  string `json:"province_id"`
	CityID      string `json:"city_id"`
	Lat         string `json:"lat"`
	Long        string `json:"long"`
	Logo        string `json:"logo"`
	Address     string `json:"address"`
	Mobile      string `json:"mobile"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Description string `json:"desc"`
	ExpireAt    string `json:"expire_at"`
	Status      string `json:"status"`
	PostalCode  string `json:"postal_code"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}
