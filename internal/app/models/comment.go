package models

import (
	"github.com/amirex128/selloora_backend/internal/app/DTOs"
	"github.com/amirex128/selloora_backend/internal/app/utils"
	"context"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

type Comment struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	PostID    uint64 `json:"post_id"`
	Post      Post   `gorm:"foreignKey:post_id" json:"post"`
	Name      string `json:"title"`
	EmailHash string `gorm:"-:all" json:"email_hash"`
	Body      string `json:"body"`
	Email     string `json:"email"`
	Approve   bool   `json:"accept"`
	CreatedAt string `json:"created_at"`
}

func InitComment(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Comment{})
	for i := 0; i < 10; i++ {
		manager.CreateComment(&gin.Context{}, context.Background(), DTOs.CreateComment{
			PostID: 1,
			Name:   "test test test",
			Body:   "test test test",
			Email:  "test@test.com",
		})
	}
}

func (m *MysqlManager) CreateComment(c *gin.Context, ctx context.Context, dto DTOs.CreateComment) (err error) {
	span, ctx := apm.StartSpan(ctx, "CreateComment", "model")
	defer span.End()
	comment := Comment{
		PostID:    dto.PostID,
		Name:      dto.Name,
		Body:      dto.Body,
		Email:     dto.Email,
		CreatedAt: utils.NowTime(),
	}
	err = m.GetConn().Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در ایجاد دیدگاه",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return
}

func (m *MysqlManager) GetAllCommentWithPagination(c *gin.Context, ctx context.Context, dto DTOs.IndexComment) (pagination *DTOs.Pagination, err error) {
	span, ctx := apm.StartSpan(ctx, "GetAllCommentWithPagination", "model")
	defer span.End()
	conn := m.GetConn()
	var comments []Comment
	pagination = &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("comments", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("title LIKE ?", "%"+dto.Search+"%").Order("id DESC")
	}
	err = conn.Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در یافتن پست ها پیش آمده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return nil, err
	}
	pagination.Data = comments
	return pagination, nil
}

func (m *MysqlManager) GetAllComments(c *gin.Context, ctx context.Context, postID uint64) (comments []*Comment, err error) {
	span, ctx := apm.StartSpan(ctx, "GetAllComments", "model")
	defer span.End()
	err = m.GetConn().Where("post_id = ?", postID).Where("approve = ?", true).Order("id DESC").Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در یافتن دیدگاه ها",
			"error":   err.Error(),
			"type":    "model",
		})
		return nil, err
	}
	return comments, nil

}

func (m *MysqlManager) DeleteComment(c *gin.Context, ctx context.Context, id uint64) (err error) {
	span, ctx := apm.StartSpan(ctx, "DeleteComment", "model")
	defer span.End()
	conn := m.GetConn()
	err = conn.Where("id = ?", id).Delete(&Comment{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در حذف دیدگاه",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return
}

func (m *MysqlManager) ApproveComment(c *gin.Context, ctx context.Context, id uint64) (err error) {
	span, ctx := apm.StartSpan(ctx, "ApproveComment", "model")
	defer span.End()
	conn := m.GetConn()
	err = conn.Model(&Comment{}).Where("id = ?", id).Update("approve", true).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در تایید دیدگاه",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return

}
