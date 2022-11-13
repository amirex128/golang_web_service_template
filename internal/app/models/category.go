package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/v2"
	"net/http"
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

func initCategory(manager *MysqlManager) bool {
	manager.GetConn().AutoMigrate(&Category{})
	manager.GetConn().AutoMigrate(&CategoryRelated{})

	if manager.GetConn().First(&Category{}).Error == nil {
		return false
	}

	categories := utils.ReadCsvFile("./csv/categories.csv")
	manager.CreateAllCategories(categories)
	categoryRelated := utils.ReadCsvFile("./csv/category_related.csv")
	manager.CreateAllCategoryRelated(categoryRelated)

	for i := 0; i < 10; i++ {
		manager.CreateCategory(&gin.Context{}, context.Background(), DTOs.CreateCategory{
			Name:        "دسته بندی " + utils.IntToString(i),
			Type:        "post",
			GalleryID:   0,
			Equivalent:  "کلمه مترادف" + utils.IntToString(i),
			Description: "توضیحات دسته بندی " + utils.IntToString(i),
		})
	}
	return true

}

func (m *MysqlManager) CreateCategory(c *gin.Context, ctx context.Context, dto DTOs.CreateCategory) error {
	span, ctx := apm.StartSpan(ctx, "CreateCategory", "model")
	defer span.End()
	lastCategory := &Category{}
	err := m.GetConn().Where("type = ?", dto.Type).Order("id desc").Find(lastCategory).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "در دریافت دسته بندی ها مشکلی به وجود آمده است",
		})
		return err
	}
	userID := GetUser(c)
	category := Category{
		Name:     dto.Name,
		ParentID: 0,
		UserID:   &userID,
		Type:     dto.Type,
		GalleryID: func() *uint64 {
			if dto.GalleryID == 0 {
				return nil
			}
			return &dto.GalleryID
		}(),
		Sort: func() uint32 {
			if lastCategory != nil {
				return lastCategory.Sort + 1
			}
			return 1
		}(),
		Equivalent:  dto.Equivalent,
		Description: dto.Description,
	}
	err = m.GetConn().Create(&category).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "در ایجاد دسته بندی مشکلی به وجود آمده است",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) GetLevel1Categories(c *gin.Context, ctx context.Context) ([]*Category, error) {
	span, ctx := apm.StartSpan(ctx, "GetLevel1Categories", "model")
	defer span.End()
	categories := make([]*Category, 0)
	err := m.GetConn().Where("type = ? and parent_id = ?", "post", 0).Find(&categories).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "در دریافت دسته بندی ها مشکلی به وجود آمده است",
		})
		return nil, err
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

func (m *MysqlManager) GetAllCategoryWithPagination(c *gin.Context, ctx context.Context, dto DTOs.IndexCategory) (pagination *DTOs.Pagination, err error) {
	span, ctx := apm.StartSpan(ctx, "GetAllCategoryWithPagination", "model")
	defer span.End()
	conn := m.GetConn()
	var categories []Category
	pagination = &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("categories", pagination, conn)).Order("id DESC")
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%")
	}
	err = conn.Find(&categories).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در یافتن پست ها پیش آمده است",
		})
		return nil, err
	}
	pagination.Data = categories
	return pagination, nil
}

func (m *MysqlManager) FindCategoryByID(c *gin.Context, ctx context.Context, id uint64) (*Category, error) {
	span, ctx := apm.StartSpan(ctx, "FindCategoryByID", "model")
	defer span.End()
	category := &Category{}
	err := m.GetConn().Where("id = ?", id).First(category).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در یافتن دسته بندی پیش آمده است",
		})
		return nil, err
	}
	return category, nil
}

func (m *MysqlManager) GetAllCategoryPostWithPagination(c *gin.Context, ctx context.Context, dto DTOs.IndexPost, categoryID uint32) (pagination *DTOs.Pagination, err error) {
	span, ctx := apm.StartSpan(ctx, "GetAllCategoryPostWithPagination", "model")
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در یافتن پست ها پیش آمده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return nil, err
	}
	pagination.Data = posts
	return pagination, nil
}

func (m *MysqlManager) UpdateCategory(c *gin.Context, ctx context.Context, dto DTOs.UpdateCategory) error {
	span, ctx := apm.StartSpan(ctx, "UpdateCategory", "model")
	defer span.End()
	category, err := m.FindCategoryByID(c, ctx, dto.ID)
	if err != nil {
		return nil
	}

	userID := GetUser(c)
	if *category.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "شما اجازه ویرایش این دسته بندی را ندارید",
		})
		return nil
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در ویرایش دسته بندی پیش آمده است",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) DeleteCategory(c *gin.Context, ctx context.Context, id uint64) error {
	span, ctx := apm.StartSpan(ctx, "DeleteCategory", "model")
	defer span.End()
	category, err := m.FindCategoryByID(c, nil, id)
	if err != nil {
		return nil
	}

	userID := GetUser(c)
	if *category.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "شما اجازه حذف این دسته بندی را ندارید",
		})
		return nil
	}

	err = m.GetConn().Delete(category).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در حذف دسته بندی پیش آمده است",
		})
		return err
	}
	return nil
}
