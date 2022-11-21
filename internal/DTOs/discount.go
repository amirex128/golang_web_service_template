package DTOs

type ProductListDiscount struct {
	//شناسه محصول
	ProductID uint64 `json:"product_id" example:"1"`
	//تعداد محصول
	Count uint32 `json:"count" example:"10"`
}
type CheckDiscount struct {
	//شناسه مالک فروشگاه
	UserID uint64 `form:"user_id" json:"user_id" validate:"required,numeric" example:"1"`
	//کد تخفیف
	Code string `form:"code" json:"code" validate:"required" example:"asdf"`
	// لیست محصولات موجود در سبد خرید
	ProductIDs []ProductListDiscount `form:"product_ids[]" json:"product_ids" validate:"required"`
}

type CreateDiscount struct {
	//کد تخفیف
	Code string `form:"code" json:"code" validate:"required" example:"asdf"`
	//تاریخ شروع تخفیف
	StartedAt string `form:"started_at" json:"started_at" validate:"required" example:"2021-01-01 00:00:00"`
	//تاریخ پایان تخفیف
	EndedAt string `form:"ended_at" json:"ended_at" validate:"required" example:"2025-01-01 00:00:00"`
	//تعداد تخفیف
	Count uint32 `form:"count" json:"count" validate:"required,numeric" example:"10"`
	//نوع تخفیف
	Type string `form:"type" json:"type" validate:"required,oneof=percent amount" example:"percent" Enums:"percent,amount"`
	//مقدار تخفیف
	Amount float32 `form:"amount" json:"amount" validate:"omitempty,numeric" example:"0.0"`
	//درصد تخفیف
	Percent float32 `form:"percent" json:"percent" validate:"omitempty,numeric" example:"50.0"`
	//شناسه محصولات
	ProductIDs []uint64 `form:"product_ids[]" json:"product_ids" validate:"omitempty,dive,numeric"`
	//وضعیت تخفیف
	Status bool `form:"status" json:"status" validate:"required,numeric" example:"true"`
}

type UpdateDiscount struct {
	//شناسه تخفیف برای ویرایش
	ID uint64 `form:"id" json:"id" validate:"required,numeric" example:"1"`
	//کد تخفیف
	Code string `form:"code" json:"code" validate:"omitempty" example:"asdf"`
	//تاریخ شروع تخفیف
	StartedAt string `form:"started_at" json:"started_at" validate:"omitempty" example:"2021-01-01 00:00:00"`
	//تاریخ پایان تخفیف
	EndedAt string `form:"ended_at" json:"ended_at" validate:"omitempty" example:"2025-01-01 00:00:00"`
	//تعداد تخفیف
	Count uint32 `form:"count" json:"count" validate:"omitempty,numeric" example:"10"`
	//نوع تخفیف
	Type string `form:"type" json:"type" validate:"omitempty,oneof=percent amount" example:"percent" Enums:"percent,amount"`
	//مقدار تخفیف
	Amount float32 `form:"amount" json:"amount" validate:"omitempty,numeric" example:"0.0"`
	//درصد تخفیف
	Percent float32 `form:"percent" json:"percent" validate:"omitempty,numeric" example:"40.0"`
	//شناسه محصولات
	ProductIDs []uint64 `form:"product_ids[]" json:"product_ids" validate:"omitempty,dive,numeric"`
	//وضعیت تخفیف
	Status bool `form:"status" json:"status" validate:"omitempty,numeric" example:"true"`
}
type IndexDiscount struct {
	Index
}
