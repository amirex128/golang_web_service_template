package DTOs

type IndexCategory struct {
	Index
}

type CreateCategory struct {
	ParentID    uint64 `form:"parent_id" json:"parent_id" validate:"required,numeric"`
	Type        string `form:"type" json:"type" validate:"required"`
	Name        string `form:"name" json:"name" validate:"required"`
	Equivalent  string `form:"equivalent" json:"equivalent" validate:"omitempty"`
	Description string `form:"description" json:"description" validate:"omitempty"`
	GalleryID   uint64 `form:"gallery_id" json:"gallery_id" validate:"omitempty,numeric"`
}

type UpdateCategory struct {
	ID          uint64 `form:"id" json:"id" validate:"required,numeric"`
	ParentID    uint64 `form:"parent_id" json:"parent_id" validate:"omitempty,numeric"`
	Name        string `form:"name" json:"name" validate:"omitempty"`
	Equivalent  string `form:"equivalent" json:"equivalent" validate:"omitempty"`
	Description string `form:"description" json:"description" validate:"omitempty"`
	Sort        uint32 `form:"sort" json:"sort" validate:"omitempty,numeric"`
	GalleryID   uint64 `form:"gallery_id" json:"gallery_id" validate:"omitempty,numeric"`
}
