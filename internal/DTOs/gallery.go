package DTOs

import "mime/multipart"

type CreateGallery struct {
	File   *multipart.FileHeader `form:"file" json:"file" validate:"required" swaggerignore:"true"`
	Width  uint32                `form:"width" json:"width" validate:"required"`
	Height uint32                `form:"height" json:"height" validate:"required"`
}
