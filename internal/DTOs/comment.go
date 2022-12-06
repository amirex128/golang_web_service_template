package DTOs

type CreateComment struct {
	//شناسه پست
	PostID uint64 `form:"post_id" json:"post_id" validate:"required,numeric" example:"1" fake:"{number:1,100}"`
	//نام کاربر
	Name string `form:"name" json:"name" validate:"required" example:"نام" fake:"{name}"`
	// شناسه صاحب فروشگاه
	UserID uint64 `form:"user_id" json:"user_id" validate:"required,numeric" example:"1" fake:"{number:1,100}"`
	//متن نظر
	Body string `form:"body" json:"body" validate:"required" example:"متن نظر" fake:"{sentence:10}"`
	//ایمیل
	Email string `form:"email" json:"email" validate:"required" example:"amirex128@gmail.com" fake:"{email}"`
}

type IndexComment struct {
	Index
}
