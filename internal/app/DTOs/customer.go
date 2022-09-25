package DTOs

type LoginCustomer struct {
	Mobile string `json:"mobile" validate:"required,min=11,max=11,startswith=09"`
}
type VerifyCustomer struct {
	Mobile     string `json:"mobile" validate:"required,min=11,max=11,startswith=09"`
	VerifyCode string `json:"verify_code" validate:"required"`
}
type UpdateCustomer struct {
	Mobile     string `json:"mobile" validate:"required,min=11,max=11,startswith=09"`
	VerifyCode string `json:"verify_code" validate:"required,min=4,max=20"`
	FullName   string `json:"full_name" validate:"required"`
	ProvinceID uint32 `json:"province_id" validate:"required,numeric"`
	CityID     uint32 `json:"city_id" validate:"required,numeric"`
	Address    string `json:"address" validate:"required"`
	PostalCode uint64 `json:"postal_code" validate:"required,startswith=9"`
}
type CreateCustomer struct {
	Mobile     string `json:"mobile" validate:"required,min=11,max=11,startswith=09"`
	VerifyCode string `json:"verify_code" validate:"required,min=4,max=20"`
	FullName   string `json:"full_name" validate:"required"`
	ProvinceID uint32 `json:"province_id" validate:"required,numeric"`
	CityID     uint32 `json:"city_id" validate:"required,numeric"`
	Address    string `json:"address" validate:"required"`
	PostalCode uint64 `json:"postal_code" validate:"required,startswith=9"`
}
