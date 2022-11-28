package DTOs

type CreateAddress struct {
	//عنوان
	Title string `form:"title" json:"title" validate:"required" example:"عنوان" fake:"{street_name}"`
	//شناسه استان
	ProvinceID uint64 `form:"province_id" json:"province_id" validate:"required,numeric" example:"1" fake:"{custom_uint64:1}"`
	//شناسه شهر
	CityID uint64 `form:"city_id" json:"city_id" validate:"required,numeric" example:"1" fake:"{custom_uint64:1}"`
	//ادرس کامل
	Address string `form:"address" json:"address" validate:"required" example:"ادرس کامل" fake:"{address}"`
	//کد پستی
	PostalCode string `form:"postal_code" json:"postal_code" validate:"required,numeric,min=10,max=10" example:"9111111111" fake:"{zip}"`
	//موبایل
	Mobile string `form:"mobile" json:"mobile" validate:"required,numeric,startswith,min=11,max=11" example:"09024809750" fake:"{phone}"`
	//نام گیرنده
	FullName string `form:"full_name" json:"full_name" validate:"required" example:"نام گیرنده" fake:"{name}"`
	//عرض جغذافیایی
	Lat string `form:"lat" json:"lat" example:"" example:"35.123456" fake:"{latitude}"`
	//طول جقرافیایی
	Long string `form:"long" json:"long" example:"" example:"35.123456" fake:"{longitude}"`
}

type UpdateAddress struct {
	//شناسه آدرس
	ID uint64 `form:"id" json:"id" validate:"required" example:"1" `
	//عنوان
	Title string `form:"title" json:"title" validate:"omitempty" example:"عنوان" `
	//شناسه استان
	ProvinceID uint64 `form:"province_id" json:"province_id" validate:"omitempty,numeric" example:"1" `
	//شناسه شهر
	CityID uint64 `form:"city_id" json:"city_id" validate:"omitempty,numeric" example:"1"`
	//ادرس کامل
	Address string `form:"address" json:"address" validate:"omitempty" example:"ادرس کامل" `
	//کد پستی
	PostalCode string `form:"postal_code" json:"postal_code" validate:"omitempty,numeric,min=10,max=10" example:"1111111111" `
	//موبایل
	Mobile string `form:"mobile" json:"mobile" validate:"omitempty,numeric,startswith" example:"09024809750" `
	//نام گیرنده
	FullName string `form:"full_name" json:"full_name" validate:"omitempty" example:"نام گیرنده" `
	//عرض جغذافیایی
	Lat string `form:"lat" json:"lat" example:"35.123456" `
	//طول جقرافیایی
	Long string `form:"long" json:"long" example:"35.123456" `
}

type IndexAddress struct {
	Index
}
