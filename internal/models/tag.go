package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/brianvoe/gofakeit/v6"
	"go.elastic.co/apm/v2"
)

type Tag struct {
	ID       uint64    `json:"id"`
	UserID   *uint64   `json:"user_id"`
	Name     string    `json:"name"`
	Slug     string    `json:"slug"`
	Posts    []Post    `gorm:"many2many:post_tag;" json:"posts"`
	Products []Product `gorm:"many2many:product_tag;" json:"products"`
}

func initTag(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&Tag{}) {
		manager.GetConn().Migrator().CreateTable(&Tag{})

		for i := 0; i < 100; i++ {
			model := new(DTOs.CreateTag)
			gofakeit.Struct(model)

			manager.CreateTag(*model)
		}
	}

}

func (m *MysqlManager) CreateTag(dto DTOs.CreateTag) (*Tag, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreateTag", "model")
	defer span.End()
	tag := &Tag{
		Name:   dto.Name,
		UserID: utils.GetUserID(m.Ctx),
		Slug:   dto.Slug,
	}
	err := m.GetConn().Create(tag).Error
	if err != nil {
		return tag, errorx.New("خطا در ایجاد تگ", "model", err)
	}
	return tag, nil
}

func (m *MysqlManager) GetAllTagsWithPagination(dto DTOs.IndexTag) (pagination *DTOs.Pagination, err error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetAllTagsWithPagination", "model")
	defer span.End()
	conn := m.GetConn()
	var tags []Tag
	pagination = &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("tags", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%")
	}
	err = conn.Order("id DESC").Find(&tags).Error
	if err != nil {
		return nil, errorx.New("مشکلی در یافتن پست ها پیش آمده است", "model", err)
	}
	pagination.Data = tags
	return pagination, nil
}

func (m *MysqlManager) DeleteTag(id uint64) (err error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:DeleteTag", "model")
	defer span.End()
	tag := &Tag{}
	err = m.GetConn().Where("id = ?", id).First(tag).Error
	if err != nil {
		return errorx.New("خطا در یافتن تگ", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, tag.UserID); err != nil {
		return err
	}
	err = m.GetConn().Delete(tag).Error
	if err != nil {
		return errorx.New("خطا در حذف تگ", "model", err)
	}
	return
}

func (m *MysqlManager) AddTag(dto DTOs.AddTag) (err error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:AddTag", "model")
	defer span.End()
	if dto.Type == "post" {
		err = m.GetConn().Model(&Post{ID: dto.PostID}).Association("Tags").Append(&Tag{ID: dto.TagID})

	} else {
		err = m.GetConn().Model(&Product{ID: dto.ProductID}).Association("Tags").Append(&Tag{ID: dto.TagID})
	}
	if err != nil {
		return errorx.New("خطا در افزودن تگ", "model", err)
	}
	return
}

func (m *MysqlManager) RandomTags(count int) (tags []*Tag, err error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:RandomTags", "model")
	defer span.End()
	err = m.GetConn().Order("RAND()").Limit(count).Find(&tags).Error
	if err != nil {
		return nil, errorx.New("خطا در یافتن تگ ها", "model", err)
	}
	return tags, nil
}

func (m *MysqlManager) GetAllTagPostWithPagination(dto DTOs.IndexPost, tagID uint64) (pagination *DTOs.Pagination, err error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetAllTagPostWithPagination", "model")
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
		return nil, errorx.New("مشکلی در یافتن پست ها پیش آمده است", "model", err)
	}
	pagination.Data = posts
	return pagination, nil
}

func (m *MysqlManager) FindTagBySlug(slug string) (tag *Tag, err error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindTagBySlug", "model")
	defer span.End()
	tag = &Tag{}
	err = m.GetConn().Where("slug = ?", slug).First(tag).Error
	if err != nil {
		return nil, errorx.New("خطا در یافتن تگ", "model", err)
	}
	return tag, nil
}
