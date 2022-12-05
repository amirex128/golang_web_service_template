package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/v2"
)

type Category struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	ParentID    uint64    `json:"parent_id"`
	UserID      *uint64   `gorm:"default:null" json:"user_id"`
	Type        string    `json:"type" sql:"type:ENUM('post','product')"`
	Name        string    `json:"name"`
	Sort        uint32    `json:"sort"`
	Equivalent  string    `json:"equivalent"`
	Description string    `json:"description"`
	GalleryID   *uint64   `gorm:"default:null" json:"gallery_id"`
	Gallery     *Gallery  `json:"gallery"`
	Products    []Product `gorm:"many2many:category_product;"`
	Posts       []Post    `gorm:"many2many:category_post;"`
	Options     []Option  `gorm:"many2many:category_option;"`
}

type CategoryRelated struct {
	CategoryID        uint32 `json:"category_id"`
	CategoryRelatedID uint32 `json:"category_related_id"`
}

func initCategory(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&Category{}) {
		manager.GetConn().AutoMigrate(&Category{})
		categories := utils.ReadCsvFile("./csv/categories.csv")
		manager.CreateAllCategories(categories)

		for i := 0; i < 100; i++ {
			model := new(DTOs.CreateCategory)
			gofakeit.Struct(model)

			manager.CreateCategory(*model)
		}
	}
	if !manager.GetConn().Migrator().HasTable(&CategoryRelated{}) {
		manager.GetConn().AutoMigrate(&CategoryRelated{})
		categoryRelated := utils.ReadCsvFile("./csv/category_related.csv")
		manager.CreateAllCategoryRelated(categoryRelated)
	}

}

func (m *MysqlManager) CreateCategory(dto DTOs.CreateCategory) (*Category, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreateCategory", "model")
	defer span.End()
	var lastSort Category
	err := m.GetConn().Where("type = ?", dto.Type).Order("sort desc").First(&lastSort).Error
	if err != nil {
		lastSort.Sort = 0
	}
	userID := GetUser(m.Ctx)
	category := &Category{
		Name:     dto.Name,
		ParentID: 0,
		UserID:   userID,
		Type:     dto.Type,
		GalleryID: func() *uint64 {
			if dto.GalleryID == 0 {
				return nil
			}
			return &dto.GalleryID
		}(),
		Sort:        lastSort.Sort + 1,
		Equivalent:  dto.Equivalent,
		Description: dto.Description,
	}
	err = m.GetConn().Create(category).Error
	if err != nil {
		return category, errorx.New("در ایجاد دسته بندی مشکلی به وجود آمده است", "model", err)
	}
	return category, nil
}

func (m *MysqlManager) GetLevel1Categories() ([]*Category, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetLevel1Categories", "model")
	defer span.End()
	categories := make([]*Category, 0)
	err := m.GetConn().Where("type = ? and parent_id = ?", "post", 0).Find(&categories).Error
	if err != nil {
		return nil, errorx.New("در دریافت دسته بندی ها مشکلی به وجود آمده است", "model", err)
	}
	return categories, nil
}

func (m *MysqlManager) CreateAllCategories(files [][]string) {
	categories := make([]Category, 0)
	for i := range files {
		value := files[i]
		categories = append(categories, Category{
			ID:          utils.StringToUint64(value[0]),
			ParentID:    utils.StringToUint64(value[1]),
			GalleryID:   nil,
			UserID:      nil,
			Name:        value[2],
			Type:        "product",
			Sort:        utils.StringToUint32(value[3]),
			Equivalent:  value[4],
			Description: value[7],
		})
	}
	err := m.GetConn().CreateInBatches(categories, 100).Error
	if err != nil {
		logrus.Error("seed categories error: ", err)
	}
}

func (m *MysqlManager) CreateAllCategoryRelated(files [][]string) {
	categoryRelated := make([]CategoryRelated, 0)
	for i := range files {
		value := files[i]
		categoryRelated = append(categoryRelated, CategoryRelated{
			CategoryID:        utils.StringToUint32(value[0]),
			CategoryRelatedID: utils.StringToUint32(value[1]),
		})
	}
	err := m.GetConn().CreateInBatches(categoryRelated, 100).Error
	if err != nil {
		logrus.Error("seed categoryRelated error: ", err)
	}
}

func (m *MysqlManager) GetAllCategoryWithPagination(dto DTOs.IndexCategory) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetAllCategoryWithPagination", "model")
	defer span.End()
	conn := m.GetConn()
	var categories []Category
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("categories", pagination, conn)).Order("id DESC")
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%")
	}
	err := conn.Find(&categories).Error
	if err != nil {
		return nil, errorx.New("در دریافت دسته بندی ها مشکلی به وجود آمده است", "model", err)
	}
	pagination.Data = categories
	return pagination, nil
}

func (m *MysqlManager) FindCategoryByID(id uint64) (*Category, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindCategoryByID", "model")
	defer span.End()
	category := &Category{}
	err := m.GetConn().Where("id = ?", id).First(category).Error
	if err != nil {
		return nil, errorx.New("در یافت دسته بندی مشکلی به وجود آمده است", "model", err)
	}
	return category, nil
}

func (m *MysqlManager) GetAllCategoryPostWithPagination(dto DTOs.IndexPost, categoryID uint32) (pagination *DTOs.Pagination, err error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetAllCategoryPostWithPagination", "model")
	defer span.End()
	conn := m.GetConn()
	var posts []Post
	pagination = &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("posts", pagination, conn)).Where("id IN (?)", conn.Table("category_post").Where("category_id = ?", categoryID).Select("post_id")).Preload("User").Preload("Categories").Order("id DESC")
	if dto.Search != "" {
		conn = conn.Where("title LIKE ?", "%"+dto.Search+"%")
	}
	err = conn.Find(&posts).Error
	if err != nil {
		return nil, errorx.New("در دریافت پست ها مشکلی به وجود آمده است", "model", err)
	}
	pagination.Data = posts
	return pagination, nil
}

func (m *MysqlManager) UpdateCategory(dto DTOs.UpdateCategory) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:UpdateCategory", "model")
	defer span.End()
	category, err := m.FindCategoryByID(dto.ID)
	if err != nil {
		return nil
	}

	userID := GetUser(m.Ctx)
	if *category.UserID != *userID {
		return errorx.New("شما اجازه ویرایش این دسته بندی را ندارید", "model", err)
	}

	if category.ParentID != dto.ParentID {
		category.ParentID = dto.ParentID
	}
	if category.Name != dto.Name {
		category.Name = dto.Name
	}
	if category.Equivalent != dto.Equivalent {
		category.Equivalent = dto.Equivalent
	}
	if category.Description != dto.Description {
		category.Description = dto.Description
	}
	if category.Sort != dto.Sort {
		category.Sort = dto.Sort
	}
	if *category.GalleryID != dto.GalleryID {
		category.GalleryID = &dto.GalleryID
	}
	err = m.GetConn().Save(category).Error
	if err != nil {
		return errorx.New("در ویرایش دسته بندی مشکلی به وجود آمده است", "model", err)
	}
	return nil
}

func (m *MysqlManager) DeleteCategory(id uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:DeleteCategory", "model")
	defer span.End()
	category, err := m.FindCategoryByID(id)
	if err != nil {
		return nil
	}

	userID := GetUser(m.Ctx)
	if *category.UserID != *userID {
		return errorx.New("شما اجازه حذف این دسته بندی را ندارید", "model", err)
	}

	err = m.GetConn().Delete(category).Error
	if err != nil {
		return errorx.New("در حذف دسته بندی مشکلی به وجود آمده است", "model", err)
	}
	return nil
}
