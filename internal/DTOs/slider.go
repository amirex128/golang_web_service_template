package DTOs

type CreateSlider struct {
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validation:"required" example:"1" fake:"{number:1,100}"`
	//لینک اسلایدر
	Link string `form:"link" json:"link" validation:"omitempty" example:"https://google.com" fake:"{url}"`
	//عنوان اسلایدر
	Title string `form:"title" json:"title" validation:"required" example:"عنوان اسلایدر" fake:"{word}"`
	//توضیحات اسلایدر
	Description string `form:"description" json:"description" validation:"omitempty" example:"توضیحات اسلایدر" fake:"{sentence:2}"`
	//تصویر اسلایدر
	GalleryID uint64 `form:"gallery_id" json:"gallery_id" validation:"required" example:"1" fake:"{number:1,100}"`
	//محل قرار گیری اسلایدر
	Position string `form:"position" json:"position" validation:"required" example:"top" enums:"top,bottom,right,left" fake:"{custom_enum:top,bottom,right,left}"`
}

type UpdateSlider struct {
	//شناسه اسلایدر برای ویرایش
	ID uint64 `form:"id" json:"id" validation:"required" example:"1"`
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validation:"omitempty" example:"1"`
	//لینک اسلایدر
	Link string `form:"link" json:"link" validation:"omitempty" example:"https://google.com"`
	//عنوان اسلایدر
	Title string `form:"title" json:"title" validation:"omitempty" example:"عنوان اسلایدر"`
	//توضیحات اسلایدر
	Description string `form:"description" json:"description" validation:"omitempty" example:"توضیحات اسلایدر"`
	//تصویر اسلایدر
	GalleryID uint64 `form:"gallery_id" json:"gallery_id" validation:"omitempty" example:"1"`
	//ترتیب اسلایدر
	Sort uint32 `form:"sort" json:"sort" validation:"omitempty" example:"1"`
	//محل قرار گیری اسلایدر
	Position string `form:"position" json:"position" validation:"omitempty" example:"top" enums:"top,bottom,right,left"`
}

type IndexSlider struct {
	Index
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validation:"omitempty" example:"1"`
}
