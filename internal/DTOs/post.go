package DTOs

type CreatePost struct {
	//عنوان مقاله
	Title string `form:"title" json:"title" validate:"required" example:"آموزش برنامه نویسی"`
	//محتوای مقاله که میتواند از اچ تی ام ال تشکیل شده باشد
	Body string `form:"body" json:"body" validate:"required" example:"<p>متن مقاله</p>"`
	//شناسه تصویر شاخص مقاله
	GalleryID uint64 `form:"gallery_id" json:"gallery_id" validate:"required" example:"1"`
	//آدرس صفحه که باید به صورت انگلیسی و منحصر به فرد باشد با دش از هم جدا شده باشد
	Slug string `form:"slug" json:"slug" validate:"required" example:"amoozesh-barnamenevisi"`
	//شناسه دسته بندی مقاله
	CategoryID uint64 `form:"category_id" json:"category_id" validate:"required" example:"1"`
}

type UpdatePost struct {
	//شناسه مقاله برای ویرایش
	ID uint64 `form:"id" json:"id" validate:"required" example:"1"`
	//عنوان مقاله
	Title string `form:"title" json:"title" validate:"omitempty" example:"آموزش برنامه نویسی"`
	//محتوای مقاله که میتواند از اچ تی ام ال تشکیل شده باشد
	Body string `form:"body" json:"body" validate:"omitempty" example:"<p>متن مقاله</p>"`
	//شناسه تصویر شاخص مقاله
	GalleryID uint64 `form:"gallery_id" json:"gallery_id" validate:"omitempty" example:"1"`
	//آدرس صفحه که باید به صورت انگلیسی و منحصر به فرد باشد با دش از هم جدا شده باشد
	Slug string `form:"slug" json:"slug" validate:"omitempty" example:"amoozesh-barnamenevisi"`
	//شناسه دسته بندی مقاله
	CategoryID uint64 `form:"category_id" json:"category_id" validate:"omitempty" example:"1"`
}

type IndexPost struct {
	Index
	//شناسه فروشگاه
	ShopID uint64 `form:"shop_id" json:"shop_id" validate:"required" example:"1"`
}
