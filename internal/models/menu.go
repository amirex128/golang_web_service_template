package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/brianvoe/gofakeit/v6"
	"go.elastic.co/apm/v2"
)

type Menu struct {
	ID       uint64  `gorm:"primary_key;auto_increment" json:"id"`
	UserID   *uint64 `json:"user_id"`
	Name     string  `json:"name"`
	Link     string  `json:"link"`
	ShopID   *uint64 `gorm:"default:null" json:"shop_id"`
	ParentID uint64  `gorm:"default:null" json:"parent_id"`
	Position string  `gorm:"default:top" json:"position" sql:"type:ENUM('top','bottom','left','right')"`
	Sort     uint32  `json:"sort"`
}

func initMenu(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&Menu{}) {
		manager.GetConn().Migrator().CreateTable(&Menu{})
		for i := 0; i < 100; i++ {
			model := new(DTOs.CreateMenu)
			gofakeit.Struct(model)

			manager.CreateMenu(*model)
		}
	}

}

func (m *MysqlManager) CreateMenu(dto DTOs.CreateMenu) (*Menu, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showMenu", "model")
	defer span.End()

	// find last sort number
	var lastSort Menu
	err := m.GetConn().Order("sort desc").Where("shop_id = ?", dto.ShopID).First(&lastSort).Error
	if err != nil {
		lastSort.Sort = 0
	}

	menu := &Menu{
		Name:     dto.Name,
		Link:     dto.Link,
		UserID:   utils.GetUserID(m.Ctx),
		ShopID:   &dto.ShopID,
		ParentID: dto.ParentID,
		Position: dto.Position,
		Sort:     lastSort.Sort + 1,
	}
	err = m.GetConn().Create(&menu).Error
	if err != nil {
		return menu, errorx.New("خطا در ایجاد کد منو", "model", err)
	}
	return menu, nil
}
func (m *MysqlManager) UpdateMenu(dto DTOs.UpdateMenu) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showMenu", "model")
	defer span.End()
	menu := &Menu{}
	err := m.GetConn().Where("id = ?", dto.ID).First(menu).Error
	if err != nil {
		return errorx.New("منو یافت نشد", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, menu.UserID); err != nil {
		return err
	}
	if dto.Name != "" {
		menu.Name = dto.Name
	}
	if dto.Link != "" {
		menu.Link = dto.Link
	}
	if dto.ParentID != 0 {
		menu.ParentID = dto.ParentID
	}
	if dto.Position != "" {
		menu.Position = dto.Position
	}
	if dto.Sort != 0 {
		menu.Sort = dto.Sort
	}
	err = m.GetConn().Save(menu).Error
	if err != nil {
		return errorx.New("خطا در ویرایش منو", "model", err)
	}
	return nil
}
func (m *MysqlManager) DeleteMenu(menuID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showMenu", "model")
	defer span.End()
	menu := Menu{}
	err := m.GetConn().Where("id = ?", menuID).First(&menu).Error
	if err != nil {
		return errorx.New("منو یافت نشد", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, menu.UserID); err != nil {
		return err
	}
	err = m.GetConn().Delete(&menu).Error
	if err != nil {
		return errorx.New("خطا در حذف منو", "model", err)
	}
	return nil
}
func (m *MysqlManager) GetAllMenuWithPagination(dto DTOs.IndexMenu) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showMenu", "model")
	defer span.End()
	conn := m.GetConn()
	var menus []Menu
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("menus", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%")
	}
	err := conn.Where("user_id = ?", utils.GetUserID(m.Ctx)).Where("shop_id = ? ", dto.ShopID).Order("id DESC").Find(&menus).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت منو ها", "model", err)
	}
	pagination.Data = menus
	return pagination, nil
}

func (m *MysqlManager) FindMenuByID(id uint64) (*Menu, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindMenuByID", "model")
	defer span.End()
	menu := &Menu{}
	err := m.GetConn().Where("id = ?", id).First(menu).Error
	if err != nil {
		return menu, errorx.New("منو مورد نظر یافت نشد", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, menu.UserID); err != nil {
		return nil, err
	}
	return menu, nil
}
