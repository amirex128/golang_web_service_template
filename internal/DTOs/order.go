package DTOs

type OrderItem struct {
	//شناسه محصول
	ProductID uint64 `form:"product_id" json:"product_id" validate:"required,numeric" example:"1"`
	//شناسه آپشن محصول
	OptionID uint64 `form:"option_id" json:"option_id" validate:"numeric" example:"1"`
	//تعداد
	Count uint32 `form:"count" json:"count" validate:"required,numeric" example:"10"`
}

type CreateOrder struct {
	//شناسه مالک فروشگاه
	UserID uint64 `form:"user_id" json:"user_id" validate:"required,numeric" example:"1" `
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validate:"required,numeric" example:"1" `
	//لیست اقلام داخل سفارش
	OrderItems []OrderItem `form:"order_items" json:"order_items" validate:"required,dive" `
	//شناسه مشتری
	CustomerID uint64 `form:"customer_id" json:"customer_id" validate:"required,numeric" example:"1" `
	//کد تخفیف
	DiscountCode string `form:"discount_code" json:"discount_code" validate:"omitempty" example:"asdf"`
	//توضیحات سفارش
	Description string `form:"description" json:"description" example:"توضیحات ارسال سفارش"`
	//کد تائید سفارش که حکم تکن را برای مشاری دارد
	VerifyCode string `form:"verify_code" json:"verify_code" validate:"required,numeric,min=2,max=20" example:"1524"`
}

type SendOrder struct {
	//شناسه سفارش
	OrderID uint64 `form:"order_id" json:"order_id" validate:"required,numeric" example:"1"`
	//شرکت سرویس دهنده جهت ارسال سفارش
	Courier string `form:"courier" json:"courier" validate:"required,oneof=tipax post" example:"tipax"`
	//ابعاد بسته مورد نیاز جهت ارسال
	PackageSize string `form:"package_size" json:"package_size" validate:"required" example:"10x10x10" enums:"10x10x10,20x20x20,30x30x30"`
	//وزن کل سفارشات بر گرم
	Weight uint32 `form:"weight" json:"weight" validate:"required,numeric" example:"1000"`
	//ارزش تقریبی کل سفارشات جهت بیمه ارسال به تومان
	Value uint32 `form:"value" json:"value" validate:"required,numeric" example:"10000"`
	//شناسه آدرس مبدا ارسال سفارش
	AddressID uint64 `form:"address_id" json:"address_id" validate:"required,numeric" example:"1"`
}

type CalculateOrder struct {
	//ابعاد بسته مورد نیاز جهت ارسال
	PackageSize string `form:"package_size" json:"package_size" validate:"required" example:"10x10x10" enums:"10x10x10,20x20x20,30x30x30"`
	//وزن کل سفارشات
	Weight uint32 `form:"weight" json:"weight" validate:"required,numeric" example:"1000"`
	//ارزش تقریبی کل سفارشات جهت بیمه ارسال
	Value uint32 `form:"value" json:"value" validate:"required,numeric" example:"10000"`
	//شناسه آدرس مبدا ارسال سفارش
	AddressID uint64 `form:"address_id" json:"address_id" validate:"required,numeric" example:"1"`
}
