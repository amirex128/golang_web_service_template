package admin

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"net/http"
)

func CreatePost(c *gin.Context) {
	dto, err := validations.CreatePost(c)
	if err != nil {
		return
	}
	userID := utils.GetUser(c)
	dto.Slug = slug.MakeLang(dto.Slug, "en")
	image, err := utils.UploadImage(c, dto.Thumbnail, "post")
	if err != nil {
		return
	}
	dto.ThumbnailPath = image
	err = models.NewMainManager().CreatePost(c, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله شما با موفقیت ایجاد شد",
	})
}

func UpdatePost(c *gin.Context) {
	dto, err := validations.UpdatePost(c)
	if err != nil {
		return
	}
	postID := c.Param("id")
	dto.Slug = slug.MakeLang(dto.Slug, "en")
	image, err := utils.UploadImage(c, dto.Thumbnail, "post")
	if err != nil {
		return
	}
	dto.ThumbnailPath = image
	utils.RemoveImages([]string{dto.ThumbnailRemove})
	err = models.NewMainManager().UpdatePost(c, dto, utils.StringToUint64(postID))
}

func ShowPost(c *gin.Context) {
	postID := c.Param("id")
	post, err := models.NewMainManager().FindPostByID(c, utils.StringToUint64(postID))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func IndexPost(c *gin.Context) {
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

func DeletePost(c *gin.Context) {
	postID := c.Param("id")
	err := models.NewMainManager().DeletePost(c, utils.StringToUint64(postID))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "مقاله با موفقیت حذف شد",
	})
}
