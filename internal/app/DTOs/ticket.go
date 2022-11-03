package DTOs

type CreateTicket struct {
	ParentID    uint64 `json:"parent_id" validate:"omitempty"`
	IsAnswer    bool   `json:"is_answer" validate:"omitempty"`
	GuestName   string `json:"guest_name" validate:"omitempty"`
	GuestMobile string `json:"guest_mobile" validate:"omitempty"`
	Title       string `json:"title" validate:"required"`
	Body        string `json:"body" validate:"required"`
	GalleryID   uint64 `json:"gallery_id" validate:"omitempty"`
}

type IndexTicket struct {
	Index
}
