package models

type OrderItem struct {
	ID        int64 `json:"id"`
	ProductID int64 `json:"product_id"`
	Count     uint  `json:"count"`
}
