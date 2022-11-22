package DTOs

type CreateDomain struct {
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validate:"required,numeric" example:"1"`
	//نام دامنه که باید یا به صورت
	//example.com
	//یا به صورت ساب دامنه سلورا
	//example.selloora.com
	Name string `form:"name" json:"name" validate:"required" example:"example.com"`
	//نوع دامنه که یا ساب دامنه میباشد یا دامنه اختصاصی
	Type string `form:"type" json:"type" validate:"required" example:"domain" enums:"domain,subdomain"`
}

type IndexDomain struct {
	Index
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validate:"omitempty,numeric" example:"1"`
}