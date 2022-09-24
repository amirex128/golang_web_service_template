package DTOs

type CheckDiscount struct {
	UserID     uint64   `json:"user_id" validate:"required,numeric"`
	Code       string   `json:"code" binding:"required"`
	ProductIDs []uint64 `json:"product_ids" binding:"required,dive"`
}

type CreateDiscount struct {
	Code       string   `json:"code" validate:"required"`
	StartedAt  string   `json:"started_at" validate:"required"`
	EndedAt    string   `json:"ended_at" validate:"required"`
	Count      uint32   `json:"count" validate:"required,numeric"`
	Type       string   `json:"type" validate:"required"`
	Amount     float32  `json:"amount" validate:"required,numeric"`
	Percent    float32  `json:"percent" validate:"required,numeric"`
	ProductIDs []uint64 `json:"product_ids" validate:"required"`
	Status     byte     `json:"status" validate:"required,numeric"`
}

type UpdateDiscount struct {
	DiscountID uint64   `json:"discount_id" validate:"required,numeric"`
	Code       string   `json:"code" validate:"required"`
	StartedAt  string   `json:"started_at" validate:"required"`
	EndedAt    string   `json:"ended_at" validate:"required"`
	Count      uint32   `json:"count" validate:"required,numeric"`
	Type       string   `json:"type" validate:"required"`
	Amount     float32  `json:"amount" validate:"required,numeric"`
	Percent    float32  `json:"percent" validate:"required,numeric"`
	ProductIDs []uint64 `json:"product_ids" validate:"required"`
	Status     byte     `json:"status" validate:"required,numeric"`
}
type IndexDiscount struct {
	Search string `json:"search"`
}
