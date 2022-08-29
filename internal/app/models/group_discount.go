package models

type GroupDiscount struct {
	ID        int64   `json:"id"`
	ProductID int64   `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Value     float32 `json:"value"`
	Type      string  `json:"type"` // percent or fixed
	Percent   string  `json:"percent"`
	StartedAt string  `json:"started_at"`
	EndedAt   string  `json:"ended_at"`
	CreatedAt string  `json:"created_at"`
}
