package DTOs

type CreateComment struct {
	PostID uint64 `form:"post_id" json:"post_id" validate:"required,numeric"`
	Title  string `form:"title" json:"title" validate:"required,min=3,max=255"`
	Body   string `form:"body" json:"body" validate:"required,min=3"`
	Email  string `form:"email" json:"email" validate:"required,email"`
}

type IndexComment struct {
	Index
}
