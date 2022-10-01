package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
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
	c.Set("template", "blog.html")
	c.Set("data", map[string]interface{}{
		"title":      "بلاگ" + siteName,
		"posts":      posts,
		"last_posts": "",
		"categories": "",
		"tags":       "",
	})
}

func detailsLanding(c *gin.Context) {
	slug := c.Param("slug")
	post, err := models.NewMainManager().FindPostBySlug(c, slug)
	if err != nil {
		return
	}

	post.CreatedAt = utils.DateToJalaali(post.CreatedAt)
	post.UpdatedAt = utils.DateToJalaali(post.UpdatedAt)

	c.Set("template", "blog-details.html")
	c.Set("data", map[string]interface{}{
		"title":        post.Title + siteName,
		"post":         post,
		"last_posts":   "",
		"random_posts": "",
		"categories":   "",
		"comments":     "",
		"tags":         "",
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