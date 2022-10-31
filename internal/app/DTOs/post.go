package DTOs

type CreatePost struct {
	Title      string   `form:"title" json:"title" validate:"required"`
	Body       string   `form:"body" json:"body" validate:"required"`
	GalleryIDs []uint64 `form:"gallery_ids" json:"gallery_ids" validate:"required,dive,numeric"`
	Slug       string   `form:"slug" json:"slug" validate:"required"`
	CategoryID uint64   `form:"category_id" json:"category_id" validate:"required"`
	CreatedAt  string   `form:"created_at" json:"created_at" validate:"required"`
	UpdatedAt  string   `form:"updated_at" json:"updated_at" validate:"required"`
}

type UpdatePost struct {
	Title      string   `form:"title" json:"title" validate:"omitempty"`
	Body       string   `form:"body" json:"body" validate:"omitempty"`
	GalleryIDs []uint64 `form:"gallery_ids" json:"gallery_ids" validate:"omitempty,dive,numeric"`
	Slug       string   `form:"slug" json:"slug" validate:"omitempty"`
	CategoryID uint64   `form:"category_id" json:"category_id" validate:"omitempty"`
}

type IndexPost struct {
	Index
}
