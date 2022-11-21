package DTOs

type CreateAddress struct {
	Title      string `form:"title" json:"title" validate:"required" fake:"{street_name}"`
	ProvinceID uint64 `form:"province_id" json:"province_id" validate:"required,numeric" fake:"{number:1,1}"`
	CityID     uint64 `form:"city_id" json:"city_id" validate:"required,numeric" fake:"{number:1,1}"`
	Address    string `form:"address" json:"address" validate:"required" fake:"{address}"`
	PostalCode string `form:"postal_code" json:"postal_code" validate:"required,numeric,min=10,max=10" fake:"{zip}"`
	Mobile     string `form:"mobile" json:"mobile" validate:"required,numeric,startswith" fake:"{phone}"`
	FullName   string `form:"full_name" json:"full_name" validate:"required" fake:"{name}"`
	Lat        string `form:"lat" json:"lat" fake:"{latitude}"`
	Long       string `form:"long" json:"long" fake:"{longitude}"`
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

type IndexAddress struct {
	Index
}
