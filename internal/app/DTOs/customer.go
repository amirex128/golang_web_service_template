package DTOs

type LoginCustomer struct {
	Mobile string `form:"mobile" json:"mobile" validate:"required,min=11,max=11,startswith=09"`
}
type VerifyCustomer struct {
	Mobile     string `form:"mobile" json:"mobile" validate:"required,min=11,max=11,startswith=09"`
	VerifyCode string `form:"verify_code" json:"verify_code" validate:"required"`
}
type UpdateCustomer struct {
	Mobile     string `form:"mobile" json:"mobile" validate:"required,min=11,max=11,startswith=09"`
	VerifyCode string `form:"verify_code" json:"verify_code" validate:"required,min=4,max=20"`
	FullName   string `form:"full_name" json:"full_name" validate:"required"`
	ProvinceID uint32 `form:"province_id" json:"province_id" validate:"required,numeric"`
	CityID     uint32 `form:"city_id" json:"city_id" validate:"required,numeric"`
	Address    string `form:"address" json:"address" validate:"required"`
	PostalCode uint64 `form:"postal_code" json:"postal_code" validate:"required,startswith=9"`
}
type CreateCustomer struct {
	Mobile     string `form:"mobile" json:"mobile" validate:"required,min=11,max=11,startswith=09"`
	VerifyCode string `form:"verify_code" json:"verify_code" validate:"required,min=4,max=20"`
	FullName   string `form:"full_name" json:"full_name" validate:"required"`
	ProvinceID uint32 `form:"province_id" json:"province_id" validate:"required,numeric"`
	CityID     uint32 `form:"city_id" json:"city_id" validate:"required,numeric"`
	Address    string `form:"address" json:"address" validate:"required"`
	PostalCode uint64 `form:"postal_code" json:"postal_code" validate:"required,startswith=9"`
}
