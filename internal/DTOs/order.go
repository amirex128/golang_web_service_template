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
	Description  string      `form:"description" json:"description"`
	VerifyCode   string      `form:"verify_code" json:"verify_code" validate:"required,numeric,min=2,max=20"`
}

type SendOrder struct {
	OrderID     uint64 `form:"order_id" json:"order_id" validate:"required,numeric"`
	Courier     string `form:"courier" json:"courier" validate:"required,oneof=tipax post"`
	PackageSize string `form:"package_size" json:"package_size" validate:"required"`
	Weight      uint32 `form:"weight" json:"weight" validate:"required,numeric"`
	Value       uint32 `form:"value" json:"value" validate:"required,numeric"`
	AddressID   uint64 `form:"address_id" json:"address_id" validate:"required,numeric"`
}

type CalculateOrder struct {
	PackageSize string `form:"package_size" json:"package_size" validate:"required"`
	Weight      uint32 `form:"weight" json:"weight" validate:"required,numeric"`
	Value       uint32 `form:"value" json:"value" validate:"required,numeric"`
	AddressID   uint64 `form:"address_id" json:"address_id" validate:"required,numeric"`
}
