package DTOs

type IndexProduct struct {
	Index
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" example:"0"`
}

type CreateProduct struct {
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" validate:"required" example:"1" fake:"{number:1,100}"`
	//برند محصول
	Manufacturer string `form:"manufacturer_id" validate:"omitempty" example:"سامسونگ" fake:"{company}"`
	//توضیحات محصول
	Description string `form:"description" validate:"omitempty" example:"توضیحات محصول" fake:"{lorem_ipsum_sentence:2}"`
	//نام محصول
	Name string `form:"name" validate:"required,min=3,max=100" example:"گوشی موبایل گلگسی نوت ۱۰" fake:"{lorem_ipsum_sentence:1}"`
	//تعداد موجودی
	Quantity uint32 `form:"quantity" validate:"required,numeric" example:"10" fake:"{number:1,100}"`
	//قیمت محصول
	Price float32 `form:"price" validate:"required,numeric" example:"1000000" fake:"{number:1000000,10000000}"`
	//تاریخ شروع فروش
	StartedAt string `form:"started_at" validate:"omitempty,datetime" example:"2020-01-01 00:00:00" fake:"{date}"`
	//تاریخ پایان فروش
	EndedAt string `form:"ended_at" validate:"omitempty,datetime" example:"2025-01-01 00:00:00" fake:"{date}"`
	//شناسه آشپن محصول
	OptionId uint32 `form:"option_id" validate:"numeric" example:"1" fake:"{number:1,100}"`
	//شناسه آیتم انتخاب شده از آپشن
	OptionItemID uint32 `form:"option_item_id" validate:"numeric" example:"1" fake:"{number:1,100}"`
	//شناسه دسته بندی
	CategoryID uint64 `form:"category_id" validate:"required,numeric" example:"1" fake:"{number:1,100}"`
	//شناسه تصاویر انتخاب شده برای این محصول
	GalleryIDs []uint64 `form:"gallery_ids" json:"gallery_ids" validate:"required,dive,numeric" example:"[1,2,3]" fakesize:"3"`
}

type UpdateProduct struct {
	//شناسه محصول برای ویرایش
	ID uint64 `form:"id" validate:"omitempty,numeric" example:"1"`
	//شناسه فروشگاه جهت انتقال یک محصول از یک فروشگاه به فروشگاه دیگر
	ShopID uint64 `form:"shop_id" json:"shop_id" validate:"omitempty,numeric" example:"2"`
	//برند محصول
	Manufacturer string `form:"manufacturer_id" validate:"omitempty,numeric" example:"سامسونگ"`
	//توضیحات محصول
	Description string `form:"description" validate:"omitempty,min=3,max=1000" example:"توضیحات محصول"`
	//نام محصول
	Name string `form:"name" validate:"omitempty,min=3,max=100" example:"گوشی موبایل گلگسی نوت ۱۰"`
	//تعداد موجودی
	Quantity uint32 `form:"quantity" validate:"omitempty,numeric" example:"10"`
	//قیمت محصول
	Price float32 `form:"price" validate:"omitempty,numeric" example:"1000000"`
	//وضعیت فعال یا غیر فعال بودن محصول
	Active bool `form:"active" validate:"omitempty,numeric" example:"true"`
	//تاریخ شروع فروش
	StartedAt string `form:"started_at" validate:"omitempty,datetime" example:"2020-01-01 00:00:00"`
	//تاریخ پایان فروش
	EndedAt string `form:"ended_at" validate:"omitempty,datetime" example:"2025-01-01 00:00:00"`
	//شناسه آیتم انتخاب شده از آپشن
	OptionId uint32 `form:"option_id" validate:"omitempty,numeric" example:"1"`
	//شناسه آیتم انتخاب شده از آپشن
	OptionItemID uint32 `form:"option_item_id" validate:"omitempty,numeric" example:"1"`
	//شناسه دسته بندی
	CategoryID uint64 `form:"category_id" validate:"required,numeric" example:"1"`
	//شناسه تصاویر انتخاب شده برای این محصول
	GalleryIDs []uint64 `form:"gallery_ids" json:"gallery_ids" validate:"omitempty,dive,numeric"`
}
