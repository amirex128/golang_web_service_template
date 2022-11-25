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

// UpdatePost
// @Summary ویرایش مقاله
// @description فروشگاه ها میتوانند برای خود مقاله بسازند تا در سئو بهتر باشند
// @Tags post
// @Router       /user/post/update/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	false "شناسه " SchemaExample(1)
// @Param	message	 body   DTOs.UpdatePost  	true "ورودی"
func UpdatePost(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:updatePost", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.UpdatePost(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	if dto.Slug != "" {
		dto.Slug = slug.MakeLang(dto.Slug, "en")
		err = models.NewMysqlManager(c).CheckSlug(dto.Slug)
		if err != nil {
			return
		}
	}

	err = models.NewMysqlManager(c).UpdatePost(dto)
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله با موفقیت ویرایش شد",
	})
}
