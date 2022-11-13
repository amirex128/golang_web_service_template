package DTOs

type CreateAddress struct {
	Title      string `form:"title" json:"title" validate:"required"`
	ProvinceID uint64 `form:"province_id" json:"province_id" validate:"required,numeric"`
	CityID     uint64 `form:"city_id" json:"city_id" validate:"required,numeric"`
	Address    string `form:"address" json:"address" validate:"required"`
	PostalCode string `form:"postal_code" json:"postal_code" validate:"required,numeric,min=10,max=10"`
	Mobile     string `form:"mobile" json:"mobile" validate:"required,numeric,startswith"`
	FullName   string `form:"full_name" json:"full_name" validate:"required"`
	Lat        string `form:"lat" json:"lat"`
	Long       string `form:"long" json:"long"`
}

type UpdateAddress struct {
	ID         uint64 `form:"id" json:"id" validate:"required"`
	Title      string `form:"title" json:"title" validate:"omitempty"`
	ProvinceID uint64 `form:"province_id" json:"province_id" validate:"omitempty,numeric"`
	CityID     uint64 `form:"city_id" json:"city_id" validate:"omitempty,numeric"`
	Address    string `form:"address" json:"address" validate:"omitempty"`
	PostalCode string `form:"postal_code" json:"postal_code" validate:"omitempty,numeric,min=10,max=10"`
	Mobile     string `form:"mobile" json:"mobile" validate:"omitempty,numeric,startswith"`
	FullName   string `form:"full_name" json:"full_name" validate:"omitempty"`
	Lat        string `form:"lat" json:"lat"`
	Long       string `form:"long" json:"long"`
}
