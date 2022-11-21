package DTOs

type IndexCategory struct {
	Index
}

type CreateCategory struct {
	//شناسه دسته بندی پدر
	ParentID uint64 `form:"parent_id" json:"parent_id" validate:"required,numeric" example:"0"`
	//نوع دسته بندی
	Type string `form:"type" json:"type" validate:"required" example:"product" enums:"product,post"`
	//نام دسته بندی
	Name string `form:"name" json:"name" validate:"required" example:"نام دسته بندی"`
	//کلمات مترادف که با کاما از هم جدا شده اند
	Equivalent string `form:"equivalent" json:"equivalent" validate:"omitempty" example:"آموزش,یادگیری"`
	//توضیحات دسته بندی
	Description string `form:"description" json:"description" validate:"omitempty" example:"توضیحات دسته بندی"`
	//شناسه تصویر آپلود شده
	GalleryID uint64 `form:"gallery_id" json:"gallery_id" validate:"omitempty,numeric" example:"1"`
}

type UpdateCategory struct {
	//شناسه دسته بندی برای ویرایش
	ID uint64 `form:"id" json:"id" validate:"required,numeric" example:"1"`
	//شناسه دسته بندی پدر
	ParentID uint64 `form:"parent_id" json:"parent_id" validate:"omitempty,numeric" example:"0"`
	//نام دسته بندی
	Name string `form:"name" json:"name" validate:"omitempty" example:"نام دسته بندی"`
	//کلمات مترادف که با کاما از هم جدا شده اند
	Equivalent string `form:"equivalent" json:"equivalent" validate:"omitempty" example:"آموزش,یادگیری"`
	//توضیحات دسته بندی
	Description string `form:"description" json:"description" validate:"omitempty" example:"توضیحات دسته بندی"`
	//ترتیب مرتب سازی دسته بندی
	Sort uint32 `form:"sort" json:"sort" validate:"omitempty,numeric" example:"1"`
	//شناسه تصویر آپلود شده
	GalleryID uint64 `form:"gallery_id" json:"gallery_id" validate:"omitempty,numeric" example:"1"`
}
