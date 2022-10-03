package DTOs

type IndexCategory struct {
	Search   string `form:"search" json:"search"`
	Page     uint32 `form:"page" json:"page" validate:"numeric"`
	PageSize uint32 `form:"page_size" json:"page_size" validate:"numeric"`
	Sort     string `form:"sort" json:"sort"`
}

type CreateCategory struct {
	ParentID    uint32 `form:"parent_id" json:"parent_id" validate:"numeric"`
	Type        string `form:"type" json:"type" validate:"required"`
	Name        string `form:"name" json:"name" validate:"required"`
	Equivalent  string `form:"equivalent" json:"equivalent" validate:"required"`
	Description string `form:"description" json:"description" validate:"required"`
	Icon        string `form:"icon" json:"icon" validate:"required"`
}
