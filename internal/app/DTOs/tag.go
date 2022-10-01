package DTOs

type CreateTag struct {
	Name string `form:"name" json:"name" binding:"required,min=3,max=40"`
	Slug string `form:"slug" json:"slug" binding:"required,min=3,max=40"`
}

type IndexTag struct {
	Search   string `form:"search" json:"search"`
	Page     uint32 `form:"page" json:"page" validate:"numeric"`
	PageSize uint32 `form:"page_size" json:"page_size" validate:"numeric"`
	Sort     string `form:"sort" json:"sort"`
}

type AddTag struct {
	PostID uint64 `form:"post_id" json:"post_id" binding:"required,numeric"`
	TagID  uint64 `form:"tag_id" json:"tag_id" binding:"required,numeric"`
}
