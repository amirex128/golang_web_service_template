package DTOs

type CreatePage struct {
	Title  string `form:"title" json:"title" validate:"required"`
	Body   string `form:"body" json:"body" validate:"required"`
	Slug   string `form:"slug" json:"slug" validate:"required"`
	Type   string `form:"type" json:"type" validate:"required"`
	ShopID uint   `form:"shop_id" json:"shop_id" validate:"required"`
}
type UpdatePage struct {
	ID    uint   `form:"id" json:"id" validate:"required"`
	Title string `form:"title" json:"title" validate:"omitempty"`
	Body  string `form:"body" json:"body" validate:"omitempty"`
	Type  string `form:"type" json:"type" validate:"omitempty"`
	Slug  string `form:"slug" json:"slug" validate:"omitempty"`
}

type IndexPage struct {
	Index
	ShopID uint64 `form:"shop_id" json:"shop_id"`
}
