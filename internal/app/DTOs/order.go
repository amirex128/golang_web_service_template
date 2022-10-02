package DTOs

type OrderItem struct {
	ProductID uint64 `form:"product_id" json:"product_id" validate:"required,numeric"`
	OptionID  uint64 `form:"option_id" json:"option_id" validate:"numeric"`
	Count     uint32 `form:"count" json:"count" validate:"required,numeric"`
}

type CreateOrder struct {
	UserID       uint64      `form:"user_id" json:"user_id" validate:"required,numeric"`
	ShopID       uint64      `form:"shop_id" json:"shop_id" validate:"required,numeric"`
	OrderItems   []OrderItem `form:"order_items" json:"order_items" validate:"required,dive"`
	CustomerID   uint64      `form:"customer_id" json:"customer_id" validate:"required,numeric"`
	DiscountCode string      `form:"discount_code" json:"discount_code" validate:"omitempty"`
	SendType     string      `form:"send_type" json:"send_type" validate:"required"`
	Description  string      `form:"description" json:"description"`
	VerifyCode   string      `form:"verify_code" json:"verify_code" validate:"required,numeric,min=2,max=20"`
}
