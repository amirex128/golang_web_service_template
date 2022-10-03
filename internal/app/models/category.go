package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type Category struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	ParentID    uint32    `json:"parent_id"`
	Type        string    `json:"type" sql:"type:ENUM('post','product')"`
	Name        string    `json:"name"`
	Sort        uint32    `json:"sort"`
	Equivalent  string    `json:"equivalent"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Products    []Product `gorm:"many2many:category_product;"`
	Posts       []Post    `gorm:"many2many:category_post;"`
	Options     []Option  `gorm:"many2many:category_option;"`
}

type CategoryRelated struct {
	CategoryID        uint32 `json:"category_id"`
	CategoryRelatedID uint32 `json:"category_related_id"`
}

type CategoryArr []Category

func (s CategoryArr) Len() int {
	return len(s)
}
func (s CategoryArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s CategoryArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Category) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Category) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}

func initCategory(manager *MysqlManager) bool {
	manager.GetConn().AutoMigrate(&Category{})
	manager.GetConn().AutoMigrate(&CategoryRelated{})

	if manager.GetConn().First(&Category{}).Error == nil {
		return false
	}

	categories := utils.ReadCsvFile("../../csv/categories.csv")
	manager.CreateAllCategories(categories)
	categoryRelated := utils.ReadCsvFile("../../csv/category_related.csv")
	manager.CreateAllCategoryRelated(categoryRelated)

	for i := 0; i < 10; i++ {
		manager.CreateCategory(&gin.Context{}, DTOs.CreateCategory{
			Name:        "دسته بندی " + utils.IntToString(i),
			Type:        "post",
			Equivalent:  "کلمه مترادف" + utils.IntToString(i),
			Description: "توضیحات دسته بندی " + utils.IntToString(i),
			Icon:        "icon " + utils.IntToString(i),
		})
	}
	return true

}
func (m *MysqlManager) CreateCategory(c *gin.Context, dto DTOs.CreateCategory) error {
	lastCategory := &Category{}
	err := m.GetConn().Where("type = ?", dto.Type).Order("id desc").Find(lastCategory).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "در دریافت دسته بندی ها مشکلی به وجود آمده است",
		})
		return nil
	}
	category := Category{
		Name:        dto.Name,
		Type:        dto.Type,
		Sort:        lastCategory.Sort + 1,
		Equivalent:  dto.Equivalent,
		Description: dto.Description,
		Icon:        dto.Icon,
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

func (m *MysqlManager) GetLevel1Categories(c *gin.Context) ([]*Category, error) {
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
			ID:          utils.StringToUint32(value[0]),
			ParentID:    utils.StringToUint32(value[1]),
			Name:        value[2],
			Type:        "product",
			Sort:        utils.StringToUint32(value[3]),
			Equivalent:  value[4],
			Description: value[7],
			Icon:        value[8],
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

func (m *MysqlManager) GetAllCategoryWithPagination(c *gin.Context, dto DTOs.IndexCategory) (pagination *DTOs.Pagination, err error) {
	conn := m.GetConn()
	var categories []Category
	pagination = &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate(CategoryTable, pagination, conn)).Order("id DESC")
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

func (m *MysqlManager) FindCategoryByID(c *gin.Context, id uint32) (*Category, error) {
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

func (m *MysqlManager) GetAllCategoryPostWithPagination(c *gin.Context, dto DTOs.IndexPost, categoryID uint32) (pagination *DTOs.Pagination, err error) {
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
		})
		return nil, err
	}
	pagination.Data = posts
	return pagination, nil
}
