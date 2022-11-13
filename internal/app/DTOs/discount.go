package DTOs

type ProductListDiscount struct {
	ProductID uint64 `json:"product_id"`
	Count     uint32 `json:"count"`
}
type CheckDiscount struct {
	UserID     uint64                `form:"user_id" json:"user_id" validate:"required,numeric"`
	Code       string                `form:"code" json:"code" validate:"required"`
	ProductIDs []ProductListDiscount `form:"product_ids[]" json:"product_ids" validate:"required"`
}

type CreateDiscount struct {
	Code       string   `form:"code" json:"code" validate:"required"`
	StartedAt  string   `form:"started_at" json:"started_at" validate:"required"`
	EndedAt    string   `form:"ended_at" json:"ended_at" validate:"required"`
	Count      uint32   `form:"count" json:"count" validate:"required,numeric"`
	Type       string   `form:"type" json:"type" validate:"required,oneof=percent amount"`
	Amount     float32  `form:"amount" json:"amount" validate:"omitempty,numeric"`
	Percent    float32  `form:"percent" json:"percent" validate:"omitempty,numeric"`
	ProductIDs []uint64 `form:"product_ids[]" json:"product_ids" validate:"omitempty,dive,numeric"`
	Status     byte     `form:"status" json:"status" validate:"required,numeric"`
}

type UpdateDiscount struct {
	ID         uint64   `form:"id" json:"id" validate:"required,numeric"`
	Code       string   `form:"code" json:"code" validate:"omitempty"`
	StartedAt  string   `form:"started_at" json:"started_at" validate:"omitempty"`
	EndedAt    string   `form:"ended_at" json:"ended_at" validate:"omitempty"`
	Count      uint32   `form:"count" json:"count" validate:"omitempty,numeric"`
	Type       string   `form:"type" json:"type" validate:"omitempty,oneof=percent amount"`
	Amount     float32  `form:"amount" json:"amount" validate:"omitempty,numeric"`
	Percent    float32  `form:"percent" json:"percent" validate:"omitempty,numeric"`
	ProductIDs []uint64 `form:"product_ids[]" json:"product_ids" validate:"omitempty,dive,numeric"`
	Status     byte     `form:"status" json:"status" validate:"omitempty,numeric"`
}
type IndexDiscount struct {
	Index
}
