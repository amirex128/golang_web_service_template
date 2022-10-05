package DTOs

import "mime/multipart"

type IndexProduct struct {
	Search   string `form:"search" json:"search"`
	Page     uint32 `form:"page" json:"page" validate:"numeric"`
	PageSize uint32 `form:"page_size" json:"page_size" validate:"numeric"`
	Sort     string `form:"sort" json:"sort"`
}

type CreateProduct struct {
	ShopID           uint64                  `form:"shop_id" validate:"required"`
	ManufacturerId   uint32                  `form:"manufacturer_id" validate:"numeric"`
	Description      string                  `form:"description" validate:"required,min=3,max=1000"`
	Name             string                  `form:"name" validate:"required,min=3,max=100"`
	ShortDescription string                  `form:"short_description" validate:"required,min=3,max=300"`
	Quantity         uint32                  `form:"quantity" validate:"numeric"`
	Price            float32                 `form:"price" validate:"numeric"`
	Weight           uint32                  `form:"weight" validate:"numeric"`
	Height           uint32                  `form:"height" validate:"numeric"`
	Width            uint32                  `form:"width" validate:"numeric"`
	StartedAt        string                  `form:"started_at" validate:"datetime"`
	EndedAt          string                  `form:"ended_at" validate:"datetime"`
	DeliveryTime     uint32                  `form:"delivery_time" validate:"numeric"` // مدت زمان ارسال
	OptionId         uint32                  `form:"option_id" validate:"numeric"`
	OptionItemID     uint32                  `form:"option_item_id" validate:"numeric"`
	CategoryID       uint64                  `form:"category_id" validate:"required,numeric"`
	Images           []*multipart.FileHeader `form:"images[]"`
	ImagePath        []string                `form:"-"`
}

type UpdateProduct struct {
	ID               uint64                  `form:"id" validate:"omitempty,numeric"`
	ShopID           uint64                  `form:"shop_id" validate:"omitempty,required"`
	ManufacturerId   uint32                  `form:"manufacturer_id" validate:"omitempty,numeric"`
	Description      string                  `form:"description" validate:"omitempty,min=3,max=1000"`
	Name             string                  `form:"name" validate:"omitempty,min=3,max=100"`
	ShortDescription string                  `form:"short_description" validate:"omitempty,min=3,max=300"`
	Quantity         uint32                  `form:"quantity" validate:"omitempty,numeric"`
	Price            float32                 `form:"price" validate:"omitempty,numeric"`
	Weight           uint32                  `form:"weight" validate:"omitempty,numeric"`
	Height           uint32                  `form:"height" validate:"omitempty,numeric"`
	Width            uint32                  `form:"width" validate:"omitempty,numeric"`
	Active           string                  `form:"active" validate:"omitempty,numeric"`
	StartedAt        string                  `form:"started_at" validate:"omitempty,datetime"`
	EndedAt          string                  `form:"ended_at" validate:"omitempty,datetime"`
	DeliveryTime     uint32                  `form:"delivery_time" validate:"omitempty,numeric"` // مدت زمان ارسال
	OptionId         uint32                  `form:"option_id" validate:"omitempty,numeric"`
	OptionItemID     uint32                  `form:"option_item_id" validate:"omitempty,numeric"`
	Images           []*multipart.FileHeader `form:"images[]"`
	ImagePath        []string                `json:"-"`
	ImageRemove      []string                `form:"-"`
}
