package DTOs

import "mime/multipart"

type CreateGallery struct {
	File *multipart.FileHeader `form:"file" json:"file" validate:"required" swaggerignore:"true"`
	//عرض تصویر
	Width uint32 `form:"width" json:"width" validate:"required" example:"300"`
	//طول تصویر
	Height uint32 `form:"height" json:"height" validate:"required" example:"300"`
}
