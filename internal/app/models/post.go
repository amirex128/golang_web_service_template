package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Post struct {
	ID         uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Title      string     `json:"title"`
	Body       string     `json:"body"`
	Slug       string     `json:"slug"`
	UserID     uint64     `json:"user_id"`
	User       User       `gorm:"foreignKey:user_id" json:"user"`
	Categories []Category `gorm:"many2many:category_post;" json:"categories"`
	Tags       []Tag      `gorm:"many2many:post_tag;" json:"tags"`
	Comments   []Comment  `gorm:"foreignKey:post_id" json:"comments"`
	CreatedAt  string     `json:"created_at"`
	UpdatedAt  string     `json:"updated_at"`
	Galleries  []Gallery  `gorm:"polymorphic:Owner;"`
}
type PostArr []Post

func (s PostArr) Len() int {
	return len(s)
}
func (s PostArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s PostArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Post) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Post) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func InitPost(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Post{})
	for i := 0; i < 10; i++ {
		manager.CreatePost(&gin.Context{}, DTOs.CreatePost{
			Title:     "آموزش برنامه نویس گولنگ" + fmt.Sprintf("%d", i),
			Body:      "این یک پست آموزشی برنامه نویسی گولنگ است" + fmt.Sprintf("%d", i),
			Slug:      "amoozesh-barnamenevis-golang" + fmt.Sprintf("%d", i),
			CreatedAt: utils.NowTime(),
			UpdatedAt: utils.NowTime(),
		}, 1)
	}
}
func (m *MysqlManager) CheckSlug(c *gin.Context, slug string) (err error) {
	rowsAffected := m.GetConn().Where("slug = ?", slug).First(&Post{}).RowsAffected
	if rowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "این نامک قبلا استفاده شده است",
		})
		return errors.New("slug not valid")
	}
	return
}

func (m *MysqlManager) CreatePost(c *gin.Context, dto DTOs.CreatePost, userID uint64) (err error) {
	post := Post{
		Title:     dto.Title,
		Body:      dto.Body,
		Slug:      dto.Slug,
		UserID:    userID,
		CreatedAt: utils.NowTime(),
		UpdatedAt: utils.NowTime(),
	}
	err = m.GetConn().Create(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در ایجاد پست پیش آمده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return
}
func (m *MysqlManager) UpdatePost(c *gin.Context, dto DTOs.UpdatePost, postID uint64) (err error) {
	post := Post{}
	err = m.GetConn().Where("id = ?", postID).First(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در ویرایش پست پیش آمده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	if dto.Title != "" {
		post.Title = dto.Title
	}
	if dto.Body != "" {
		post.Body = dto.Body
	}
	if dto.Slug != "" {
		post.Slug = dto.Slug
	}
	post.UpdatedAt = utils.NowTime()
	err = m.GetConn().Save(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در ویرایش پست پیش آمده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return
}

func (m *MysqlManager) DeletePost(c *gin.Context, postID uint64) (err error) {
	err = m.GetConn().Where("id = ?", postID).Delete(&Post{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در حذف پست پیش آمده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return
	}
	return
}

func (m *MysqlManager) FindPostByID(c *gin.Context, postID uint64) (post Post, err error) {
	err = m.GetConn().Where("id = ?", postID).Preload("User").Preload("Categories").First(&post).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در یافتن پست پیش آمده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return
	}
	return
}
func (m *MysqlManager) FindPostBySlug(slug string) (Post, error) {
	var post Post
	err := m.GetConn().Where("slug = ?", slug).Preload("User").Preload("Categories").First(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (m *MysqlManager) GetAllPostWithPagination(c *gin.Context, dto DTOs.IndexPost) (pagination *DTOs.Pagination, err error) {
	conn := m.GetConn()
	var posts []Post
	pagination = &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("posts", pagination, conn)).Preload("User").Preload("Categories").Order("id DESC")
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

func (m *MysqlManager) RandomPost(c *gin.Context, count int) (posts []Post, err error) {
	err = m.GetConn().Order("RAND()").Limit(count).Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "مشکلی در یافتن پست ها پیش آمده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return nil, err
	}
	return posts, nil
}
func (m *MysqlManager) GetLastPost(c *gin.Context, count int) (posts []Post, err error) {
	err = m.GetConn().Preload("User").Order("id DESC").Limit(count).Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "مشکلی در یافتن پست ها پیش آمده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return nil, err
	}
	return posts, nil
}
