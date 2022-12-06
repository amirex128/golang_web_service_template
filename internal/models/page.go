package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"go.elastic.co/apm/v2"
)

type Page struct {
	ID        uint64  `json:"id" gorm:"primary_key"`
	Title     string  `json:"title"`
	Body      string  `json:"body"`
	Slug      string  `json:"slug"`
	Type      string  `json:"type" sql:"type:ENUM('blank','normal')"`
	UserID    *uint64 `json:"user_id"`
	ShopID    *uint64 `gorm:"default:null" json:"shop_id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func initPage(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&Page{}) {
		manager.GetConn().Migrator().CreateTable(&Page{})

	}

}

func (m *MysqlManager) CreatePage(dto DTOs.CreatePage) (*Page, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreatePage", "model")
	defer span.End()
	page := &Page{
		Title:     dto.Title,
		Body:      dto.Body,
		Slug:      dto.Slug,
		UserID:    GetUserID(m.Ctx),
		ShopID:    &dto.ShopID,
		Type:      dto.Type,
		CreatedAt: utils.NowTime(),
		UpdatedAt: utils.NowTime(),
	}
	err := m.GetConn().Create(page).Error
	if err != nil {
		return page, errorx.New("خطایی در ایجاد صفحه رخ داده است", "model", err)
	}
	return page, nil
}

func (m *MysqlManager) UpdatePage(dto DTOs.UpdatePage) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:UpdatePage", "model")
	defer span.End()

	page := &Page{}
	err := m.GetConn().Where("id = ?", dto.ID).First(page).Error
	if err != nil {
		return errorx.New("صفحه مورد نظر یافت نشد", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, page.UserID); err != nil {
		return err
	}
	if page.Title != dto.Title {
		page.Title = dto.Title
	}
	if page.Body != dto.Body {
		page.Body = dto.Body
	}
	if page.Slug != dto.Slug {
		page.Slug = dto.Slug
	}
	if page.Type != dto.Type {
		page.Type = dto.Type
	}
	page.UpdatedAt = utils.NowTime()
	err = m.GetConn().Save(page).Error
	return err
}

func (m *MysqlManager) FindPageBySlug(slug string) (*Page, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindPageBySlug", "model")
	defer span.End()
	page := &Page{}
	err := m.GetConn().Where("slug = ?", slug).First(page).Error
	if err != nil {
		return nil, errorx.New("صفحه مورد نظر یافت نشد", "model", err)
	}
	return page, nil
}

func (m *MysqlManager) FindPageByID(id uint64) (*Page, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindPageByID", "model")
	defer span.End()
	page := &Page{}
	err := m.GetConn().Where("id = ?", id).First(page).Error
	if err != nil {
		return nil, errorx.New("صفحه مورد نظر یافت نشد", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, page.UserID); err != nil {
		return nil, err
	}
	return page, nil
}
func (m *MysqlManager) FindPageByShopID(id uint64) ([]*Page, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindPageByID", "model")
	defer span.End()
	pages := make([]*Page, 0)
	err := m.GetConn().Where("shop_id = ?", id).Find(pages).Error
	if err != nil {
		return nil, errorx.New("صفحه مورد نظر یافت نشد", "model", err)
	}
	return pages, nil
}
func (m *MysqlManager) DeletePage(pageID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showPage", "model")
	defer span.End()
	page := Page{}
	err := m.GetConn().Where("id = ?", pageID).First(&page).Error
	if err != nil {
		return errorx.New("صفحه یافت نشد", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, page.UserID); err != nil {
		return err
	}
	err = m.GetConn().Delete(&page).Error
	if err != nil {
		return errorx.New("خطا در حذف صفحه", "model", err)
	}
	return nil
}

func (m *MysqlManager) GetAllPageWithPagination(dto DTOs.IndexPage) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showPage", "model")
	defer span.End()
	conn := m.GetConn()
	var pages []Page
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("pages", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("title LIKE ?", "%"+dto.Search+"%")
	}
	err := conn.Where("user_id = ?", GetUserID(m.Ctx)).Where("shop_id = ? ", dto.ShopID).Order("id DESC").Find(&pages).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت صفحه ها", "model", err)
	}
	pagination.Data = pages
	return pagination, nil
}
