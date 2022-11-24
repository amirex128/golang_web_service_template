package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"go.elastic.co/apm/v2"
	"os"
)

type Page struct {
	ID        uint64 `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Slug      string `json:"slug"`
	Type      string `json:"type" sql:"type:ENUM('blank','normal')"`
	ShopID    *uint  `gorm:"default:null" json:"shop_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func initPage(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Page{})
	dirs, err := os.ReadDir("./csv/themes")
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			files, err := os.ReadDir("./csv/themes/" + dir.Name())
			if err != nil {
				panic(err)
			}
			for _, file := range files {
				if !file.IsDir() {
					var body string
					readFile, err := os.ReadFile("./csv/themes/" + dir.Name() + "/" + file.Name())
					if err != nil {
						panic(err)
					}
					body = string(readFile)
					_, err = manager.CreatePage(DTOs.CreatePage{
						Title:  dir.Name(),
						Body:   body,
						Type:   "blank",
						Slug:   dir.Name(),
						ShopID: 1,
					})
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}

}

func (m *MysqlManager) CreatePage(dto DTOs.CreatePage) (*Page, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreatePage", "model")
	defer span.End()
	page := &Page{
		Title:     dto.Title,
		Body:      dto.Body,
		Slug:      dto.Slug,
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
	return page, nil
}

func (m *MysqlManager) DeletePage(pageID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showPage", "model")
	defer span.End()
	page := Page{}
	err := m.GetConn().Where("id = ?", pageID).First(&page).Error
	if err != nil {
		return errorx.New("صفحه یافت نشد", "model", err)
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
		conn = conn.Where("title LIKE ?", "%"+dto.Search+"%").Where("shop_id = ? ", dto.ShopID).Order("id DESC")
	}
	err := conn.Find(&pages).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت صفحه ها", "model", err)
	}
	pagination.Data = pages
	return pagination, nil
}
