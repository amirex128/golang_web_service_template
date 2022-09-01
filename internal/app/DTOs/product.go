package DTOs

import "mime/multipart"

type IndexProduct struct {
	Search   string `form:"search" json:"search" validate:"min=3,max=50"`
	Page     int    `form:"page" json:"page" validate:"numeric"`
	PageSize int    `form:"page_size" json:"page_size" validate:"numeric"`
	Sort     string `form:"sort" json:"sort"`
}
type CreateProduct struct {
	ManufacturerId   uint                    `form:"manufacturer_id" validate:"numeric"`
	Description      string                  `form:"description" validate:"required,min=3,max=1000"`
	Name             string                  `form:"name" validate:"required,min=3,max=100"`
	ShortDescription string                  `form:"short_description" validate:"required,min=3,max=300"`
	Quantity         int                     `form:"quantity" validate:"numeric"`
	Price            float32                 `form:"price" validate:"numeric"`
	FreeSend         byte                    `form:"free_send" validate:"numeric"`
	Weight           uint                    `form:"weight" validate:"numeric"`
	Height           uint                    `form:"height" validate:"numeric"`
	Width            uint                    `form:"width" validate:"numeric"`
	StartedAt        string                  `form:"started_at" validate:"datetime"`
	EndedAt          string                  `form:"ended_at" validate:"datetime"`
	DeliveryTime     uint                    `form:"delivery_time" validate:"numeric"` // مدت زمان ارسال
	OptionId         uint                    `form:"option_id" validate:"numeric"`
	OptionItemID     uint                    `form:"option_item_id" validate:"numeric"`
	Images           []*multipart.FileHeader `form:"images[]"`
	ImagePath        []string
}
