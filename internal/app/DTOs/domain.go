package DTOs

type CreateDomain struct {
	ShopID uint64 `form:"shop_id" json:"shop_id" validate:"required,numeric"`
	Name   string `form:"name" json:"name" validate:"required"`
	Type   string `form:"type" json:"type" validate:"required"`
}
