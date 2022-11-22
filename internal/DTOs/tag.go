package DTOs

type CreateTag struct {
	//عنوان تگ
	Name string `form:"name" json:"name" binding:"required,min=3,max=40" example:"عنوان تگ"`
	//آدرس تگ
	Slug string `form:"slug" json:"slug" binding:"required,min=3,max=40" example:"slug-tag"`
}

type IndexTag struct {
	Index
}

type AddTag struct {
	//نوع تگ
	Type string `form:"type" json:"type" binding:"required,oneof=post category" example:"post" enums:"post,category"`
	//شناسه مقاله
	PostID uint64 `form:"post_id" json:"post_id" binding:"required,numeric" example:"1"`
	// شناسه محصول
	ProductID uint64 `form:"product_id" json:"product_id" binding:"required,numeric" example:"1"`
	//شناسه تگ
	TagID uint64 `form:"tag_id" json:"tag_id" binding:"required,numeric" example:"1"`
}
