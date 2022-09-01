package DTOs

type CheckDiscount struct {
	Code       string   `json:"code" binding:"required"`
	ProductIDs []uint64 `json:"product_ids" binding:"required"`
}

type CreateDiscount struct {
	Code       string  `json:"code"`
	StartedAt  string  `json:"started_at"`
	EndedAt    string  `json:"ended_at"`
	Count      uint32  `json:"count"`
	Type       string  `json:"type"`
	Amount     float32 `json:"amount"`
	Percent    float32 `json:"percent"`
	ProductIDs string  `json:"product_ids"`
	Status     byte    `json:"status"`
}

type UpdateDiscount struct {
	Code       string  `json:"code"`
	StartedAt  string  `json:"started_at"`
	EndedAt    string  `json:"ended_at"`
	Count      uint32  `json:"count"`
	Type       string  `json:"type"`
	Amount     float32 `json:"amount"`
	Percent    float32 `json:"percent"`
	ProductIDs string  `json:"product_ids"`
	Status     byte    `json:"status"`
}
type IndexDiscount struct {
	Search string `json:"search"`
}
