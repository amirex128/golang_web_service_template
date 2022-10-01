package DTOs

import "mime/multipart"

type CreatePost struct {
	Title         string                `json:"title" validate:"required"`
	Body          string                `json:"body" validate:"required"`
	Thumbnail     *multipart.FileHeader `json:"thumbnail" validate:"required"`
	ThumbnailPath string                `json:"thumbnail_path"`
	Slug          string                `json:"slug" validate:"required"`
	CategoryID    uint64                `json:"category_id" validate:"required"`
	CreatedAt     string                `json:"created_at" validate:"required"`
	UpdatedAt     string                `json:"updated_at" validate:"required"`
}

type UpdatePost struct {
	Title           string                `json:"title" validate:"required"`
	Body            string                `json:"body" validate:"required"`
	Thumbnail       *multipart.FileHeader `json:"thumbnail" validate:"required"`
	ThumbnailPath   string                `json:"thumbnail_path"`
	ThumbnailRemove string                `form:"thumbnail_remove"`
	Slug            string                `json:"slug" validate:"required"`
	CategoryID      uint64                `json:"category_id" validate:"required"`
}

type IndexPost struct {
	Search   string `form:"search" json:"search"`
	Page     uint32 `form:"page" json:"page" validate:"numeric"`
	PageSize uint32 `form:"page_size" json:"page_size" validate:"numeric"`
	Sort     string `form:"sort" json:"sort"`
}
