package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

type Post struct {
	ID         uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Title      string     `json:"title"`
	Body       string     `json:"body"`
	Slug       string     `json:"slug"`
	GalleryID  *uint64    `gorm:"default:null" json:"gallery_id"`
	Gallery    *Gallery   `gorm:"foreignKey:gallery_id" json:"gallery"`
	UserID     *uint64    `json:"user_id"`
	User       *User      `gorm:"foreignKey:user_id" json:"user"`
	ShopID     *uint64    `gorm:"default:null" json:"shop_id"`
	Categories []Category `gorm:"many2many:category_post;" json:"categories"`
	Tags       []Tag      `gorm:"many2many:post_tag;" json:"tags"`
	Comments   []Comment  `gorm:"foreignKey:post_id" json:"comments"`
	CreatedAt  string     `json:"created_at"`
	UpdatedAt  string     `json:"updated_at"`
}

func InitPost(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Post{})
	for i := 0; i < 10; i++ {
		manager.CreatePost(&gin.Context{}, context.Background(), DTOs.CreatePost{
			Title:     "آموزش برنامه نویس گولنگ" + fmt.Sprintf("%d", i),
			Body:      "این یک پست آموزشی برنامه نویسی گولنگ است" + fmt.Sprintf("%d", i),
			Slug:      "amoozesh-barnamenevis-golang" + fmt.Sprintf("%d", i),
			GalleryID: 1,
		}, 1)
	}
}
func (m *MysqlManager) CheckSlug(c *gin.Context, ctx context.Context, slug string) (err error) {
	span, ctx := apm.StartSpan(ctx, "CheckSlug", "model")
	defer span.End()
	rowsAffected := m.GetConn().Where("slug = ?", slug).First(&Post{}).RowsAffected
	if rowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "این نامک قبلا استفاده شده است",
		})
		return errors.New("slug not valid")
	}
	return
}

func (m *MysqlManager) CreatePost(c *gin.Context, ctx context.Context, dto DTOs.CreatePost, userID uint64) (err error) {
	span, ctx := apm.StartSpan(ctx, "CreatePost", "model")
	defer span.End()
	post := Post{
		Title: dto.Title,
		Body:  dto.Body,
		Slug:  dto.Slug,
		GalleryID: func() *uint64 {
			if dto.GalleryID != 0 {
				return &dto.GalleryID
			}
			return nil
		}(),
		UserID:    &userID,
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

func (m *MysqlManager) UpdatePost(c *gin.Context, ctx context.Context, dto DTOs.UpdatePost) (err error) {
	span, ctx := apm.StartSpan(ctx, "UpdatePost", "model")
	defer span.End()
	post := Post{}
	err = m.GetConn().Where("id = ?", dto.ID).First(&post).Error
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
	if dto.GalleryID != 0 {
		post.GalleryID = &dto.GalleryID
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

func (m *MysqlManager) DeletePost(c *gin.Context, ctx context.Context, postID uint64) (err error) {
	span, ctx := apm.StartSpan(ctx, "DeletePost", "model")
	defer span.End()
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

func (m *MysqlManager) FindPostByID(c *gin.Context, ctx context.Context, postID uint64) (post Post, err error) {
	span, ctx := apm.StartSpan(ctx, "FindPostByID", "model")
	defer span.End()
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

func (m *MysqlManager) FindPostBySlug(slug string, ctx context.Context) (Post, error) {
	span, ctx := apm.StartSpan(ctx, "FindPostBySlug", "model")
	defer span.End()
	var post Post
	err := m.GetConn().Where("slug = ?", slug).Preload("User").Preload("Categories").Preload("Gallery").First(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (m *MysqlManager) GetAllPostWithPagination(c *gin.Context, ctx context.Context, dto DTOs.IndexPost) (pagination *DTOs.Pagination, err error) {
	span, ctx := apm.StartSpan(ctx, "GetAllPostWithPagination", "model")
	defer span.End()
	conn := m.GetConn()
	var posts []Post
	pagination = &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("posts", pagination, conn)).Preload("User").Preload("Categories").Preload("Gallery").Order("id DESC")
	if dto.Search != "" {
		conn = conn.Where("title LIKE ?", "%"+dto.Search+"%")
	}
	err = conn.Where("shop_id = ? ", dto.ShopID).Find(&posts).Error
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

func (m *MysqlManager) RandomPost(c *gin.Context, ctx context.Context, count int) (posts []Post, err error) {
	span, ctx := apm.StartSpan(ctx, "RandomPost", "model")
	defer span.End()
	err = m.GetConn().Order("RAND()").Limit(count).Preload("User").Preload("Gallery").Find(&posts).Error
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

func (m *MysqlManager) GetLastPost(c *gin.Context, ctx context.Context, count int) (posts []Post, err error) {
	span, ctx := apm.StartSpan(ctx, "GetLastPost", "model")
	defer span.End()
	err = m.GetConn().Preload("User").Preload("Gallery").Order("id DESC").Limit(count).Find(&posts).Error
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
