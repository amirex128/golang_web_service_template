package DTOs

import "mime/multipart"

type CreateGallery struct {
	File      *multipart.FileHeader `form:"file" json:"file" validate:"required"`
	OwnerID   uint64                `form:"owner_id" json:"owner_id" validate:"required"`
	OwnerType string                `form:"owner_type" json:"owner_type" validate:"required"`
	Width     uint32                `form:"width" json:"width" validate:"required"`
	Height    uint32                `form:"height" json:"height" validate:"required"`
}
