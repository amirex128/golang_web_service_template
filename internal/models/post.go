package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/brianvoe/gofakeit/v6"
	"go.elastic.co/apm/v2"
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
	if !manager.GetConn().Migrator().HasTable(&Post{}) {
		manager.GetConn().Migrator().CreateTable(&Post{})
		for i := 0; i < 100; i++ {
			model := new(DTOs.CreatePost)
			gofakeit.Struct(model)

			manager.CreatePost(*model)
		}

	}

}
func (m *MysqlManager) CheckSlug(slug string) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CheckSlug", "model")
	defer span.End()
	rowsAffected := m.GetConn().Where("slug = ?", slug).First(&Post{}).RowsAffected
	if rowsAffected > 0 {
		return errorx.New("این نامک قبلا استفاده شده است", "model", nil)
	}
	return nil
}

func (m *MysqlManager) CreatePost(dto DTOs.CreatePost) (*Post, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreatePost", "model")
	defer span.End()
	userID := GetUserID(m.Ctx)
	post := &Post{
		Title: dto.Title,
		Body:  dto.Body,
		Slug:  dto.Slug,
		GalleryID: func() *uint64 {
			if dto.GalleryID != 0 {
				return &dto.GalleryID
			}
			return nil
		}(),
		UserID:    userID,
		CreatedAt: utils.NowTime(),
		UpdatedAt: utils.NowTime(),
	}
	err := m.GetConn().Create(&post).Error
	if err != nil {
		return post, errorx.New("مشکلی در ایجاد پست پیش آمده است", "model", err)
	}
	return post, nil
}

func (m *MysqlManager) UpdatePost(dto DTOs.UpdatePost) (err error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:UpdatePost", "model")
	defer span.End()
	post := Post{}
	err = m.GetConn().Where("id = ?", dto.ID).First(&post).Error
	if err != nil {
		return errorx.New("مشکلی در ویرایش پست پیش آمده است", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, post.UserID); err != nil {
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
		return errorx.New("مشکلی در ویرایش پست پیش آمده است", "model", err)
	}
	return
}

func (m *MysqlManager) DeletePost(postID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:DeletePost", "model")
	defer span.End()
	post := &Post{}
	err := m.GetConn().Where("id = ?", postID).First(&post).Error
	if err != nil {
		return errorx.New("مشکلی در حذف پست پیش آمده است", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, post.UserID); err != nil {
		return err
	}
	err = m.GetConn().Delete(post).Error
	if err != nil {
		return errorx.New("مشکلی در حذف پست پیش آمده است", "model", err)
	}
	return nil
}

func (m *MysqlManager) FindPostByID(postID uint64) (*Post, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindPostByID", "model")
	defer span.End()
	post := &Post{}
	err := m.GetConn().Where("id = ?", postID).Preload("User").Preload("Categories").First(post).Error
	if err != nil {
		return nil, errorx.New("مشکلی در یافتن پست پیش آمده است", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, post.UserID); err != nil {
		return nil, err
	}
	return post, nil
}

func (m *MysqlManager) FindPostBySlug(slug string) (*Post, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindPostBySlug", "model")
	defer span.End()
	var post *Post
	err := m.GetConn().Where("slug = ?", slug).Preload("User").Preload("Categories").Preload("Gallery").First(&post).Error
	if err != nil {
		return nil, errorx.New("مشکلی در یافتن پست پیش آمده است", "model", err)
	}
	return post, nil
}

func (m *MysqlManager) GetAllPostWithPagination(dto DTOs.IndexPost) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetAllPostWithPagination", "model")
	defer span.End()
	conn := m.GetConn()
	var posts []Post
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("posts", pagination, conn)).Preload("User").Preload("Categories").Preload("Gallery").Order("id DESC")
	if dto.Search != "" {
		conn = conn.Where("title LIKE ?", "%"+dto.Search+"%")
	}
	err := conn.Where("user_id = ?", GetUserID(m.Ctx)).Where("shop_id = ? ", dto.ShopID).Find(&posts).Error
	if err != nil {
		return nil, errorx.New("مشکلی در یافتن پست ها پیش آمده است", "model", err)
	}
	pagination.Data = posts
	return pagination, nil
}

func (m *MysqlManager) RandomPost(count int) ([]*Post, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:RandomPost", "model")
	defer span.End()
	posts := make([]*Post, 0)
	err := m.GetConn().Order("RAND()").Limit(count).Preload("User").Preload("Gallery").Find(posts).Error
	if err != nil {
		return nil, errorx.New("مشکلی در یافتن پست ها پیش آمده است", "model", err)
	}
	return posts, nil
}

func (m *MysqlManager) GetLastPost(count int) ([]*Post, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:GetLastPost", "model")
	defer span.End()
	posts := make([]*Post, 0)
	err := m.GetConn().Preload("User").Preload("Gallery").Order("id DESC").Limit(count).Find(&posts).Error
	if err != nil {
		return nil, errorx.New("مشکلی در یافتن پست ها پیش آمده است", "model", err)
	}
	return posts, nil
}
