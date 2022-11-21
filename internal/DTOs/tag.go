package DTOs

type CreateTag struct {
	Name string `form:"name" json:"name" binding:"required,min=3,max=40"`
	Slug string `form:"slug" json:"slug" binding:"required,min=3,max=40"`
}

type IndexTag struct {
	Index
}

type AddTag struct {
	PostID uint64 `form:"post_id" json:"post_id" binding:"required,numeric"`
	TagID  uint64 `form:"tag_id" json:"tag_id" binding:"required,numeric"`
}
