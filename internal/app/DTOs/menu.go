package DTOs

type CreateMenu struct {
	Name     string `form:"name" json:"name" validation:"required"`
	Link     string `form:"link" json:"link" validation:"omitempty"`
	ShopID   uint64 `form:"shop_id" json:"shop_id" validation:"required"`
	ParentID uint64 `form:"parent_id" json:"parent_id" validation:"required"`
	Position string `form:"position" json:"position" validation:"required"`
}

type UpdateMenu struct {
	ID       uint64 `form:"id" json:"id" validation:"required"`
	Name     string `form:"name" json:"name" validation:"omitempty"`
	Link     string `form:"link" json:"link" validation:"omitempty"`
	ParentID uint64 `form:"parent_id" json:"parent_id" validation:"omitempty"`
	Position string `form:"position" json:"position" validation:"omitempty"`
	Sort     uint32 `form:"sort" json:"sort" validation:"omitempty"`
}

type IndexMenu struct {
	Index
	ShopID uint64 `form:"shop_id" json:"shop_id" validation:"omitempty"`
}
