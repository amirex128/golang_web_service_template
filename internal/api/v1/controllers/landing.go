package controllers

import (
	"fmt"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

const (
	siteName = " | سلورا دستیار فروش اینستاگرام شما"
)

func IndexLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()

	shop, domain, theme, err := models.NewMysqlManager(c).FindShopByDomain(c.Request.Host)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:blogLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()

	dto, err := validations.IndexPost(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	posts, err := models.NewMysqlManager(c).GetAllPostWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMysqlManager(c).RandomTags(20)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	categories, err := models.NewMysqlManager(c).GetLevel1Categories()
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	randomPosts, err := models.NewMysqlManager(c).RandomPost(6)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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

	shop, domain, theme, err := models.NewMysqlManager(c).FindShopByDomain(c.Request.Host)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:categoryLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexPost(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	categoryID := c.Param("id")
	category, err := models.NewMysqlManager(c).FindCategoryByID(utils.StringToUint64(categoryID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	posts, err := models.NewMysqlManager(c).GetAllCategoryPostWithPagination(dto, utils.StringToUint32(categoryID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMysqlManager(c).RandomTags(20)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	categories, err := models.NewMysqlManager(c).GetLevel1Categories()
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	randomPosts, err := models.NewMysqlManager(c).RandomPost(6)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	shop, domain, theme, err := models.NewMysqlManager(c).FindShopByDomain(c.Request.Host)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:tagLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexPost(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	tagSlug := c.Param("slug")
	tag, err := models.NewMysqlManager(c).FindTagBySlug(tagSlug)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	posts, err := models.NewMysqlManager(c).GetAllTagPostWithPagination(dto, tag.ID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMysqlManager(c).RandomTags(20)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	categories, err := models.NewMysqlManager(c).GetLevel1Categories()
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	randomPosts, err := models.NewMysqlManager(c).RandomPost(6)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	shop, domain, theme, err := models.NewMysqlManager(c).FindShopByDomain(c.Request.Host)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:detailsLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	slug := c.Param("slug")
	post, err := models.NewMysqlManager(c).FindPostBySlug(slug)
	if err != nil {
		c.Set("template", "404-error.html")
		return
	}

	post.CreatedAt = utils.DateToJalaali(post.CreatedAt)
	post.UpdatedAt = utils.DateToJalaali(post.UpdatedAt)

	tags, err := models.NewMysqlManager(c).RandomTags(20)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	categories, err := models.NewMysqlManager(c).GetLevel1Categories()
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	randomPosts, err := models.NewMysqlManager(c).RandomPost(6)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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

	lastPost, err := models.NewMysqlManager(c).GetLastPost(4)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	for i := range lastPost {
		lastPost[i].CreatedAt = utils.DateToJalaali(lastPost[i].CreatedAt)
		lastPost[i].UpdatedAt = utils.DateToJalaali(lastPost[i].UpdatedAt)
	}
	comments, err := models.NewMysqlManager(c).GetAllComments(post.ID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	for i := range comments {
		comments[i].EmailHash = utils.GetMD5Hash(comments[i].Email)
	}
	shop, domain, theme, err := models.NewMysqlManager(c).FindShopByDomain(c.Request.Host)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:SearchLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexPost(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	posts, err := models.NewMysqlManager(c).GetAllPostWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMysqlManager(c).RandomTags(20)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	categories, err := models.NewMysqlManager(c).GetLevel1Categories()
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	randomPosts, err := models.NewMysqlManager(c).RandomPost(6)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	shop, domain, theme, err := models.NewMysqlManager(c).FindShopByDomain(c.Request.Host)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:PageLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	shop, domain, theme, err := models.NewMysqlManager(c).FindShopByDomain(c.Request.Host)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	page, err := models.NewMysqlManager(c).FindPageBySlug(c.Param("slug"))
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:contactLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	c.Set("template", "contact.html")
	c.Set("data", map[string]interface{}{
		"title": "تماس با ما" + siteName,
	})
}

func FaqLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:faqLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	c.Set("template", "faq.html")
	c.Set("data", map[string]interface{}{
		"title": "سوالات متداول" + siteName,
	})
}

func PricingLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:pricingLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	c.Set("template", "pricing.html")
	c.Set("data", map[string]interface{}{
		"title": "تعرفه ها" + siteName,
	})
}

func ServicesLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:servicesLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	c.Set("template", "services.html")
	c.Set("data", map[string]interface{}{
		"title": "خدمات" + siteName,
	})
}

func TestimonialLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:testimonialLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	c.Set("template", "testimonial.html")
	c.Set("data", map[string]interface{}{
		"title": "نظرات مشتریان" + siteName,
	})
}

func LearnLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:learnLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	c.Set("template", "learn.html")
	c.Set("data", map[string]interface{}{
		"title": "آموزش کار با سامانه" + siteName,
	})
}

func RulesLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:rulesLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	c.Set("template", "rules.html")
	c.Set("data", map[string]interface{}{
		"title": "قوانین فروش" + siteName,
	})
}

func ReturnRulesLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:returnRulesLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	c.Set("template", "return-rules.html")
	c.Set("data", map[string]interface{}{
		"title": "قوانین مرجوعی" + siteName,
	})
}
