package models

type Ticket struct {
	ID          uint64  `gorm:"primary_key;auto_increment" json:"id"`
	UserID      uint64  `json:"user_id"`
	User        User    `json:"user"`
	GuestName   string  `json:"guest_name"`
	GuestMobile string  `json:"guest_mobile"`
	Title       string  `json:"title"`
	Body        string  `json:"body"`
	GalleryID   uint64  `json:"gallery_id"`
	Gallery     Gallery `json:"gallery"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
