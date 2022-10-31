package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"net/http"
)

func createPost(c *gin.Context) {
	dto, err := validations.CreatePost(c)
	if err != nil {
		return
	}
	userID := utils.GetUser(c)
	dto.Slug = slug.MakeLang(dto.Slug, "en")
	err = models.NewMainManager().CheckSlug(c, dto.Slug)
	if err != nil {
		return
	}

	err = models.NewMainManager().CreatePost(c, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله شما با موفقیت ایجاد شد",
	})
}

func updatePost(c *gin.Context) {
	dto, err := validations.UpdatePost(c)
	if err != nil {
		return
	}
	postID := c.Param("id")
	if dto.Slug != "" {
		dto.Slug = slug.MakeLang(dto.Slug, "en")
		err = models.NewMainManager().CheckSlug(c, dto.Slug)
		if err != nil {
			return
		}
	}

	err = models.NewMainManager().UpdatePost(c, dto, utils.StringToUint64(postID))
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله با موفقیت ویرایش شد",
	})
}

func showPost(c *gin.Context) {
	postID := c.Param("id")
	post, err := models.NewMainManager().FindPostByID(c, utils.StringToUint64(postID))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func indexPost(c *gin.Context) {
	dto, err := validations.IndexPost(c)
	if err != nil {
		return
	}
	posts, err := models.NewMainManager().GetAllPostWithPagination(c, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}

func deletePost(c *gin.Context) {
	postID := c.Param("id")
	err := models.NewMainManager().DeletePost(c, utils.StringToUint64(postID))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله با موفقیت حذف شد",
	})
}
