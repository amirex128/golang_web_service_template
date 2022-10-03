package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	siteName = " | سلورا دستیار فروش اینستاگرام شما"
)

func indexLanding(c *gin.Context) {
	c.Set("template", "index.html")
	c.Set("data", map[string]interface{}{
		"title": "صفحه اصلی" + siteName,
	})
}

func blogLanding(c *gin.Context) {
	dto, err := validations.IndexPost(c)
	if err != nil {
		return
	}
	posts, err := models.NewMainManager().GetAllPostWithPagination(c, dto)
	if err != nil {
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMainManager().RandomTags(c, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMainManager().GetLevel1Categories(c)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMainManager().RandomPost(c, 6)
	if err != nil {
		return
	}
	for i := range randomPosts {
		randomPosts[i].CreatedAt = func() string {
			ago := utils.DateAgo(randomPosts[i].CreatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
		randomPosts[i].UpdatedAt = func() string {
			ago := utils.DateAgo(randomPosts[i].UpdatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
	}
	c.Set("template", "blog.html")
	c.Set("data", map[string]interface{}{
		"title":        "بلاگ" + siteName,
		"posts":        posts,
		"random_posts": randomPosts,
		"categories":   categories,
		"tags":         tags,
	})
}

func categoryLanding(c *gin.Context) {
	dto, err := validations.IndexPost(c)
	if err != nil {
		return
	}
	categoryID := c.Param("id")
	category, err := models.NewMainManager().FindCategoryByID(c, utils.StringToUint32(categoryID))
	if err != nil {
		return
	}
	posts, err := models.NewMainManager().GetAllCategoryPostWithPagination(c, dto, utils.StringToUint32(categoryID))
	if err != nil {
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMainManager().RandomTags(c, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMainManager().GetLevel1Categories(c)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMainManager().RandomPost(c, 6)
	if err != nil {
		return
	}
	for i := range randomPosts {
		randomPosts[i].CreatedAt = func() string {
			ago := utils.DateAgo(randomPosts[i].CreatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
		randomPosts[i].UpdatedAt = func() string {
			ago := utils.DateAgo(randomPosts[i].UpdatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
	}
	c.Set("template", "category.html")
	c.Set("data", map[string]interface{}{
		"title":        category.Name + siteName,
		"category":     category,
		"posts":        posts,
		"random_posts": randomPosts,
		"categories":   categories,
		"tags":         tags,
	})
}

func tagLanding(c *gin.Context) {
	dto, err := validations.IndexPost(c)
	if err != nil {
		return
	}
	tagSlug := c.Param("slug")
	tag, err := models.NewMainManager().FindTagBySlug(c, tagSlug)
	if err != nil {
		return
	}
	posts, err := models.NewMainManager().GetAllTagPostWithPagination(c, dto, tag.ID)
	if err != nil {
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMainManager().RandomTags(c, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMainManager().GetLevel1Categories(c)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMainManager().RandomPost(c, 6)
	if err != nil {
		return
	}
	for i := range randomPosts {
		randomPosts[i].CreatedAt = func() string {
			ago := utils.DateAgo(randomPosts[i].CreatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
		randomPosts[i].UpdatedAt = func() string {
			ago := utils.DateAgo(randomPosts[i].UpdatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
	}
	c.Set("template", "tag.html")
	c.Set("data", map[string]interface{}{
		"title":        tag.Name + siteName,
		"tag":          tag,
		"posts":        posts,
		"random_posts": randomPosts,
		"categories":   categories,
		"tags":         tags,
	})
}

func detailsLanding(c *gin.Context) {
	slug := c.Param("slug")
	post, err := models.NewMainManager().FindPostBySlug(slug)
	if err != nil {
		c.Set("template", "404-error.html")
		return
	}

	post.CreatedAt = utils.DateToJalaali(post.CreatedAt)
	post.UpdatedAt = utils.DateToJalaali(post.UpdatedAt)

	tags, err := models.NewMainManager().RandomTags(c, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMainManager().GetLevel1Categories(c)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMainManager().RandomPost(c, 6)
	if err != nil {
		return
	}
	for i := range randomPosts {
		randomPosts[i].CreatedAt = func() string {
			ago := utils.DateAgo(randomPosts[i].CreatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
		randomPosts[i].UpdatedAt = func() string {
			ago := utils.DateAgo(randomPosts[i].UpdatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
	}

	lastPost, err := models.NewMainManager().GetLastPost(c, 4)
	if err != nil {
		return
	}
	for i := range lastPost {
		lastPost[i].CreatedAt = utils.DateToJalaali(lastPost[i].CreatedAt)
		lastPost[i].UpdatedAt = utils.DateToJalaali(lastPost[i].UpdatedAt)
	}
	comments, err := models.NewMainManager().GetAllComments(c)
	if err != nil {
		return
	}
	c.Set("template", "blog-details.html")
	c.Set("data", map[string]interface{}{
		"title":        post.Title + siteName,
		"post":         post,
		"last_posts":   lastPost,
		"random_posts": randomPosts,
		"categories":   categories,
		"comments":     comments,
		"tags":         tags,
	})
}

func contactLanding(c *gin.Context) {
	c.Set("template", "contact.html")
	c.Set("data", map[string]interface{}{
		"title": "تماس با ما" + siteName,
	})
}

func faqLanding(c *gin.Context) {
	c.Set("template", "faq.html")
	c.Set("data", map[string]interface{}{
		"title": "سوالات متداول" + siteName,
	})
}

func pricingLanding(c *gin.Context) {
	c.Set("template", "pricing.html")
	c.Set("data", map[string]interface{}{
		"title": "تعرفه ها" + siteName,
	})
}

func servicesLanding(c *gin.Context) {
	c.Set("template", "services.html")
	c.Set("data", map[string]interface{}{
		"title": "خدمات" + siteName,
	})
}

func testimonialLanding(c *gin.Context) {
	c.Set("template", "testimonial.html")
	c.Set("data", map[string]interface{}{
		"title": "نظرات مشتریان" + siteName,
	})
}

func learnLanding(c *gin.Context) {
	c.Set("template", "learn.html")
	c.Set("data", map[string]interface{}{
		"title": "آموزش کار با سامانه" + siteName,
	})
}

func rulesLanding(c *gin.Context) {
	c.Set("template", "rules.html")
	c.Set("data", map[string]interface{}{
		"title": "قوانین فروش" + siteName,
	})
}

func returnRulesLanding(c *gin.Context) {
	c.Set("template", "return-rules.html")
	c.Set("data", map[string]interface{}{
		"title": "قوانین مرجوعی" + siteName,
	})
}
