package DTOs

type CheckDiscount struct {
	UserID     uint64   `form:"user_id" json:"user_id" validate:"required,numeric"`
	Code       string   `form:"code" json:"code" binding:"required"`
	ProductIDs []uint64 `form:"product_ids" json:"product_ids" binding:"required,dive"`
}

type CreateDiscount struct {
	Code       string   `form:"code" json:"code" validate:"required"`
	StartedAt  string   `form:"started_at" json:"started_at" validate:"required"`
	EndedAt    string   `form:"ended_at" json:"ended_at" validate:"required"`
	Count      uint32   `form:"count" json:"count" validate:"required,numeric"`
	Type       string   `form:"type" json:"type" validate:"required"`
	Amount     float32  `form:"amount" json:"amount" validate:"required,numeric"`
	Percent    float32  `form:"percent" json:"percent" validate:"required,numeric"`
	ProductIDs []uint64 `form:"product_ids" json:"product_ids" validate:"required"`
	Status     byte     `form:"status" json:"status" validate:"required,numeric"`
}

type UpdateDiscount struct {
	DiscountID uint64   `form:"discount_id" json:"discount_id" validate:"required,numeric"`
	Code       string   `form:"code" json:"code" validate:"required"`
	StartedAt  string   `form:"started_at" json:"started_at" validate:"required"`
	EndedAt    string   `form:"ended_at" json:"ended_at" validate:"required"`
	Count      uint32   `form:"count" json:"count" validate:"required,numeric"`
	Type       string   `form:"type" json:"type" validate:"required"`
	Amount     float32  `form:"amount" json:"amount" validate:"required,numeric"`
	Percent    float32  `form:"percent" json:"percent" validate:"required,numeric"`
	ProductIDs []uint64 `form:"product_ids" json:"product_ids" validate:"required"`
	Status     byte     `form:"status" json:"status" validate:"required,numeric"`
}
type IndexDiscount struct {
	Search string `form:"search" json:"search"`
}