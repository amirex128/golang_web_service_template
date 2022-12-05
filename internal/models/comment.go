package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/brianvoe/gofakeit/v6"
	"go.elastic.co/apm/v2"
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
	if !manager.GetConn().Migrator().HasTable(&Comment{}) {
		manager.GetConn().AutoMigrate(&Comment{})
		for i := 0; i < 100; i++ {
			model := new(DTOs.CreateComment)
			gofakeit.Struct(model)

			manager.CreateComment(*model)
		}
	}

}

func (m *MysqlManager) CreateComment(dto DTOs.CreateComment) (*Comment, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreateComment", "model")
	defer span.End()
	comment := &Comment{
		PostID:    dto.PostID,
		Name:      dto.Name,
		Body:      dto.Body,
		Email:     dto.Email,
		CreatedAt: utils.NowTime(),
	}
	err := m.GetConn().Create(comment).Error
	if err != nil {
		return comment, errorx.New("خطا در ایجاد دیدگاه", "model", err)
	}
	return comment, nil
}

func (m *MysqlManager) GetAllCommentWithPagination(dto DTOs.IndexComment) (pagination *DTOs.Pagination, err error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetAllCommentWithPagination", "model")
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
		return nil, errorx.New("مشکلی در یافتن پست ها پیش آمده است", "model", err)
	}
	pagination.Data = comments
	return pagination, nil
}

func (m *MysqlManager) GetAllComments(postID uint64) ([]*Comment, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetAllComments", "model")
	defer span.End()
	comments := make([]*Comment, 0)
	err := m.GetConn().Where("post_id = ?", postID).Where("approve = ?", true).Order("id DESC").Find(&comments).Error
	if err != nil {
		return nil, errorx.New("خطا در یافتن دیدگاه ها", "model", err)
	}
	return comments, nil

}

func (m *MysqlManager) DeleteComment(id uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:DeleteComment", "model")
	defer span.End()
	conn := m.GetConn()
	err := conn.Where("id = ?", id).Delete(&Comment{}).Error
	if err != nil {
		return errorx.New("خطا در حذف دیدگاه", "model", err)
	}
	return nil
}

func (m *MysqlManager) ApproveComment(id uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:ApproveComment", "model")
	defer span.End()
	conn := m.GetConn()
	err := conn.Model(&Comment{}).Where("id = ?", id).Update("approve", true).Error
	if err != nil {
		return errorx.New("خطا در تایید دیدگاه", "model", err)
	}
	return nil

}
func (m *MysqlManager) FindCommentByID(id uint64) (*Comment, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindCommentByID", "model")
	defer span.End()
	menu := &Comment{}
	err := m.GetConn().Where("id = ?", id).First(menu).Error
	if err != nil {
		return menu, errorx.New("دیدگاه مورد نظر یافت نشد", "model", err)
	}
	return menu, nil
}
