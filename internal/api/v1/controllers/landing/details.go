package landing

import (
	"fmt"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

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
