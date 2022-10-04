package DTOs

type RequestCreateLoginCustomer struct {
	Mobile string `form:"mobile" json:"mobile" validate:"required,numeric,startswith=09,min=11,max=11"`
	ShopID uint64 `form:"shop_id" json:"shop_id" validate:"required,numeric"`
}

type CreateUpdateCustomer struct {
	ShopID        uint64 `form:"shop_id" json:"shop_id" validate:"required"`
	Mobile        string `form:"mobile" json:"mobile" validate:"omitempty,min=11,max=11,startswith=09"`
	VerifyCode    string `form:"verify_code" json:"verify_code" validate:"omitempty,min=4,max=4"`
	FullName      string `form:"full_name" json:"full_name" validate:"omitempty"`
	ProvinceID    uint32 `form:"province_id" json:"province_id" validate:"omitempty,numeric"`
	CityID        uint32 `form:"city_id" json:"city_id" validate:"omitempty,numeric"`
	Address       string `form:"address" json:"address" validate:"omitempty"`
	PostalCode    uint64 `form:"postal_code" json:"postal_code" validate:"omitempty,startswith=9"`
	LastSendSMSAt string
}
