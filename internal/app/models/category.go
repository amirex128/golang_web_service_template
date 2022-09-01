package models

import (
	"backend/internal/app/helpers"
	"database/sql"
	"encoding/gob"
	"github.com/sirupsen/logrus"
	"io"
)

type Category struct {
	ID          int            `json:"id"`
	ParentID    int            `json:"parent_id"`
	Name        string         `json:"name"`
	Sort        uint           `json:"sort"`
	Equivalent  sql.NullString `json:"equivalent"`
	GuildIds    sql.NullString `json:"guild_ids"`
	Description sql.NullString `json:"description"`
	Icon        sql.NullString `json:"icon"`
	Depth       uint           `json:"depth"`
}

type CategoryOption struct {
	CategoryID int `json:"category_id"`
	OptionID   int `json:"option_id"`
}

type CategoryProduct struct {
	CategoryID int   `json:"category_id"`
	ProductID  int64 `json:"product_id"`
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
	manager.GetConn().AutoMigrate(&CategoryOption{})
	manager.GetConn().AutoMigrate(&CategoryProduct{})
	manager.GetConn().AutoMigrate(&CategoryRelated{})

	if manager.GetConn().First(&Category{}).Error == nil {
		return false
	}

	categories := helpers.ReadCsvFile("../../csv/categories.csv")
	manager.CreateAllCategories(categories)
	categoryRelated := helpers.ReadCsvFile("../../csv/category_related.csv")
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
			ID:          intConvert(value[0]),
			ParentID:    intConvert(value[0]),
			Name:        value[2],
			Sort:        uintConvert(value[3]),
			Equivalent:  stringConvert(value[4]),
			GuildIds:    stringConvert(value[5]),
			Description: stringConvert(value[7]),
			Icon:        stringConvert(value[8]),
			Depth:       uintConvert(value[9]),
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
			CategoryID:        intConvert(value[0]),
			CategoryRelatedID: intConvert(value[1]),
		})
	}
	err := m.GetConn().CreateInBatches(categoryRelated, 100).Error
	if err != nil {
		logrus.Error("seed categoryRelated error: ", err)
	}
}
