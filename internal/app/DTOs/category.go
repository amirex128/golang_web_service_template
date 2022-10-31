package DTOs

type IndexCategory struct {
	Index
}

type CreateCategory struct {
	ParentID    uint32 `form:"parent_id" json:"parent_id" validate:"numeric"`
	Type        string `form:"type" json:"type" validate:"required"`
	Name        string `form:"name" json:"name" validate:"required"`
	Equivalent  string `form:"equivalent" json:"equivalent" validate:"required"`
	Description string `form:"description" json:"description" validate:"required"`
}
