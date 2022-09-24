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
	VerifyCode uint16 `json:"verify_code" validate:"required,min=1000,max=9999"`
	FullName   string `json:"full_name" validate:"required"`
	ProvinceID uint32 `json:"province_id" validate:"required,numeric"`
	CityID     uint32 `json:"city_id" validate:"required,numeric"`
	Address    string `json:"address" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required,len=10,numeric"`
}
