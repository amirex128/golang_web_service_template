package DTOs

type CreateSlider struct {
	ShopID      uint64 `form:"shop_id" json:"shop_id" validation:"required"`
	Link        string `form:"link" json:"link" validation:"omitempty"`
	Title       string `form:"title" json:"title" validation:"required"`
	Description string `form:"description" json:"description" validation:"omitempty"`
	GalleryID   uint64 `form:"gallery_id" json:"gallery_id" validation:"required"`
	Position    string `form:"position" json:"position" validation:"required"`
}

type UpdateSlider struct {
	ID          uint64 `form:"id" json:"id" validation:"required"`
	ShopID      uint64 `form:"shop_id" json:"shop_id" validation:"omitempty"`
	Link        string `form:"link" json:"link" validation:"omitempty"`
	Title       string `form:"title" json:"title" validation:"omitempty"`
	Description string `form:"description" json:"description" validation:"omitempty"`
	GalleryID   uint64 `form:"gallery_id" json:"gallery_id" validation:"omitempty"`
	Sort        uint32 `form:"sort" json:"sort" validation:"omitempty"`
	Position    string `form:"position" json:"position" validation:"omitempty"`
}

type IndexSlider struct {
	Index
	ShopID uint64 `form:"shop_id" json:"shop_id" validation:"omitempty"`
}
