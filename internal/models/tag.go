package models

import (
	"context"
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
	"strconv"
)

type Tag struct {
	ID       uint64    `json:"id"`
	Name     string    `json:"name"`
	Slug     string    `json:"slug"`
	Posts    []Post    `gorm:"many2many:post_tag;" json:"posts"`
	Products []Product `gorm:"many2many:product_tag;" json:"products"`
}

func initTag(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Tag{})
	for i := 0; i < 100; i++ {
		manager.CreateTag(&gin.Context{}, context.Background(), DTOs.CreateTag{
			Name: "تگ شماره" + strconv.Itoa(i),
			Slug: "tag" + strconv.Itoa(i),
		})
	}
}

func (m *MysqlManager) CreateTag(c *gin.Context, ctx context.Context, dto DTOs.CreateTag) (err error) {
	span, ctx := apm.StartSpan(ctx, "CreateTag", "model")
	defer span.End()
	tag := Tag{
		Name: dto.Name,
		Slug: dto.Slug,
	}
	err = m.GetConn().Create(&tag).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "خطا در ایجاد تگ",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return
}

func (m *MysqlManager) GetAllTagsWithPagination(c *gin.Context, ctx context.Context, dto DTOs.IndexTag) (pagination *DTOs.Pagination, err error) {
	span, ctx := apm.StartSpan(ctx, "GetAllTagsWithPagination", "model")
	defer span.End()
	conn := m.GetConn()
	var tags []Tag
	pagination = &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("tags", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%").Order("id DESC")
	}
	err = conn.Find(&tags).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در یافتن پست ها پیش آمده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return nil, err
	}
	pagination.Data = tags
	return pagination, nil
}

func (m *MysqlManager) DeleteTag(c *gin.Context, ctx context.Context, id uint64) (err error) {
	span, ctx := apm.StartSpan(ctx, "DeleteTag", "model")
	defer span.End()
	err = m.GetConn().Where("id = ?", id).Delete(&Tag{}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "خطا در حذف تگ",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return
}

func (m *MysqlManager) AddTag(c *gin.Context, ctx context.Context, dto DTOs.AddTag) (err error) {
	span, ctx := apm.StartSpan(ctx, "AddTag", "model")
	defer span.End()
	if dto.Type == "post" {
		err = m.GetConn().Model(&Post{ID: dto.PostID}).Association("Tags").Append(&Tag{ID: dto.TagID})

	} else {
		err = m.GetConn().Model(&Product{ID: dto.ProductID}).Association("Tags").Append(&Tag{ID: dto.TagID})
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "خطا در افزودن تگ",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return
}

func (m *MysqlManager) RandomTags(c *gin.Context, ctx context.Context, count int) (tags []*Tag, err error) {
	span, ctx := apm.StartSpan(ctx, "RandomTags", "model")
	defer span.End()
	err = m.GetConn().Order("RAND()").Limit(count).Find(&tags).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "خطا در یافتن تگ ها",
			"error":   err.Error(),
			"type":    "model",
		})
		return nil, err
	}
	return tags, nil
}

func (m *MysqlManager) GetAllTagPostWithPagination(c *gin.Context, ctx context.Context, dto DTOs.IndexPost, tagID uint64) (pagination *DTOs.Pagination, err error) {
	span, ctx := apm.StartSpan(ctx, "GetAllTagPostWithPagination", "model")
	defer span.End()
	conn := m.GetConn()
	var posts []Post
	pagination = &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("posts", pagination, conn)).Where("id IN (?)", conn.Table("post_tag").Where("tag_id = ?", tagID).Select("post_id")).Preload("User").Preload("Categories").Order("id DESC")
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

func (m *MysqlManager) FindTagBySlug(c *gin.Context, ctx context.Context, slug string) (tag *Tag, err error) {
	span, ctx := apm.StartSpan(ctx, "FindTagBySlug", "model")
	defer span.End()
	tag = &Tag{}
	err = m.GetConn().Where("slug = ?", slug).First(tag).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "خطا در یافتن تگ",
			"error":   err.Error(),
			"type":    "model",
		})
		return nil, err
	}
	return tag, nil
}
