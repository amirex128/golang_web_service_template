package post

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreatePost
// @Summary ایجاد مقاله
// @description فروشگاه ها میتوانند برای خود مقاله بسازند تا در سئو بهتر باشند
// @Tags post
// @Router       /user/post/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreatePost  	true "ورودی"
func CreatePost(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createPost", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreatePost(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	dto.Slug = slug.MakeLang(dto.Slug, "en")
	err = models.NewMysqlManager(c).CheckSlug(dto.Slug)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	post, err := models.NewMysqlManager(c).CreatePost(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله شما با موفقیت ایجاد شد",
		"data":    post,
	})
}
