package DTOs

type CreateMenu struct {
	//نام منو برای نمایش در لیست ها
	Name string `form:"name" json:"name" validation:"required" example:"خانه" fake:"{word}"`
	//لینک منو
	Link string `form:"link" json:"link" validation:"omitempty" example:"https://example.selloora.conf/page/test" fake:"{url}"`
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validation:"required" example:"1" fake:"{number:1,100}"`
	//شناسه منو بالا سری برای منو های اصلی و باید صفر ارسال شود
	ParentID uint64 `form:"parent_id" json:"parent_id" validation:"required" example:"0" fake:"{custom_uint64:0}"`
	//محل قرار گیری منو در قالب که با توجه به تنظیمات قالب میباشد
	Position string `form:"position" json:"position" validation:"required" example:"top" enums:"top,bottom,right,left" fake:"{custom_enum:top,bottom,right,left}"`
}

type UpdateMenu struct {
	//شناسه منو برای ویرایش
	ID uint64 `form:"id" json:"id" validation:"required" example:"1"`
	//نام منو
	Name string `form:"name" json:"name" validation:"omitempty" example:"خانه"`
	//لینک منو
	Link string `form:"link" json:"link" validation:"omitempty" example:"https://example.selloora.conf/page/test"`
	//شناسه منو بالا سری
	ParentID uint64 `form:"parent_id" json:"parent_id" validation:"omitempty" example:"0"`
	//شناسه منو بالا سری برای منو های اصلی و باید صفر ارسال شود
	Position string `form:"position" json:"position" validation:"omitempty"  example:"top" enums:"top,bottom,right,left"`
	//ترتیب منو
	Sort uint32 `form:"sort" json:"sort" validation:"omitempty" example:"1"`
}

type IndexMenu struct {
	Index
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validation:"omitempty" example:"1"`
}
