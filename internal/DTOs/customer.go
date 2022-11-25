package DTOs

type RequestCreateLoginCustomer struct {
	//شماره همراه
	Mobile string `form:"mobile" json:"mobile" validate:"required,numeric,startswith=09,min=11,max=11" example:"09024809750"`
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validate:"required,numeric" example:"1"`
}

type CreateUpdateCustomer struct {
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validate:"required" example:"1"`
	//شماره همراه
	Mobile string `form:"mobile" json:"mobile" validate:"omitempty,min=11,max=11,startswith=09" example:"09024809750"`
	//کد تائید
	VerifyCode string `form:"verify_code" json:"verify_code" validate:"omitempty,min=4,max=4" example:"1234"`
	//نام کامل
	FullName string `form:"full_name" json:"full_name" validate:"omitempty" example:"نام"`
	//شناسه استان
	ProvinceID uint32 `form:"province_id" json:"province_id" validate:"omitempty,numeric" example:"1"`
	//شناسه شهر
	CityID uint32 `form:"city_id" json:"city_id" validate:"omitempty,numeric" example:"1"`
	//آدرس
	Address string `form:"address" json:"address" validate:"omitempty" example:"آدرس"`
	//کد پستی
	PostalCode    string `form:"postal_code" json:"postal_code" validate:"omitempty,startswith=9" example:"9111111111"`
	LastSendSMSAt string `form:"-" json:"-"`
}

type IndexOrderCustomer struct {
	//موبایل
	Mobile string `form:"mobile" json:"mobile" validate:"omitempty,min=11,max=11,startswith=09"  example:"09024809750"`
	//کد تائید دریافتی که جکم توکن را برای مشتریان دارند و تمامی وبسرویس های آن ها با استفاده از این کد کار خواهند کرد
	VerifyCode string `form:"verify_code" json:"verify_code" validate:"omitempty,min=4,max=4" example:"1234"`
}

type IndexCustomer struct {
	Index
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" example:"1"`
}
