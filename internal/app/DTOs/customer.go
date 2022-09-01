package DTOs

type RequestCustomer struct {
	Mobile string `json:"mobile" validate:"required,min=11,max=11,startswith=09"`
}
type VerifyCustomer struct {
	Mobile     string `json:"mobile" validate:"required,min=11,max=11,startswith=09"`
	VerifyCode uint16 `json:"verify_code" validate:"required,min=1000,max=9999"`
}
type UpdateCustomer struct {
	VerifyCode uint16 `json:"verify_code" validate:"required,min=1000,max=9999"`
	FullName   string `json:"full_name" validate:"required"`
	ProvinceID uint64 `json:"province_id" validate:"required,numeric"`
	CityID     uint64 `json:"city_id" validate:"required,numeric"`
	Address    string `json:"address" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required,len=10,numeric"`
}
