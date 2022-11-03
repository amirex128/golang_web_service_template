package DTOs

type CreateTicket struct {
	ParentID    uint64 `form:"parent_id" json:"parent_id" validate:"omitempty,numeric"`
	IsAnswer    bool   `form:"is_answer" json:"is_answer" validate:"omitempty,boolean"`
	GuestName   string `form:"guest_name" json:"guest_name" validate:"omitempty"`
	GuestMobile string `form:"guest_mobile" json:"guest_mobile" validate:"omitempty"`
	Title       string `form:"title" json:"title" validate:"required"`
	Body        string `form:"body" json:"body" validate:"required"`
	GalleryID   uint64 `form:"gallery_id" json:"gallery_id" validate:"omitempty,numeric"`
}

type IndexTicket struct {
	Index
}
