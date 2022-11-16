package controllers

import (
	"fmt"
	"github.com/amirex128/selloora_backend/internal/app/models"
	"github.com/amirex128/selloora_backend/internal/app/utils"
	"github.com/amirex128/selloora_backend/internal/app/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

const (
	siteName = " | سلورا دستیار فروش اینستاگرام شما"
)

func IndexLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexLanding", "request")
	defer span.End()

	shop, domain, theme, err := models.NewMainManager().FindShopByDomain(c, ctx, c.Request.Host)
	if err != nil {
		return
	}

	c.Set("template", fmt.Sprintf("themes/%d/%s", theme.ID, "index.html"))
	c.Set("data", map[string]interface{}{
		"theme":  theme,
		"domain": domain,
		"shop":   shop,
		"title":  "صفحه اصلی" + siteName,
	})
}

func BlogLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "blogLanding", "request")
	defer span.End()

	dto, err := validations.IndexPost(c)
	if err != nil {
		return
	}
	posts, err := models.NewMainManager().GetAllPostWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMainManager().RandomTags(c, ctx, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMainManager().GetLevel1Categories(c, ctx)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMainManager().RandomPost(c, ctx, 6)
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

	shop, domain, theme, err := models.NewMainManager().FindShopByDomain(c, ctx, c.Request.Host)
	if err != nil {
		return
	}

	c.Set("template", fmt.Sprintf("themes/%d/%s", theme.ID, "blog.html"))

	c.Set("data", map[string]interface{}{
		"theme":        theme,
		"domain":       domain,
		"shop":         shop,
		"title":        "بلاگ" + siteName,
		"posts":        posts,
		"random_posts": randomPosts,
		"categories":   categories,
		"tags":         tags,
	})
}

func CategoryLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "categoryLanding", "request")
	defer span.End()
	dto, err := validations.IndexPost(c)
	if err != nil {
		return
	}
	categoryID := c.Param("id")
	category, err := models.NewMainManager().FindCategoryByID(c, ctx, utils.StringToUint64(categoryID))
	if err != nil {
		return
	}
	posts, err := models.NewMainManager().GetAllCategoryPostWithPagination(c, ctx, dto, utils.StringToUint32(categoryID))
	if err != nil {
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMainManager().RandomTags(c, ctx, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMainManager().GetLevel1Categories(c, ctx)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMainManager().RandomPost(c, ctx, 6)
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
	shop, domain, theme, err := models.NewMainManager().FindShopByDomain(c, ctx, c.Request.Host)
	if err != nil {
		return
	}

	c.Set("template", fmt.Sprintf("themes/%d/%s", theme.ID, "category.html"))

	c.Set("data", map[string]interface{}{
		"theme":        theme,
		"domain":       domain,
		"shop":         shop,
		"title":        category.Name + siteName,
		"category":     category,
		"posts":        posts,
		"random_posts": randomPosts,
		"categories":   categories,
		"tags":         tags,
	})
}

func TagLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "tagLanding", "request")
	defer span.End()
	dto, err := validations.IndexPost(c)
	if err != nil {
		return
	}
	tagSlug := c.Param("slug")
	tag, err := models.NewMainManager().FindTagBySlug(c, ctx, tagSlug)
	if err != nil {
		return
	}
	posts, err := models.NewMainManager().GetAllTagPostWithPagination(c, ctx, dto, tag.ID)
	if err != nil {
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMainManager().RandomTags(c, ctx, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMainManager().GetLevel1Categories(c, ctx)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMainManager().RandomPost(c, ctx, 6)
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
	shop, domain, theme, err := models.NewMainManager().FindShopByDomain(c, ctx, c.Request.Host)
	if err != nil {
		return
	}

	c.Set("template", fmt.Sprintf("themes/%d/%s", theme.ID, "tag.html"))

	c.Set("data", map[string]interface{}{
		"theme":        theme,
		"domain":       domain,
		"shop":         shop,
		"title":        tag.Name + siteName,
		"tag":          tag,
		"posts":        posts,
		"random_posts": randomPosts,
		"categories":   categories,
		"tags":         tags,
	})
}

func DetailsLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "detailsLanding", "request")
	defer span.End()
	slug := c.Param("slug")
	post, err := models.NewMainManager().FindPostBySlug(slug, ctx)
	if err != nil {
		c.Set("template", "404-error.html")
		return
	}

	post.CreatedAt = utils.DateToJalaali(post.CreatedAt)
	post.UpdatedAt = utils.DateToJalaali(post.UpdatedAt)

	tags, err := models.NewMainManager().RandomTags(c, ctx, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMainManager().GetLevel1Categories(c, ctx)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMainManager().RandomPost(c, ctx, 6)
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

	lastPost, err := models.NewMainManager().GetLastPost(c, ctx, 4)
	if err != nil {
		return
	}
	for i := range lastPost {
		lastPost[i].CreatedAt = utils.DateToJalaali(lastPost[i].CreatedAt)
		lastPost[i].UpdatedAt = utils.DateToJalaali(lastPost[i].UpdatedAt)
	}
	comments, err := models.NewMainManager().GetAllComments(c, ctx, post.ID)
	if err != nil {
		return
	}
	for i := range comments {
		comments[i].EmailHash = utils.GetMD5Hash(comments[i].Email)
	}
	shop, domain, theme, err := models.NewMainManager().FindShopByDomain(c, ctx, c.Request.Host)
	if err != nil {
		return
	}

	c.Set("template", fmt.Sprintf("themes/%d/%s", theme.ID, "blog-details.html"))

	c.Set("data", map[string]interface{}{
		"theme":        theme,
		"domain":       domain,
		"shop":         shop,
		"title":        post.Title + siteName,
		"post":         post,
		"last_posts":   lastPost,
		"random_posts": randomPosts,
		"categories":   categories,
		"comments":     comments,
		"tags":         tags,
	})
}

func SearchLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "SearchLanding", "request")
	defer span.End()
	dto, err := validations.IndexPost(c)
	if err != nil {
		return
	}
	posts, err := models.NewMainManager().GetAllPostWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMainManager().RandomTags(c, ctx, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMainManager().GetLevel1Categories(c, ctx)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMainManager().RandomPost(c, ctx, 6)
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
	shop, domain, theme, err := models.NewMainManager().FindShopByDomain(c, ctx, c.Request.Host)
	if err != nil {
		return
	}

	c.Set("template", fmt.Sprintf("themes/%d/%s", theme.ID, "search.html"))

	c.Set("data", map[string]interface{}{
		"theme":        theme,
		"domain":       domain,
		"shop":         shop,
		"title":        "بلاگ" + siteName,
		"posts":        posts,
		"random_posts": randomPosts,
		"categories":   categories,
		"tags":         tags,
	})
}

func PageLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "PageLanding", "request")
	defer span.End()
	shop, domain, theme, err := models.NewMainManager().FindShopByDomain(c, ctx, c.Request.Host)
	if err != nil {
		return
	}

	page, err := models.NewMainManager().FindPageBySlug(c, ctx, c.Param("slug"))
	if err != nil {
		return
	}
	c.Set("template", fmt.Sprintf("themes/%d/%s", theme.ID, "page.html"))

	c.Set("data", map[string]interface{}{
		"theme":  theme,
		"domain": domain,
		"shop":   shop,
		"page":   page,
	})
}
func ContactLanding(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "contactLanding", "request")
	defer span.End()
	c.Set("template", "contact.html")
	c.Set("data", map[string]interface{}{
		"title": "تماس با ما" + siteName,
	})
}

func FaqLanding(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "faqLanding", "request")
	defer span.End()
	c.Set("template", "faq.html")
	c.Set("data", map[string]interface{}{
		"title": "سوالات متداول" + siteName,
	})
}

func PricingLanding(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "pricingLanding", "request")
	defer span.End()
	c.Set("template", "pricing.html")
	c.Set("data", map[string]interface{}{
		"title": "تعرفه ها" + siteName,
	})
}

func ServicesLanding(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "servicesLanding", "request")
	defer span.End()
	c.Set("template", "services.html")
	c.Set("data", map[string]interface{}{
		"title": "خدمات" + siteName,
	})
}

func TestimonialLanding(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "testimonialLanding", "request")
	defer span.End()
	c.Set("template", "testimonial.html")
	c.Set("data", map[string]interface{}{
		"title": "نظرات مشتریان" + siteName,
	})
}

func LearnLanding(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "learnLanding", "request")
	defer span.End()
	c.Set("template", "learn.html")
	c.Set("data", map[string]interface{}{
		"title": "آموزش کار با سامانه" + siteName,
	})
}

func RulesLanding(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "rulesLanding", "request")
	defer span.End()
	c.Set("template", "rules.html")
	c.Set("data", map[string]interface{}{
		"title": "قوانین فروش" + siteName,
	})
}

func ReturnRulesLanding(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "returnRulesLanding", "request")
	defer span.End()
	c.Set("template", "return-rules.html")
	c.Set("data", map[string]interface{}{
		"title": "قوانین مرجوعی" + siteName,
	})
}
