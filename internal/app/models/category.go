package models

import (
	"backend/internal/app/utils"
	"database/sql"
	"encoding/gob"
	"github.com/sirupsen/logrus"
	"io"
)

type Category struct {
	ID          int            `gorm:"primary_key;auto_increment" json:"id"`
	ParentID    int            `json:"parent_id"`
	Name        string         `json:"name"`
	Sort        uint           `json:"sort"`
	Equivalent  sql.NullString `json:"equivalent"`
	GuildIds    sql.NullString `json:"guild_ids"`
	Description sql.NullString `json:"description"`
	Icon        sql.NullString `json:"icon"`
	Depth       uint           `json:"depth"`
	Products    []Product      `gorm:"many2many:product_category;"`
	Options     []Option       `gorm:"many2many:category_options;"`
}

type CategoryRelated struct {
	CategoryID        int `json:"category_id"`
	CategoryRelatedID int `json:"category_related_id"`
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

	return true

}

func (m *MysqlManager) GetAllCategories() ([]*Category, error) {
	categories := make([]*Category, 0)
	err := m.GetConn().Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (m *MysqlManager) CreateAllCategories(files [][]string) {
	categories := make([]Category, 0)
	for i := range files {
		value := files[i]
		categories = append(categories, Category{
			ID:          utils.StringToInt(value[0]),
			ParentID:    utils.StringToInt(value[0]),
			Name:        value[2],
			Sort:        utils.StringToUint(value[3]),
			Equivalent:  utils.StringConvert(value[4]),
			GuildIds:    utils.StringConvert(value[5]),
			Description: utils.StringConvert(value[7]),
			Icon:        utils.StringConvert(value[8]),
			Depth:       utils.StringToUint(value[9]),
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
			CategoryID:        utils.StringToInt(value[0]),
			CategoryRelatedID: utils.StringToInt(value[1]),
		})
	}
	err := m.GetConn().CreateInBatches(categoryRelated, 100).Error
	if err != nil {
		logrus.Error("seed categoryRelated error: ", err)
	}
}
