package DTOs

type CreateShop struct {
	//نام فروشگاه
	Name string `form:"name" json:"name" validate:"required" example:"فروشگاه امیر" fake:"{name}"`
	//تصویر لوگو فروشگاه
	GalleryID uint64 `form:"gallery_id" json:"gallery_id" validate:"omitempty" example:"1" fake:"{number:1,100}"`
	//نوع فروشگاه که میتواند به صورت اینستاگرامی باشد یا وبسایتی باشد یا روبیکا یا ربات تلگرام
	Type string `form:"type" json:"type" validate:"required,oneof=instagram telegram website rubika" example:"instagram" enums:"instagram,telegram,website,rubika" fake:"{custom_enum:instagram,telegram,website,rubika}"`
	//آدرس شبکه اجتماعیی که میخواهید متصل کنید به فروشگاه برای دریافت محصولات از ان
	SocialAddress string `form:"social_address" json:"social_address" validate:"required" example:"amirex_dev" fake:"{username}"`
	//توضیحات فروشگاه
	Description string `form:"description" json:"description" validate:"omitempty" example:"توضیحات فروشگاه" fake:"{sentence:10}"`
	//شماره تلفن برای نمایش در فروشگاه برای پشتیبانی
	Phone string `form:"phone" json:"phone" validate:"omitempty" example:"05136643278" fake:"{phone}"`
	//شماره موبایل برای نمایش در فروشگاه برای پشتیبانی
	Mobile string `form:"mobile" json:"mobile" validate:"omitempty,numeric,startswith=09" example:"09024809750" fake:"{phone}"`
	//آیدی تلگرام برای نمایش در فروشگاه برای پشتیبانی
	TelegramID string `form:"telegram_id" json:"telegram_id" validate:"omitempty" example:"amirex128" fake:"{username}"`
	//آیدی اینستاگرام برای نمایش در فروشگاه برای پشتیبانی
	InstagramID string `form:"instagram_id" json:"instagram_id" validate:"omitempty" example:"amirex_dev" fake:"{username}"`
	//آیدی واتس آپ برای نمایش در فروشگاه برای پشتیبانی
	WhatsappID string `form:"whatsapp_id" json:"whatsapp_id" validate:"omitempty" example:"amirex128" fake:"{username}"`
	//ایمیل برای نمایش در فروشگاه برای پشتیبانی
	Email string `form:"email" json:"email" validate:"omitempty,email" example:"amirex128@gmail.com" fake:"{email}"`
	//وبسایت شخصی برای نمایش در فروشگاه برای پشتیبانی
	Website string `form:"website" json:"website" validate:"omitempty" example:"https://amirshirdel.ir" fake:"{url}"`
	//هزینه ارسال تومان
	SendPrice float32 `form:"send_price" json:"send_price" validate:"omitempty" example:"20000" fake:"{number:10000,100000}"`
}

type UpdateShop struct {
	//شناسه فروشگاه برای ویرایش
	ID uint64 `form:"id" json:"id" validate:"required" example:"1"`
	//نام فروشگاه
	Name string `form:"name" json:"name" validate:"omitempty" example:"فروشگاه امیر"`
	//تصویر لوگو فروشگاه
	GalleryID uint64 `form:"gallery_id" json:"gallery_id" validate:"omitempty" example:"1"`
	//نوع فروشگاه که میتواند به صورت اینستاگرامی باشد یا وبسایتی باشد یا روبیکا یا ربات تلگرام
	Type string `form:"type" json:"type" validate:"omitempty,oneof=instagram telegram website rubika" example:"instagram" enums:"instagram,telegram,website,rubika"`
	//آدرس شبکه اجتماعیی که میخواهید متصل کنید به فروشگاه برای دریافت محصولات از ان
	SocialAddress string `form:"social_address" json:"social_address" validate:"omitempty" example:"amirex_dev"`
	//توضیحات فروشگاه
	Description string `form:"description" json:"description" validate:"omitempty" example:"توضیحات فروشگاه"`
	//شماره تلفن برای نمایش در فروشگاه برای پشتیبانی
	Phone string `form:"phone" json:"phone" validate:"omitempty" example:"05136643278"`
	//شماره موبایل برای نمایش در فروشگاه برای پشتیبانی
	Mobile string `form:"mobile" json:"mobile" validate:"omitempty,numeric,startswith=09" example:"09024809750"`
	//آیدی تلگرام برای نمایش در فروشگاه برای پشتیبانی
	TelegramID string `form:"telegram_id" json:"telegram_id" validate:"omitempty" example:"amirex128"`
	//آیدی اینستاگرام برای نمایش در فروشگاه برای پشتیبانی
	InstagramID string `form:"instagram_id" json:"instagram_id" validate:"omitempty" example:"amirex_dev"`
	//آیدی واتس آپ برای نمایش در فروشگاه برای پشتیبانی
	WhatsappID string `form:"whatsapp_id" json:"whatsapp_id" validate:"omitempty" example:"amirex128"`
	//ایمیل برای نمایش در فروشگاه برای پشتیبانی
	Email string `form:"email" json:"email" validate:"omitempty,email" example:"amirex128@gmail.com"`
	//وبسایت شخصی برای نمایش در فروشگاه برای پشتیبانی
	Website string `form:"website" json:"website" validate:"omitempty" example:"https://amirshirdel.ir"`
	//هزینه ارسال تومان
	SendPrice float32 `form:"send_price" json:"send_price" validate:"omitempty" example:"25000"`
	//شناسه قالب
	ThemeID uint64 `form:"theme_id" json:"theme_id" validate:"omitempty" example:"1"`
}

type DeleteShop struct {
	//شناسه فروشگاه برای حذف
	NewShopID uint64 `form:"new_shop_id" json:"new_shop_id" validate:"omitempty" example:"1"`
	//نوع حذف فروشگاه که محصولات فروشگاه هم حذف شود یا به فروشگاه دیگری انتقال پیدا کند
	ProductBehave string `form:"product_behave" json:"product_behave" validate:"oneof=move delete_product,required" example:"move" enums:"move,delete_product"`
}
type CheckSocial struct {
	//آدرس شبکه اجتماعی
	SocialAddress string `form:"social_address" json:"social_address" validate:"required" example:"amirex_dev"`
	//نوع شبکه اجتماعی
	Type string `form:"type" json:"type" validate:"required" example:"instagram" enums:"instagram,telegram,website,rubika"`
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validate:"required,numeric" example:"1"`
}

type SendPriceShop struct {
	//هزینه ارسال تومان
	SendPrice float32 `form:"send_price" json:"send_price" validate:"required,numeric" example:"30000"`
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validate:"required,numeric" example:"1"`
}
type SelectThemeShop struct {
	// شناسه قالب
	ThemeID uint64 `form:"theme_id" json:"theme_id" validate:"required,numeric" example:"1"`
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validate:"required,numeric" example:"1"`
}
type IndexShop struct {
	Index
}
