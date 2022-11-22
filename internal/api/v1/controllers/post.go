package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
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
	span, ctx := apm.StartSpan(c.Request.Context(), "createPost", "request")
	defer span.End()
	dto, err := validations.CreatePost(c)
	if err != nil {
		return
	}
	userID := models.GetUser(c)
	dto.Slug = slug.MakeLang(dto.Slug, "en")
	err = models.NewMysqlManager(ctx).CheckSlug(c, ctx, dto.Slug)
	if err != nil {
		return
	}

	err = models.NewMysqlManager(ctx).CreatePost(c, ctx, dto, *userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله شما با موفقیت ایجاد شد",
	})
}

// UpdatePost
// @Summary ویرایش مقاله
// @description فروشگاه ها میتوانند برای خود مقاله بسازند تا در سئو بهتر باشند
// @Tags post
// @Router       /user/post/update [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.UpdatePost  	true "ورودی"
func UpdatePost(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updatePost", "request")
	defer span.End()
	dto, err := validations.UpdatePost(c)
	if err != nil {
		return
	}
	if dto.Slug != "" {
		dto.Slug = slug.MakeLang(dto.Slug, "en")
		err = models.NewMysqlManager(ctx).CheckSlug(c, ctx, dto.Slug)
		if err != nil {
			return
		}
	}

	err = models.NewMysqlManager(ctx).UpdatePost(c, ctx, dto)
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله با موفقیت ویرایش شد",
	})
}

// ShowPost
// @Summary نمایش مقاله
// @description فروشگاه ها میتوانند برای خود مقاله بسازند تا در سئو بهتر باشند
// @Tags post
// @Router       /user/post/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه " SchemaExample(1)
func ShowPost(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "showPost", "request")
	defer span.End()
	postID := c.Param("id")
	post, err := models.NewMysqlManager(ctx).FindPostByID(c, ctx, utils.StringToUint64(postID))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

// IndexPost
// @Summary لیست مقالات
// @description فروشگاه ها میتوانند برای خود مقاله بسازند تا در سئو بهتر باشند
// @Tags post
// @Router       /user/post [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexPost(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexPost", "request")
	defer span.End()

	dto, err := validations.IndexPost(c)
	if err != nil {
		return
	}
	posts, err := models.NewMysqlManager(ctx).GetAllPostWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}

// DeletePost
// @Summary حدف مقاله
// @description فروشگاه ها میتوانند برای خود مقاله بسازند تا در سئو بهتر باشند
// @Tags post
// @Router       /user/post/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه " SchemaExample(1)
func DeletePost(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deletePost", "request")
	defer span.End()
	postID := c.Param("id")
	err := models.NewMysqlManager(ctx).DeletePost(c, ctx, utils.StringToUint64(postID))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله با موفقیت حذف شد",
	})
}
