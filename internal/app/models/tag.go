package models

import (
	"backend/internal/app/DTOs"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Tag struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Posts []Post `gorm:"many2many:post_tags;" json:"posts"`
}
type TagArr []Tag

func (s TagArr) Len() int {
	return len(s)
}
func (s TagArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s TagArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Tag) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Tag) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initTag(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Tag{})
}

func (m *MysqlManager) CreateTag(c *gin.Context, dto DTOs.CreateTag) (err error) {
	tag := Tag{
		Name: dto.Name,
		Slug: dto.Slug,
	}
	err = m.GetConn().Create(&tag).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "خطا در ایجاد تگ",
			"error":   err.Error(),
		})
		return err
	}
	return
}

func (m *MysqlManager) GetAllTagsWithPagination(c *gin.Context, dto DTOs.IndexTag) (pagination *DTOs.Pagination, err error) {
	conn := m.GetConn()
	var tags []Tag
	pagination = &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate(TagTable, pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%").Order("id DESC")
	}
	err = conn.Find(&tags).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در یافتن پست ها پیش آمده است",
			"error":   err.Error(),
		})
		return nil, err
	}
	pagination.Data = tags
	return pagination, nil
}

func (m *MysqlManager) DeleteTag(c *gin.Context, id uint64) (err error) {
	err = m.GetConn().Where("id = ?", id).Delete(&Tag{}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "خطا در حذف تگ",
			"error":   err.Error(),
		})
		return err
	}
	return
}

func (m *MysqlManager) AddTag(c *gin.Context, dto DTOs.AddTag) (err error) {
	err = m.GetConn().Model(&Post{ID: dto.PostID}).Association("Tags").Append(&Tag{ID: dto.TagID})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "خطا در افزودن تگ",
			"error":   err.Error(),
		})
		return err
	}
	return
}

func (m *MysqlManager) RandomTags(c *gin.Context, count int) (tags []Tag, err error) {
	err = m.GetConn().Order("RAND()").Limit(count).Find(&tags).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "خطا در یافتن تگ ها",
			"error":   err.Error(),
		})
		return nil, err
	}
	return tags, nil
}
