package DTOs

type IndexCategory struct {
	Search   string `form:"search" json:"search"`
	Page     uint32 `form:"page" json:"page" validate:"numeric"`
	PageSize uint32 `form:"page_size" json:"page_size" validate:"numeric"`
	Sort     string `form:"sort" json:"sort"`
}
