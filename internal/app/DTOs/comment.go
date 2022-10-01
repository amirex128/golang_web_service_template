package DTOs

type CreateComment struct {
	PostID uint64 `form:"post_id" json:"post_id" validate:"required,numeric"`
	Title  string `form:"title" json:"title" validate:"required,min=3,max=255"`
	Body   string `form:"body" json:"body" validate:"required,min=3"`
}

type IndexComment struct {
	Search   string `form:"search" json:"search"`
	Page     uint32 `form:"page" json:"page" validate:"numeric"`
	PageSize uint32 `form:"page_size" json:"page_size" validate:"numeric"`
	Sort     string `form:"sort" json:"sort"`
}
