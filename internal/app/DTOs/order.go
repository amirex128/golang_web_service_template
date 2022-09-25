package DTOs

type OrderItem struct {
	ProductID   uint64 `json:"product_id" validate:"required,numeric"`
	OptionID    uint64 `json:"option_id" validate:"numeric"`
	Count       uint64 `json:"count" validate:"required,numeric"`
	Description string `json:"description"`
}

type CreateOrder struct {
	UserID     uint64      `json:"user_id" validate:"required,numeric"`
	OrderItems []OrderItem `json:"order_items" validate:"required,dive"`
	CustomerID uint64      `json:"customer_id" validate:"required,numeric"`
	DiscountID uint64      `json:"discount_id" validate:"numeric"`
	SendType   string      `json:"send_type" validate:"required"`
	VerifyCode string      `json:"verify_code" validate:"required,numeric,min=2,max=20"`
}
