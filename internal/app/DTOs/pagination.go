package DTOs

import (
	"gorm.io/gorm"
	"math"
)

type Pagination struct {
	PageSize   int         `json:"page_size"`
	Page       int         `json:"page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPageSize()
}

func (p *Pagination) GetPageSize() int {
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	return p.PageSize
}

func (p *Pagination) GetPage() int {
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
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.PageSize)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(size).Order(sort)
	}
}
