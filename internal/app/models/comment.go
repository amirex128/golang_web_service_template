package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Comment struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	PostID    uint64 `json:"post_id"`
	Post      Post   `gorm:"foreignKey:post_id" json:"post"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Email     string `json:"email"`
	Approve   byte   `json:"accept"`
	CreatedAt string `json:"created_at"`
}
type CommentArr []Comment

func (s CommentArr) Len() int {
	return len(s)
}
func (s CommentArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s CommentArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Comment) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Comment) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initComment(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Comment{})
}

func (m *MysqlManager) CreateComment(c *gin.Context, dto DTOs.CreateComment) (err error) {
	comment := Comment{
		PostID:    dto.PostID,
		Title:     dto.Title,
		Body:      dto.Body,
		Email:     dto.Email,
		CreatedAt: utils.NowTime(),
	}
	err = m.GetConn().Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در ایجاد دیدگاه",
		})
		return err
	}
	return
}

func (m *MysqlManager) GetAllCommentWithPagination(c *gin.Context, dto DTOs.IndexComment) (pagination *DTOs.Pagination, err error) {
	conn := m.GetConn()
	var comments []Comment
	pagination = &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate(CommentTable, pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("title LIKE ?", "%"+dto.Search+"%").Order("id DESC")
	}
	err = conn.Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مشکلی در یافتن پست ها پیش آمده است",
		})
		return nil, err
	}
	pagination.Data = comments
	return pagination, nil
}
func (m *MysqlManager) GetAllComments(c *gin.Context) (comments []Comment, err error) {
	err = m.GetConn().Order("id DESC").Find(&comments).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در یافتن دیدگاه ها",
		})
		return nil, err
	}
	return comments, nil

}
func (m *MysqlManager) DeleteComment(c *gin.Context, id uint64) (err error) {
	conn := m.GetConn()
	err = conn.Where("id = ?", id).Delete(&Comment{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در حذف دیدگاه",
		})
		return err
	}
	return
}

func (m *MysqlManager) ApproveComment(c *gin.Context, id uint64) (err error) {
	conn := m.GetConn()
	err = conn.Model(&Comment{}).Where("id = ?", id).Update("approve", 1).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در تایید دیدگاه",
		})
		return err
	}
	return

}
