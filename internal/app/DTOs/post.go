package DTOs

import "mime/multipart"

type CreatePost struct {
	Title         string                `form:"title" json:"title" validate:"required"`
	Body          string                `form:"body" json:"body" validate:"required"`
	Thumbnail     *multipart.FileHeader `form:"thumbnail" json:"thumbnail" validate:"required"`
	ThumbnailPath string                `form:"-" json:"-"`
	Slug          string                `form:"slug" json:"slug" validate:"required"`
	CategoryID    uint64                `form:"category_id" json:"category_id" validate:"required"`
	CreatedAt     string                `form:"created_at" json:"created_at" validate:"required"`
	UpdatedAt     string                `form:"updated_at" json:"updated_at" validate:"required"`
}

type UpdatePost struct {
	Title           string                `form:"title" json:"title" validate:"omitempty"`
	Body            string                `form:"body" json:"body" validate:"omitempty"`
	Thumbnail       *multipart.FileHeader `form:"thumbnail" json:"thumbnail" validate:"omitempty"`
	ThumbnailPath   string                `form:"-" json:"-"`
	ThumbnailRemove string                `form:"-" json:"-"`
	Slug            string                `form:"slug" json:"slug" validate:"omitempty"`
	CategoryID      uint64                `form:"category_id" json:"category_id" validate:"omitempty"`
}

type IndexPost struct {
	Search   string `form:"search" json:"search"`
	Page     uint32 `form:"page" json:"page" validate:"numeric"`
	PageSize uint32 `form:"page_size" json:"page_size" validate:"numeric"`
	Sort     string `form:"sort" json:"sort"`
}
