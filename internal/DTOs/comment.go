package DTOs

type CreateComment struct {
	//شناسه پست
	PostID uint64 `form:"post_id" json:"post_id" validate:"required,numeric" example:"1"`
	//نام کاربر
	Name string `form:"name" json:"name" validate:"required" example:"نام"`
	//متن نظر
	Body string `form:"body" json:"body" validate:"required" example:"متن نظر"`
	//ایمیل
	Email string `form:"email" json:"email" validate:"required" example:"amirex128@gmail.com"`
}

type IndexComment struct {
	Index
}
