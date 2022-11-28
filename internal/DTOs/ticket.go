package DTOs

type CreateTicket struct {
	//شناسه تیکت قبلی
	ParentID uint64 `form:"parent_id" json:"parent_id" validate:"omitempty,numeric" example:"1" fake:"{custom_uint64:0}"`
	//نام کاربری که ثبت نام نکرده و از طریق سایت پیام میگذارد
	GuestName string `form:"guest_name" json:"guest_name" validate:"omitempty" example:"امیر" fake:"{name}"`
	//موبایل کاربری که ثبت نام نکرده و از طریق سایت پیام میگذارد
	GuestMobile string `form:"guest_mobile" json:"guest_mobile" validate:"omitempty" example:"09024809750" fake:"{phone}"`
	//عنوان تیکت
	Title string `form:"title" json:"title" validate:"required" example:"عنوان" fake:"{sentence:1}"`
	//متن تیکت بدون اچ تی ام ال باید باشد
	Body string `form:"body" json:"body" validate:"required" example:"متن پیام" fake:"{sentence:4}"`
	// فایل تصویری پیوست شده
	GalleryID uint64 `form:"gallery_id" json:"gallery_id" validate:"omitempty,numeric" example:"1" fake:"{number:1,100}"`
}

type IndexTicket struct {
	Index
}
