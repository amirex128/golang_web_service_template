package DTOs

type CreatePage struct {
	//عنوان صفحه
	Title string `form:"title" json:"title" validate:"required" example:"درباره ما"`
	//محتوا صفحه که میتواند از اچ تی ام ال تشکیل شده باشد
	Body string `form:"body" json:"body" validate:"required" example:"<p>متن صفحه درباره ما</p>"`
	//آدرس صفحه که باید به صورت انگلیسی و منحصر به فرد باشد با دش از هم جدا شده باشد
	Slug string `form:"slug" json:"slug" validate:"required" example:"about-us"`
	//نوع صفحه که همراه با چارچوب قالب باشد یا صفحه کاملا خالی
	Type string `form:"type" json:"type" validate:"required" example:"normal" enums:"normal,blank"`
	//شناسه فروشگاه
	ShopID uint `form:"shop_id" json:"shop_id" validate:"required" example:"1"`
}
type UpdatePage struct {
	ID uint `form:"id" json:"id" validate:"required" example:"1"`
	//عنوان صفحه
	Title string `form:"title" json:"title" validate:"omitempty" example:"تماس با ما"`
	//محتوا صفحه که میتواند از اچ تی ام ال تشکیل شده باشد
	Body string `form:"body" json:"body" validate:"omitempty" example:"<p>متن صفحه تماس با ما</p>"`
	//نوع صفحه که همراه با چارچوب قالب باشد یا صفحه کاملا خالی
	Type string `form:"type" json:"type" validate:"omitempty" example:"normal" enums:"normal,blank"`
	//آدرس صفحه که باید به صورت انگلیسی و منحصر به فرد باشد با دش از هم جدا شده باشد
	Slug string `form:"slug" json:"slug" validate:"omitempty" example:"contact-us"`
}

type IndexPage struct {
	Index
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" example:"1"`
}
