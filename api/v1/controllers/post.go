package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"go.elastic.co/apm/v2"
	"net/http"
)

func createPostAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createPostAdmin", "request")
	defer span.End()
	dto, err := validations.CreatePost(c)
	if err != nil {
		return
	}
	userID := models.GetUser(c)
	dto.Slug = slug.MakeLang(dto.Slug, "en")
	err = models.NewMainManager().CheckSlug(c, ctx, dto.Slug)
	if err != nil {
		return
	}

	err = models.NewMainManager().CreatePost(c, ctx, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله شما با موفقیت ایجاد شد",
	})
}

func updatePostAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updatePostAdmin", "request")
	defer span.End()
	dto, err := validations.UpdatePost(c)
	if err != nil {
		return
	}
	postID := c.Param("id")
	if dto.Slug != "" {
		dto.Slug = slug.MakeLang(dto.Slug, "en")
		err = models.NewMainManager().CheckSlug(c, ctx, dto.Slug)
		if err != nil {
			return
		}
	}

	err = models.NewMainManager().UpdatePost(c, ctx, dto, utils.StringToUint64(postID))
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله با موفقیت ویرایش شد",
	})
}

func showPostAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "showPostAdmin", "request")
	defer span.End()
	postID := c.Param("id")
	post, err := models.NewMainManager().FindPostByID(c, ctx, utils.StringToUint64(postID))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func indexPostAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexPostAdmin", "request")
	defer span.End()
	dto, err := validations.IndexPost(c)
	if err != nil {
		return
	}
	posts, err := models.NewMainManager().GetAllPostWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}

func deletePostAdmin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deletePostAdmin", "request")
	defer span.End()
	postID := c.Param("id")
	err := models.NewMainManager().DeletePost(c, ctx, utils.StringToUint64(postID))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله با موفقیت حذف شد",
	})
}
