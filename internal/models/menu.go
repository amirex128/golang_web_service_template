package models

import (
	"context"
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

type Menu struct {
	ID       uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Name     string  `json:"name"`
	Link     string  `json:"link"`
	ShopID   *uint64 `gorm:"default:null" json:"shop_id"`
	ParentID uint64  `gorm:"default:null" json:"parent_id"`
	Position string  `gorm:"default:top" json:"position" sql:"type:ENUM('top','bottom','left','right')"`
	Sort     uint32  `json:"sort"`
}

func initMenu(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Menu{})
	for i := 0; i < 20; i++ {
		manager.CreateMenu(&gin.Context{}, context.Background(), DTOs.CreateMenu{
			Name:     "dsaD",
			Link:     "DSADSA",
			ShopID:   1,
			ParentID: 0,
			Position: "top",
		})
	}
}

func (m *MysqlManager) CreateMenu(c *gin.Context, ctx context.Context, dto DTOs.CreateMenu) error {
	span, ctx := apm.StartSpan(ctx, "showMenu", "model")
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
		ShopID:   &dto.ShopID,
		ParentID: dto.ParentID,
		Position: dto.Position,
		Sort:     lastSort.Sort + 1,
	}
	err = m.GetConn().Create(&menu).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در ایجاد کد منو",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}
func (m *MysqlManager) UpdateMenu(c *gin.Context, ctx context.Context, dto DTOs.UpdateMenu) error {
	span, ctx := apm.StartSpan(ctx, "showMenu", "model")
	defer span.End()
	menu := &Menu{}
	err := m.GetConn().Where("id = ?", dto.ID).First(menu).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "منو یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در ویرایش منو"})
		return err
	}
	return nil
}
func (m *MysqlManager) DeleteMenu(c *gin.Context, ctx context.Context, menuID uint64) error {
	span, ctx := apm.StartSpan(ctx, "showMenu", "model")
	defer span.End()
	menu := Menu{}
	err := m.GetConn().Where("id = ?", menuID).First(&menu).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "منو یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}

	err = m.GetConn().Delete(&menu).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در حذف منو",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}
func (m *MysqlManager) GetAllMenuWithPagination(c *gin.Context, ctx context.Context, dto DTOs.IndexMenu) (*DTOs.Pagination, error) {
	span, ctx := apm.StartSpan(ctx, "showMenu", "model")
	defer span.End()
	conn := m.GetConn()
	var menus []Menu
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("menus", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%").Where("shop_id = ? ", dto.ShopID).Order("id DESC")
	}
	err := conn.Find(&menus).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در دریافت منو ها",
			"error":   err.Error(),
			"type":    "model",
		})
		return pagination, err
	}
	pagination.Data = menus
	return pagination, nil
}
