package controllers

import (
	"fmt"
	"github.com/amirex128/selloora_backend/internal/models"
	utils2 "github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

const (
	siteName = " | سلورا دستیار فروش اینستاگرام شما"
)

func IndexLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexLanding", "request")
	defer span.End()

	shop, domain, theme, err := models.NewMysqlManager(ctx).FindShopByDomain(c, ctx, c.Request.Host)
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
	posts, err := models.NewMysqlManager(ctx).GetAllPostWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils2.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils2.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMysqlManager(ctx).RandomTags(c, ctx, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMysqlManager(ctx).GetLevel1Categories(c, ctx)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMysqlManager(ctx).RandomPost(c, ctx, 6)
	if err != nil {
		return
	}
	for i := range randomPosts {
		randomPosts[i].CreatedAt = func() string {
			ago := utils2.DateAgo(randomPosts[i].CreatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
		randomPosts[i].UpdatedAt = func() string {
			ago := utils2.DateAgo(randomPosts[i].UpdatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
	}

	shop, domain, theme, err := models.NewMysqlManager(ctx).FindShopByDomain(c, ctx, c.Request.Host)
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
	category, err := models.NewMysqlManager(ctx).FindCategoryByID(c, ctx, utils2.StringToUint64(categoryID))
	if err != nil {
		return
	}
	posts, err := models.NewMysqlManager(ctx).GetAllCategoryPostWithPagination(c, ctx, dto, utils2.StringToUint32(categoryID))
	if err != nil {
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils2.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils2.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMysqlManager(ctx).RandomTags(c, ctx, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMysqlManager(ctx).GetLevel1Categories(c, ctx)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMysqlManager(ctx).RandomPost(c, ctx, 6)
	if err != nil {
		return
	}
	for i := range randomPosts {
		randomPosts[i].CreatedAt = func() string {
			ago := utils2.DateAgo(randomPosts[i].CreatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
		randomPosts[i].UpdatedAt = func() string {
			ago := utils2.DateAgo(randomPosts[i].UpdatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
	}
	shop, domain, theme, err := models.NewMysqlManager(ctx).FindShopByDomain(c, ctx, c.Request.Host)
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
	tag, err := models.NewMysqlManager(ctx).FindTagBySlug(c, ctx, tagSlug)
	if err != nil {
		return
	}
	posts, err := models.NewMysqlManager(ctx).GetAllTagPostWithPagination(c, ctx, dto, tag.ID)
	if err != nil {
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils2.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils2.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMysqlManager(ctx).RandomTags(c, ctx, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMysqlManager(ctx).GetLevel1Categories(c, ctx)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMysqlManager(ctx).RandomPost(c, ctx, 6)
	if err != nil {
		return
	}
	for i := range randomPosts {
		randomPosts[i].CreatedAt = func() string {
			ago := utils2.DateAgo(randomPosts[i].CreatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
		randomPosts[i].UpdatedAt = func() string {
			ago := utils2.DateAgo(randomPosts[i].UpdatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
	}
	shop, domain, theme, err := models.NewMysqlManager(ctx).FindShopByDomain(c, ctx, c.Request.Host)
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
	post, err := models.NewMysqlManager(ctx).FindPostBySlug(slug, ctx)
	if err != nil {
		c.Set("template", "404-error.html")
		return
	}

	post.CreatedAt = utils2.DateToJalaali(post.CreatedAt)
	post.UpdatedAt = utils2.DateToJalaali(post.UpdatedAt)

	tags, err := models.NewMysqlManager(ctx).RandomTags(c, ctx, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMysqlManager(ctx).GetLevel1Categories(c, ctx)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMysqlManager(ctx).RandomPost(c, ctx, 6)
	if err != nil {
		return
	}
	for i := range randomPosts {
		randomPosts[i].CreatedAt = func() string {
			ago := utils2.DateAgo(randomPosts[i].CreatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
		randomPosts[i].UpdatedAt = func() string {
			ago := utils2.DateAgo(randomPosts[i].UpdatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
	}

	lastPost, err := models.NewMysqlManager(ctx).GetLastPost(c, ctx, 4)
	if err != nil {
		return
	}
	for i := range lastPost {
		lastPost[i].CreatedAt = utils2.DateToJalaali(lastPost[i].CreatedAt)
		lastPost[i].UpdatedAt = utils2.DateToJalaali(lastPost[i].UpdatedAt)
	}
	comments, err := models.NewMysqlManager(ctx).GetAllComments(c, ctx, post.ID)
	if err != nil {
		return
	}
	for i := range comments {
		comments[i].EmailHash = utils2.GetMD5Hash(comments[i].Email)
	}
	shop, domain, theme, err := models.NewMysqlManager(ctx).FindShopByDomain(c, ctx, c.Request.Host)
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
	posts, err := models.NewMysqlManager(ctx).GetAllPostWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils2.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils2.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMysqlManager(ctx).RandomTags(c, ctx, 20)
	if err != nil {
		return
	}

	categories, err := models.NewMysqlManager(ctx).GetLevel1Categories(c, ctx)
	if err != nil {
		return
	}

	randomPosts, err := models.NewMysqlManager(ctx).RandomPost(c, ctx, 6)
	if err != nil {
		return
	}
	for i := range randomPosts {
		randomPosts[i].CreatedAt = func() string {
			ago := utils2.DateAgo(randomPosts[i].CreatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
		randomPosts[i].UpdatedAt = func() string {
			ago := utils2.DateAgo(randomPosts[i].UpdatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
	}
	shop, domain, theme, err := models.NewMysqlManager(ctx).FindShopByDomain(c, ctx, c.Request.Host)
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
	shop, domain, theme, err := models.NewMysqlManager(ctx).FindShopByDomain(c, ctx, c.Request.Host)
	if err != nil {
		return
	}

	page, err := models.NewMysqlManager(ctx).FindPageBySlug(c, ctx, c.Param("slug"))
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
