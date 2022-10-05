package DTOs

import (
	"gorm.io/gorm"
	"math"
)

type Pagination struct {
	PageSize   uint32      `form:"page_size" json:"page_size"`
	Page       uint32      `form:"page" json:"page"`
	Sort       string      `form:"sort,omitempty" json:"sort,omitempty"`
	TotalRows  uint64      `form:"total_rows" json:"total_rows"`
	TotalPages uint32      `form:"total_pages" json:"total_pages"`
	Data       interface{} `form:"data" json:"data"`
}

func (p *Pagination) GetOffset() uint32 {
	return (p.GetPage() - 1) * p.GetPageSize()
}

func (p *Pagination) GetPageSize() uint32 {
	if p.PageSize == 0 {
		p.PageSize = 3
	}
	return p.PageSize
}

func (p *Pagination) GetPage() uint32 {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}
func Paginate(table string, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	size := pagination.GetPageSize()
	sort := pagination.GetSort()
	offset := pagination.GetOffset()
	var totalRows int64

	db.Table(table).Count(&totalRows)
	pagination.TotalRows = uint64(totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.PageSize)))
	pagination.TotalPages = uint32(totalPages)

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(offset)).Limit(int(size)).Order(sort)
	}
}
